package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"os"
)


func main() {
    answer := ""
    counter := 0
    f, err := os.Open("problem.csv")
    
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