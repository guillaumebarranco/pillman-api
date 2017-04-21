package main

import (
    "encoding/json"
    "fmt"
    //"io/ioutil"
    "net/http"
    "reflect"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/davecgh/go-spew/spew"
    "github.com/gorilla/mux"
    // "github.com/json-iterator/go"
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

func checkErr(err interface{}) {

    if err != nil {
        // spew.Dump(err)
        panic(err)
    }
}

func makeResponse(items []*Medoc, w http.ResponseWriter) {

    jsonItems := json.NewEncoder(w).Encode(items)

    // debugType(jsonItems)
    spew.Dump(reflect.TypeOf(jsonItems))

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := jsonItems; err != nil {
        panic(err)
    }
}

func makeResponseMaj(items []*Maj, w http.ResponseWriter) {

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

    rows, err := con.Query("select cis, name, denomination as dci, forme, side_effect as effects from medicaments limit 0,"+limit)

    checkErr(err)

    items := make([]*Medoc, 0, 10)

    var cis string
    var name string
    var dci string
    var forme string
    var effects string

    for rows.Next() {
        err := rows.Scan(&cis, &name, &dci, &forme, &effects)
        checkErr(err)

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

    checkErr(err)

    items := make([]*Medoc, 0, 10)

    var name string
    var dci string
    var forme string
    var effects string

    for rows.Next() {

        err := rows.Scan(&name, &dci, &forme, &effects)
        checkErr(err)

        items = append(items, &Medoc{
            Name: name,
            Dci: dci,
            Forme: forme,
            Effects: effects,
        })
    }

    makeResponse(items, w)
}

func getMedocsVersion(w http.ResponseWriter, r * http.Request) {

    con := getDbUtil()
    defer con.Close()

    rows, err := con.Query("select version from maj order by date desc limit 1")

    checkErr(err)

    items := make([]*Maj, 0, 10)

    var version string

    for rows.Next() {
        err := rows.Scan(&version)
        checkErr(err)

        items = append(items, &Maj{
            Version: version,
        })
    }

    spew.Dump(reflect.TypeOf(items))

    makeResponseMaj(items, w)
}
