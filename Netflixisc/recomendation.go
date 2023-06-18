package reco

// https://ioshellboy.medium.com/circuit-breakers-in-golang-1779da9b001
import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

// - The recommendation service - exposes a route /recommendations;
//  that returns a list of recommended movies, while also logging the number of goutines every 500ms.

func main() {
	logGoroutines()
	http.HandleFunc("/recommendations", recoHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func logGoroutines() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("\n%v - %v", t, runtime.NumGoroutine())
			}
		}
	}()
}

func recoHandler(w http.ResponseWriter, r *http.Request) {
	a := `{"movies": ["Few Angry Men", "Pride & Prejudice"]}`
	w.Write([]byte(a))
}
