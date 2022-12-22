package vibe_checker

import (
	"fmt"
	"os"
)

var vibe bool

func checkVibe() bool {
	vw := "failed"
	if vibe {
		vw = "passed"
	}

	fmt.Printf("You %s the vibe test\n", vw)
	return vibe
}

// RunVibeChecker
// Reports whether a user passes the "vibe test" based on whether they can write
// "yes" as an argument when running this program.
func RunVibeChecker() {
	var vibe_state_input = "no"

	if len(os.Args) >= 3 {
		vibe_state_input = os.Args[2]
	}
	vibe = vibe_state_input == "yes"

	checkVibe()
}
