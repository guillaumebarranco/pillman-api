package main

import (
	"time"
)

/******************/
/*      Medoc     */
/******************/

type Medoc struct {
    Cis         string      `json:"cis"`
    Name        string      `json:"name"`
    Dci         string      `json:"dci"`
    Forme       string      `json:"forme"` 
    Effects     string      `json:"effects"`
    Completed   bool        `json:"completed"`
    Due         time.Time   `json:"due"`
}

type Medocs []Medoc

type Maj struct {
    Version     string  `json:version`
}

type Majs []Maj
