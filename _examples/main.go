package main

import (
	"log"
	"os"

	"github.com/ngc224/trfile"
)

var fileNameFormat = "test_20060102.log" //Golang time format

func main() {
	w, err := trfile.New(fileNameFormat)

	if err != nil {
		os.Exit(1)
	}

	log.SetOutput(w)
	log.Println("Halo World!")
}