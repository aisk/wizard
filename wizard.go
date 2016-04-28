package wizard

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

type Answer struct {
	Content string
	Handler func()
}

type Question struct {
	Content string
	Answers []Answer
}

func Ask(questions []Question) {
	for _, question := range questions {
		printQuestion(question)
		printAnswers(question)
		handler := scanAnswerNumber(question)
		handler()
	}
}

func printQuestion(qustion Question) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("[%s] %s\n", green("?"), qustion.Content)
}

func printAnswers(question Question) {
	blue := color.New(color.FgBlue).SprintFunc()
	for i, answer := range question.Answers {
		fmt.Printf(" %s) %s\n", blue(i), answer.Content)
	}
}

func scanAnswerNumber(question Question) func() {
	for true {
		fmt.Print(" >> ")
		var input string
		fmt.Scanln(&input)
		for i, answer := range question.Answers {
			if strconv.Itoa(i) == input {
				return answer.Handler
			}
		}
		fmt.Println("invalid input.")
	}
	// unreachable path
	panic("unreachable path")
}
