package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Authorization")
}

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes {

        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.HandleFunc(route.Name, saveHandler)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}
