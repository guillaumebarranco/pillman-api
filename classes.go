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

type OpenMedoc struct {
    Denomination    	string      `json:"denomination"`
    CodeCIS     		string      `json:"codeCIS"`
    Completed   		bool        `json:"completed"`
    Due         		time.Time   `json:"due"`
}

type OpenMedocs []OpenMedoc