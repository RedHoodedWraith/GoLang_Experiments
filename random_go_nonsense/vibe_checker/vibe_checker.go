package vibe_checker

import (
	"fmt"
	"os"
)

var vibe bool

// There is principle in software engineering called single responsibility principle.
// It usually applied to classes, meaning that each class must do one thing (broadly speaking).
// Same principle applied to methods/functions too and it means that function should be pure and not produce side-effects.
// That's quite large topic, but what would be best here is just print vibe and not return anything (this is sometimes useful for method chaining, but it's whole different topic.
// And finally, it is more useful to return formatted string instead of printing it and name method FmtVibe. This will allow to use this function in anything that produces text, e.g. Logs, db writers, etc.
func CheckVibe(vibe bool) bool {
	vw := "failed"
	if vibe {
		vw = "passed"
	}

	fmt.Printf("You %s the vibe test\n", vw)
	return vibe
}

type VibeCheckable interface {
	checkVibe() bool
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

	CheckVibe(vibe)
}
