package main

import (
  "net/http"
  "log"
  //"fmt"
  "github.com/gorilla/mux"
  //"github.com/gorilla/securecookie"
  //"gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

// struct geoLoc{
//   ID bson.ObjectId `bson:"_id,omitempty"`
//   Lat string
//   Long string
// }

func pageHandler404(response http.ResponseWriter, request *http.Request) {

}

func indexPageHandler(response http.ResponseWriter, request *http.Request) {

}

func challengePost(response http.ResponseWriter, request *http.Request) {

}

func challengeGet(response http.ResponseWriter, request *http.Request) {

}

func uploadPhotoHandler(response http.ResponseWriter, request *http.Request) {

}

func voteHandler(response http.ResponseWriter, request *http.Request) {

}

func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
    target := "https://" + req.Host  + ":8000" +req.URL.Path
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    log.Printf("redirect to: %s", target)
    http.Redirect(w, req, target,
            // see @andreiavrammsd comment: often 307 > 301
            http.StatusTemporaryRedirect)
}

/// create a router with the gorilla mux router and handle the requests
var router = mux.NewRouter()
func main() {
    router.NotFoundHandler = http.HandlerFunc(pageHandler404)
    ///handlers for the gorilla mux router
    router.HandleFunc("/", indexPageHandler)
    router.HandleFunc("/challenge", challengePost).Methods("POST")
    router.HandleFunc("/challenge", challengeGet).Methods("GET")
    // router.HandleFunc("/location", Location).Methods("GET")
    router.HandleFunc("/uploadphoto", uploadPhotoHandler).Methods("POST")
    router.HandleFunc("/vote", voteHandler).Methods("POST")
    http.Handle("/", router)
   port := ":80"
     go http.ListenAndServe(port, http.HandlerFunc(redirect))
     port2 := ":443"
       go http.ListenAndServeTLS(port2,"cert.pem", "key.pem", http.HandlerFunc(redirect))
    http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", nil)
}
