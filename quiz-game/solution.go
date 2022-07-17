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
    process_file(answer, counter)
}

func check_err(err error) {
    if err != nil {
        fmt.Println("Something went wrong:", err)
        log.Fatal(err)
    }
}

func process_file(answer string, counter int) {
    wordPtr := flag.String("file", "problem.csv", "file with a quiz")
    flag.Parse()
    f, err := os.Open(*wordPtr)
    
    check_err(err)  

    defer f.Close()
    
    csvReader := csv.NewReader(f)
    
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        check_err(err)
        
        fmt.Printf("%+v\n", rec[0])
        _, err = fmt.Scanf("%v", &answer)

        check_err(err)

        if answer == rec[1] {
            counter++
        } 
    }
    fmt.Printf("You have answered correct %d times\n", counter)
}