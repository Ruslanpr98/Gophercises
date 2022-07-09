package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"os"
)


func main() {
    // open file
    f, err := os.Open("problem.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        // do something with read line
        fmt.Printf("%+v\n", rec)
    }
}