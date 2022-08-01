package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"os"
    "flag"
    "strings"
    //"time"
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

    
    
    index := 0
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        check_err(err)
        
        fmt.Printf("Problem %d %+v = ", index+1, rec[0])
        _, err = fmt.Scanf("%v", &answer)

        check_err(err)

        if strings.TrimSpace(answer) == rec[1] {
            counter++
        }
        index++ 
    }
    fmt.Printf("You have answered correct %d times\n", counter)
}