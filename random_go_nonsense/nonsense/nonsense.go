package main

import (
	"GoLang_Experiments/random_go_nonsense/loopy_thing"
	"GoLang_Experiments/random_go_nonsense/vibe_checker"
	"fmt"
	"os"
	"strconv"
)

const min_args, prid_index, min_vc_args, min_lpy_args int = 2, 1, 3, 3

func NotEnoughMsg(task_name string) {
	fmt.Printf("Not enough arguments were entered for [%s]. A minimum of two arguments are required.\n", task_name)
}

func main() {
	// There is package called "flag" which is really good for command line arguments practice. Check it out.
	numargs := len(os.Args)

	if numargs < min_args {
		// You could practice returning errors here too (e.g. implementing interface that has error() method in it.
		NotEnoughMsg("Anything Here")
		return
	}

	switch os.Args[prid_index] {
	case "vibe":
		if numargs < min_vc_args {
			NotEnoughMsg("vibe")
			break
		}
		vibe_checker.RunVibeChecker()
	case "loopy":
		if numargs < min_lpy_args {
			NotEnoughMsg("loopy")
			break
		}
		intVar, _ := strconv.Atoi(os.Args[2])
		loopy_thing.GoLoopy(intVar)
	default:
		fmt.Println("You chose nothing. Add an argument into the command line part when running this program.")
	}
}
