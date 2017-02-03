package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/all/{limit}",
        getMedocs,
    },
    Route{
        "Medoc",
        "GET",
        "/medoc/{cis}",
        getMedoc,
    },
    Route{
        "OpenMedocs",
        "GET",
        "/open/all",
        getOpenMedocs,
    },
}
