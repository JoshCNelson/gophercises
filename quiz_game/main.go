package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func parseProblems(lines [][]string) []problem {

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return ret
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "the path to the csv we want to read questions from")
	limit := flag.Int("limit", 30, "how many seconds before quiz ends")

	flag.Parse()

	file, _ := os.Open(*csvFile)

	reader := csv.NewReader(file)
	lines, _ := reader.ReadAll()

	problems := parseProblems(lines)

	correct := 0
	timer := time.NewTimer(time.Second * time.Duration(*limit))

	for i, problem := range problems {
		fmt.Printf("Problem %d) %v\n", i+1, problem.question)

		answerChannel := make(chan string)

		go func(c chan string) {
			var input string
			fmt.Scanf("%s", &input)
			c <- input
		}(answerChannel)

		select {
		case <-timer.C:
			fmt.Printf("You got %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerChannel:
			if answer == problem.answer {
				fmt.Println("Correct!")
				correct++
			}
		}
	}

	fmt.Printf("You got %d out of %d\n", correct, len(problems))
}
