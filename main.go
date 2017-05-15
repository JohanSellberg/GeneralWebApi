package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hej AFA värld! Jag kör denna lilla web i en docker container med OS %s som har en %s CPU ", runtime.GOOS, runtime.GOARCH)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
