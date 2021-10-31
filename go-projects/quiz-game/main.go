package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type quiz struct {
	question, answer string
}

type result struct {
	question, wrongAnswer, correctAnswer int
}

func main() {
	fmt.Println("Mathematic Quiz")

	fmt.Println("Customize timer(Default 30seconds): (Y/N)")
	var timerSelection string
	fmt.Scanln(&timerSelection)

	var timer time.Duration = 30

	if strings.ToLower(timerSelection) == "y" {
		fmt.Println("Input the time needed for the quiz: ")
		fmt.Scanln(&timer)
	}

	fmt.Println("Quiz Start now! You have", timer, "seconds")

	quizTimer := time.NewTimer(timer * time.Second)
	go func() {
		<-quizTimer.C
		fmt.Println("Time is up!")
	}()

	readQuizQuestions()

	quizTimer.Stop()
}

func readQuizQuestions() result {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	question := 0
	correctAnswer := 0
	wrongAnswer := 0

	for _, line := range csvLines {
		question++
		q := quiz{
			question: line[0],
			answer:   line[1],
		}
		fmt.Println("Question", question, ": ", q.question)
		var answer string
		fmt.Scanln(&answer)

		if answer == q.answer {
			correctAnswer++
		} else {
			wrongAnswer++
		}
	}

	r := result{}
	r.calcResult(question, wrongAnswer, correctAnswer)
	return r
}

func (r result) calcResult(question, wrongAnswer, correctAnswer int) {
	r.question = question
	r.correctAnswer = correctAnswer
	r.wrongAnswer = wrongAnswer

	fmt.Println("Correct:", r.correctAnswer)
	fmt.Println("Wrong:", r.wrongAnswer)
	fmt.Println("Total Questions:", r.question)
}
