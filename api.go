package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "reflect"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/davecgh/go-spew/spew"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func getDbUtil() *sql.DB {

    user := "root"
    password := ""
    host := ""
    database := "medoc"

    con, err := sql.Open("mysql", user+":"+password+"@"+host+"/"+database)

    if err != nil {  }

    return con
}

func makeResponse(items []*Medoc, w http.ResponseWriter) {

    jsonItems := json.NewEncoder(w).Encode(items)

    // debugType(jsonItems)
    spew.Dump(reflect.TypeOf(jsonItems))

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := jsonItems; err != nil {
        panic(err)
    }
}

func getMedocs(w http.ResponseWriter, r * http.Request) {

    con := getDbUtil()
    defer con.Close()

    vars := mux.Vars(r)
    limit := vars["limit"]

    rows, err := con.Query("select name, denomination as dci, cis, forme, side_effect as effects from medicaments limit 0,"+limit)

    if err != nil { spew.Dump(err) }

    items := make([]*Medoc, 0, 10)

    var cis string
    var name string
    var dci string
    var forme string
    var effects string

    for rows.Next() {
        err := rows.Scan(&cis, &name, &dci, &forme, &effects)
        if err != nil { spew.Dump(err) }

        items = append(items, &Medoc{
            Cis: cis,
            Name: name,
            Dci: dci,
            Forme: forme,
            Effects: effects,
        })
    }

    spew.Dump(reflect.TypeOf(items))

    makeResponse(items, w)
}

func getMedoc(w http.ResponseWriter, r * http.Request) {

    con := getDbUtil()
    defer con.Close()

    vars := mux.Vars(r)
    cis := vars["cis"]

    rows, err := con.Query("select name, denomination as dci, forme, side_effect as effects from medicaments where cis = "+cis)

    if err != nil { spew.Dump(err) }

    items := make([]*Medoc, 0, 10)

    var name string
    var dci string
    var forme string
    var effects string

    for rows.Next() {

        err := rows.Scan(&name, &dci, &forme, &effects)
        if err != nil { spew.Dump(err) }

        items = append(items, &Medoc{
            Name: name,
            Dci: dci,
            Forme: forme,
            Effects: effects,
        })
    }

    makeResponse(items, w)
}
