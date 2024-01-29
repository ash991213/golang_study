package main

import (
	"fmt"
	"time"
)

func main() {
	// basic_conditional()
	// var_conditional()
	// switch_case()
	// no_conditional_switch_case()
	// switch_case_with_var()
}

func basic_conditional() {
	if true {
		fmt.Println("Hello, Go")
	} else if true {
		fmt.Println("Hello, Goooo")
	} else {
		fmt.Println("Hello, Gooooooo")
	}
}

func var_conditional() {
	if redTeamScore, blueTeamScore := 2, 3; redTeamScore > blueTeamScore {
		fmt.Println("Winner : Red Team ")
	} else if redTeamScore == blueTeamScore {
		fmt.Println("Winner : No One")
	} else {
		fmt.Println("Winner : Blue Team")
	}
}

func switch_case() {
	Score := 60

	switch Score {
	case 100:
		fmt.Println("Very Good")
	case 60:
		fmt.Println("Not Bad")
	case 20:
		fmt.Println("So Bad")
	default:
		fmt.Println("are you sure??")
	}
}

func no_conditional_switch_case() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func switch_case_with_var() {
	switch Score := 100; Score {
	case 100:
		fmt.Println("Very Good")
	case 60:
		fmt.Println("Not Bad")
	case 20:
		fmt.Println("So Bad")
	default:
		fmt.Println("are you sure??")
	}
}
