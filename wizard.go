package wizard

import (
	"fmt"
	"github.com/bgentry/speakeasy"
	"github.com/fatih/color"
	"strconv"
)

type Input struct {
	Hidden bool
	Result *string
}

type Answer struct {
	Content string
	Handler func()
}

type Question struct {
	Content string
	Answers []Answer
	Input   *Input
}

func Ask(questions []Question) {
	for _, question := range questions {
		if question.Input != nil {
			print(">> ")
			if question.Input.Hidden {
				password, _ := speakeasy.Ask("Your LeanCloud login password (will hidden while input): ")
				question.Input.Result = &password
			} else {
				fmt.Scanln(question.Input.Result)
			}
			continue
		}
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
