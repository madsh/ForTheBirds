package main

import (    
    "fmt"
    "log"
    "net/http"
)

func root(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "<html><a href=\"https://www.pixar.com/for-the-birds\">For The Birds</a> <br> - Pixar<htlm>\n")
}

func snob(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "<html>I am OK!<br> - Snob<htlm>\n")
}

func neurotic(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.)
    fmt.Fprintf(w, "<html>Sorry, I can't do that... for legal reasons.<br> - Neurotic<htlm>\n")
}

func main() {
    http.HandleFunc("/root", root)
    http.HandleFunc("/snob", snob)
    http.HandleFunc("/neurotic", neurotic)
    log.Fatal( http.ListenAndServe(":8080", nil) )    
}


