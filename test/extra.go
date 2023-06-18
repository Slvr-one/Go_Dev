package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	// "gopkg.in/mgo.v2/bson"
	// "rsc.io/quote"
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/propagation"
	// "go.opentelemetry.io/otel/trace"
)

func main() {
	if os.Getenv("SERVICE_NAME") != "" {
		serviceName = os.Getenv("SERVICE_NAME")
	}

	// checking for a local .env file containing vars - redundant as of now
	envLoadErr := godotenv.Load()
	helpers.HandleErr(envLoadErr, "err loading .env file.")

	// init a server client with custom spec - for listen & serve
	mainServer := &http.Server{
		Addr:           ":" + *listenAddr,
		Handler:        InitR(), //router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // "1 times 2, 20 times" or 1048576 - standard size of header :)
		// BaseContext: func(l net.Listener) context.Context {
		// 	ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
		// 	return ctx
		// },
	}

	// log.Fatal(mainServer.ListenAndServe())
	// c := cors.New(cors.Options{
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	// })

	// handler := c.Handler(router)
	// http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
	// log.Fatal(http.ListenAndServe(":"+serverPort, router))
	// log.Fatal(http.ListenAndServe(":"+*listenAddr, router))
}

// using mux router -> (old)
func InitR() *mux.Router {
	router := mux.NewRouter()
	// fs := http.FileServer(http.Dir("static"))

	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")
	// router.HandleFunc("/GH/{name}", GetHorse).Methods("GET")       //Get a specific horse

	router.HandleFunc("/", Welcom).Methods("GET")
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/metrics", Monitor).Methods("GET")

	router.HandleFunc("/api/horses", GetHorses).Methods("GET") //List all available horses
	router.HandleFunc("/api/horses", CreateHorse).Methods("POST")
	router.HandleFunc("/api/horses/{name}", updateHorse).Methods("GET") //Get a specific horse
	router.HandleFunc("/api/horses/{name}", updateHorse).Methods("PUT") //Update a specific horse
	router.HandleFunc("/api/horses/{name}", DeleteHorse).Methods("DELETE")

	router.HandleFunc("/invest/{investor}/{horse}/{amount}", Invest).Methods("GET")

	// router.Handle("/", http.FileServer(http.Dir("templates/styles/")))
	// router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/styles"))))

	styles := http.FileServer(http.Dir("./templates/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	// router.Handle("/styles/", styles)
	return router
}

// http://localhost:9000/invest/%7BDangerous%7D/%7B500%7D

// // Define some constants representing the rules of the game
// const MaxBet = 10 // The maximum allowable bet

func httpErrorBadRequest(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusBadRequest)
}

func httpErrorInternalServerError(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusInternalServerError)
}

func httpError(err error, ctx *gin.Context, status int) {
	log.Println(err.Error())
	ctx.String(status, err.Error())
}

func pingHandler(ctx *gin.Context) {
	req := resty.New().R().SetHeader("Content-Type", "application/text")
	otelCtx := ctx.Request.Context()
	span := trace.SpanFromContext(otelCtx)
	defer span.End()
	otel.GetTextMapPropagator().Inject(otelCtx, propagation.HeaderCarrier(req.Header))
	url := ctx.Query("url")
	if len(url) == 0 {
		url = os.Getenv("PING_URL")
		if len(url) == 0 {
			httpErrorBadRequest(errors.New("url is empty"), ctx)
			return
		}
	}
	log.Printf("Sending a ping to %s", url)
	resp, err := req.Get(url)
	if err != nil {
		httpErrorBadRequest(err, ctx)
		return
	}
	log.Println(resp.String())
	ctx.String(http.StatusOK, resp.String())
}

// ///////////////////////////////////
func execute() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	}

	out, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

// func useCPU is a long loop for wasting cpu
func useCPU(w http.ResponseWriter, r *http.Request) {
	count := 1

	for i := 1; i <= 1000000; i++ {
		count = i
	}

	fmt.Printf("count: %d", count)
	w.Write([]byte(fmt.Sprint(count)))
}

func userHandler(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("id")
	if id == "" {
		http.Error(ctx.Writer, "The id query parameter is missing", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(ctx.Writer, "<h1>The user id is: %s</h1>", id)
}

func searchHandler(ctx *gin.Context) {
	u, err := url.Parse(ctx.Request.URL.String())

	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")

	if page == "" {
		page = "1"
	}

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)
}

func serve(ctx *gin.Context) {
	env := map[string]string{}
	for _, keyval := range os.Environ() {
		keyval := strings.SplitN(keyval, "=", 2)
		env[keyval[0]] = keyval[1]
	}
	bytes, err := json.Marshal(env)
	if err != nil {
		ctx.Writer.Write([]byte("{}"))
		return
	}
	ctx.Writer.Write([]byte(bytes))
}

// a func to get the static assets
func staticHandler() http.Handler {
	return http.FileServer(http.Dir("./static"))
}

// LogRequest -> logs req info
// func LogRequest(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(ctx *gin.Context) {
// 		color.Yellow("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
// 		handler.ServeHTTP(w, r)
// 	})
// }

// removeBets remove the occurrences of an element from a slice, return the new slice
func removeBets(slice []string, elem string) []string {
	// Create a new slice to store the result
	newSlice := make([]string, 0, len(slice))
	// Loop over the elements in the original slice
	for _, value := range slice {
		// If the element is not the one to remove, add it to the new slice
		if value != elem {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func Status(ctx *gin.Context) {
	ctx.Writer.Write([]byte("API is up and running"))
}

// type Game struct {
// 	Players []Player         // A list of all active players in the game
// 	Horses  map[string]Horse // A mapping from each horse name to its details
// 	Stakes  map[int64]*Stake // A mapping from each player ID to their current stake
// }

var Html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="description" content="This is a demo of the Quote API">
<meta name="go-import" content="github.com/rsc/quote"
<style>
body {
	font-family: Arial, Helvetica, sans-serif;
	font-size: 14px;
	font-weight: 400;
	color: #000000;
	background-color: #fff;
	padding: 0;
	margin: 0;
	text-align: center;
	text-decoration: none;
	border: 1px solid #ccc;
	border-radius: 5px;
	border-color: #ccc;
	border-style: solid;
	border-width: 1px;
	border-top-left-radius: 5px;
	border-top-right-radius: 5px;
	border-bottom-left-radius: 5px;
	border-bottom-right-radius: 5px;
	border-top: 1px solid #ccc;
	border-bottom: 1px solid #ccc;
	border-left: 1px solid #ccc;
	border-right: 1px solid #ccc;
}
.container {
	margin: 0;
	padding: 0;
}
.quote {
	margin: 0;
	padding: 0;
}
</style>
</head>
<body>
<div class="container">
<div class="quote">
<h1>Welcome to the Quote API!</h1>
<p>This is a demo of the Quote API. You can find the source code <a href="https://github.com/rsc/quote">here</a>.</p>
</div>
</body>
</html>`
