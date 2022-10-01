package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	Question string
	Answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question/answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Deu ruim ao abrir o arquivo: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Something goes wrong on read csv.")
	}

	problems := parseLines(lines)
	var correct int32
	var wrong int32

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.Answer {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Wrong!")
			wrong++
		}
	}

	fmt.Printf("The final score is: Correct %v | Wrong %v\n", correct, wrong)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			Question: line[0],
			Answer:   line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
