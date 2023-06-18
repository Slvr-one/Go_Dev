package main


type Person struct {
    Firstname string
    Lastname string
}
// Write a middleware that makes sure request has Header "Content-Type" application/json
// Middleware #1
func filterContentType(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    if r.Header.Get("Content-Type") != "application/json" {
      w.WriteHeader(http.StatusUnsupportedMediaType)
      w.Write([]byte("405 - Header Content-Type incorrect"))
      return
    }
    handler.ServeHTTP(w, r)
  })
}

// Write a middleware that adds current server time to the reponse cookie
// Middleware #2
func setTimeCookie(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    // Cookie here is a struct that represents an HTTP
    // cookie as sent in the Set-Cookie header of HTTP request
    cookie := http.Cookie{
      Name: "Server Time (UTC)" // can be anything
      Value: strconv.Itoa(int(time.Now().Unix()))
      // 👆 converted time to string
    }
    // now set the cookie to response 
    http.SetCookie(w, &cookie)
    handler.ServeHTTP(w, r)
  })
}
Now let us create a handler to handle the POST request, and then we will use the middlewares:

Handler


// main handler 
func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("405 - Method Not Allowed"))
        return
    }

    var p Person
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&p)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("500 - Internal Server Error"))
        return
    }
    defer r.Body.Close()

    fmt.Printf("Got firstName and lastName as %s, %s",
    p.Firstname, p.Lastname)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("201 - Created"))
}