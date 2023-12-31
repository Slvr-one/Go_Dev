package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

var downstreamErrCount int
var circuitOpenErrCount int

// Movies App
type MovieResponse struct {
	Feed           []string
	Recommendation []string
}

func main() {
	downstreamErrCount = 0
	circuitOpenErrCount = 0
	hystrix.ConfigureCommand("recommendation", hystrix.CommandConfig{
		Timeout:                100,  // The wait time for upstream service response.
		RequestVolumeThreshold: 25,   // The number of requests untill circuit breaker wouldn’t evaluate if the state needs to be changed.
		ErrorPercentThreshold:  5,    //The threshold percentage of errors at which the circuit would short.
		SleepWindow:            1000, // The delay between open, and half open state.
	})

	http.HandleFunc("/movies", fetchMoviesFeedHandlerWithCircuitBreaker)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func fetchMoviesFeedHandler(w http.ResponseWriter, r *http.Request) {
// 	mr := MovieResponse{
// 		Feed: []string{"Transformers", "Fault in our stars", "The Old Boy"},
// 	}
// 	rms, err := fetchRecommendations()
// 	if err != nil {
// 		w.WriteHeader(500)
// 	}
// 	mr.Recommendation = rms
// 	bytes, err := json.Marshal(mr)
// 	if err != nil {
// 		w.WriteHeader(500)
// 	}
// 	w.Write(bytes)
// }

func fetchMoviesFeedHandlerWithCircuitBreaker(w http.ResponseWriter, r *http.Request) {
	mr := MovieResponse{
		Feed: []string{"Transformers", "Fault in our stars", "The Old Boy"},
	}

	//circuit breaker
	output := make(chan bool, 1)
	errors := hystrix.Go("recommendation", func() error {
		// talk to other services
		rms, err := fetchRecommendations()
		if err != nil {
			return err
		}
		mr.Recommendation = rms
		output <- true
		return nil
	}, func(err error) error {
		// Your fallback logic goes here
		return nil
	})

	select {
	case err := <-errors:
		if err == hystrix.ErrCircuitOpen {
			circuitOpenErrCount = circuitOpenErrCount + 1
		} else {
			downstreamErrCount = downstreamErrCount + 1
		}

	case _ = <-output:

	}

	bytes, err := json.Marshal(mr)
	if err != nil {
		w.WriteHeader(500)
	}
	fmt.Printf("\ndownstreamErrCount=%d, circuitOpenErrCount=%d", downstreamErrCount, circuitOpenErrCount)
	w.Write(bytes)
}

func fetchRecommendations() ([]string, error) {
	resp, err := http.Get("http://localhost:9090/recommendations")
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var mvsr map[string]interface{}
	err = json.Unmarshal(body, &mvsr)
	if err != nil {
		return []string{}, err
	}
	mvsb, err := json.Marshal(mvsr["movies"])
	if err != nil {
		return []string{}, err
	}
	var mvs []string
	err = json.Unmarshal(mvsb, &mvs)
	if err != nil {
		return []string{}, err
	}
	return mvs, nil
}
