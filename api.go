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

type User struct {

    // The `json` struct tag maps between the json name
    // and actual name of the field
    Denomination string `json:"denomination"`
}

func getOpenMedocs(w http.ResponseWriter, r * http.Request) {

    // var query = "a"
    // var page = "1"
    // var limit = "10"
    // url := "https://www.open-medicaments.fr/api/v1/medicaments?query="+query+"&page="+page+"&limit="+limit

    url := "https://api.stackexchange.com/2.2/tags?page=1&pagesize=100&order=desc&sort=popular&site=stackoverflow"

    res, _ := http.Get(url)
    defer res.Body.Close()

    var data struct {
        Items []struct {
            Name                string
            Count               int
            Is_required         bool
            Is_moderator_only   bool
            Has_synonyms        bool
        }
    }

    // var data struct {
    //     Test []struct {
    //         Denomination        string
    //     }
    // }

    dec := json.NewDecoder(res.Body)
    dec.Decode(&data)

    for _, item := range data.Items {
        fmt.Printf("%s = %d\n", item.Name, item.Count)
        // fmt.Printf("%s = %d\n", item.Denomination)
    }

}

// func getOpenMedocs(w http.ResponseWriter, r * http.Request) {
//     var data struct {
//         Items []struct {
//             Name              string
//             Count             int
//             Is_required       bool
//             Is_moderator_only bool
//             Has_synonyms      bool
//         }
//     }

//     res, _ := http.Get("https://api.stackexchange.com/2.2/tags?page=1&pagesize=100&order=desc&sort=popular&site=stackoverflow")
//     defer res.Body.Close()

//     dec := json.NewDecoder(res.Body)
//     dec.Decode(&data)

//     for _, item := range data.Items {
//         fmt.Printf("%s = %d\n", item.Name, item.Count)
//     }
// }
