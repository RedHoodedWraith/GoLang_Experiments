package main

import (
	"GoLang_Experiments/random_go_nonsense/loopy"
	"GoLang_Experiments/random_go_nonsense/vibe_checker"
	"fmt"
	"os"
)

func NotEnoughMsg() {
	fmt.Println("Not enough arguments were entered. A minimum of two arguments are required.")
}

func main() {

	if len(os.Args) < 3 {
		NotEnoughMsg()
		return
	}

	switch os.Args[2] {
	case "vibe":
		if len(os.Args) < 4 {
			NotEnoughMsg()
			break
		}
		vibe_checker.RunVibeChecker()
	case "loopy":
		loopy.GoLoopy()
	default:
		fmt.Println("You chose nothing. Add an argument into the command line part when running this program.")
	}
}
