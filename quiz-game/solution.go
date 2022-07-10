package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"os"
    "flag"
)


func main() {
    answer := ""
    counter := 0
    wordPtr := flag.String("file", "problem.csv", "file with a quiz")
    flag.Parse()
    f, err := os.Open(*wordPtr)
    
    if err != nil {
        log.Fatal(err)
    }    

    defer f.Close()
    
    csvReader := csv.NewReader(f)
    
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        
        fmt.Printf("%+v\n", rec[0])
        fmt.Scanf("%v", &answer)
        if answer == rec[1] {
            counter++
        } 
    }
    fmt.Printf("You have answered correct %d times\n", counter)
}