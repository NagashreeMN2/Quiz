package main

import (
"encoding/csv"
"fmt"
"flag"
"os"
"strings"
"time"
)

func main(){
	csvFile := flag.String("csv", "problem.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("timer", 30, "time Limit to answer in seconds")
	flag.Parse()
	
	file, err := os.Open(*csvFile)
	if err != nil{
		exit(fmt.Sprintf("Cannot open csvfile %s\n", err))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
       fmt.Println("Cannot read lines")
	}

	problems := parserLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()
		select{
		case <-timer.C:
			fmt.Printf("you scored %d out of %d \n", correct, len(problems))
			return
		case answer:= <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("you scored %d out of %d \n", correct, len(problems))
}

func parserLines(lines [][]string) []problem{
	ret := make([]problem, len(lines))
	for i, line := range lines{
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
    return ret
}

type problem struct{
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}