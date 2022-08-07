package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"os"
    "flag"
    "strings"
    "time"
)


func main() {
    sourcefile := flag.String("file", "problem.csv", "file with a quiz")
    timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
    flag.Parse()
    f, err := os.Open(*sourcefile)
    
    check_err(err)  

    defer f.Close()
    
    csvReader := csv.NewReader(f)

    timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

    //fmt.Printf("%T\n%T\n", csvReader, timer)

    process_strings(csvReader, timer)

    
    
}

func check_err(err error) {
    if err != nil {
        fmt.Println("Something went wrong:", err)
        log.Fatal(err)
    }
}

func process_strings(csvReader *csv.Reader, timer *time.Timer) {
    answer := ""
    index := 0
    counter := 0
    MainLoop:
        for {
            rec, err := csvReader.Read()
            if err == io.EOF {
                break
            }
            check_err(err)
            answerCh := make(chan string)
            fmt.Printf("Problem %d %+v = ", index+1, rec[0])
            go func(){
                _, err = fmt.Scanf("%v", &answer)
                check_err(err)
                answerCh <- answer
            }()
            select {
            case <-timer.C:
                fmt.Println("\nTime us up!")
                break MainLoop
            
            case answer = <-answerCh:
                if strings.TrimSpace(answer) == rec[1] {
                    counter++
                }
            index++
            } 
        }
        fmt.Printf("\nYou have answered correct %d times\n", counter)
}