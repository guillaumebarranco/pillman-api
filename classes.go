package main

import (
	"time"
)

/******************/
/*      Medoc     */
/******************/

type Medoc struct {
    Name        string      `json:"name"`
    Cis         string      `json:"cis"`
    Dci         string      `json:"dci"`
    Forme       string      `json:"forme"` 
    Effects     string      `json:"effects"`
    Completed   bool        `json:"completed"`
    Due         time.Time   `json:"due"`
}

type Medocs []Medoc
