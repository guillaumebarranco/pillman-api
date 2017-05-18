/*
 *	Main app function creating the router and make calls listening on port 8181
 */

package main

import (
    "log"
    "net/http"
)

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8181", router))
    print("running")
}