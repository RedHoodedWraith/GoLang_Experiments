package loopy_thing

import (
	"GoLang_Experiments/random_go_nonsense/vibe_checker"
	"fmt"
)

var true_thing_id_nums = []int{1, 2}

type Thing struct {
	number int
	name   string
	canon  bool
}

func (th *Thing) String() string {
	return fmt.Sprintf("Name: %s | Number: %v | Canon?: %v", th.name, th.number, th.canon)
}

func (th *Thing) printThing() {
	fmt.Println(th)
}

// Demonstrates using an interface function
func (th *Thing) checkVibe() bool {
	fmt.Printf("\nVibe Check: Does %s pass the Vibe Test?\n\t%s, ", th.name, th.name)
	return vibe_checker.CheckVibe(th.canon)
}

func (th *Thing) UpdateThing(new_num int, new_name string, new_fake_state bool) {
	th.number = new_num
	th.name = new_name
	th.canon = new_fake_state
}

func (th *Thing) updateNumber(new_num int) {
	th.UpdateThing(new_num, fmt.Sprintf("Thing %v", new_num), th.canon)
}

func (th *Thing) makeThingCanon() {
	th.canon = true
}

func (th *Thing) makeThingNotCanon() {
	th.canon = false
}

func makeThing(number int, name string, canon bool) *Thing {
	return &Thing{number, name, canon}
}

func makeThingSimple(number int, canon bool) *Thing {
	return makeThing(number, fmt.Sprintf("Thing %v", number), canon)
}

// PrintThingTable Prints all the Things from the Collection
func PrintThingTable(things_arr []*Thing) {
	for _, th := range things_arr {
		th.printThing()
	}
}

// Deprecated Demonstration Code
func demoLoopy() {
	t1 := &Thing{0, "Zero", false}
	t2 := makeThing(1, "One", true)
	t1.printThing()
	t2.printThing()

	fmt.Println("\nAfter Updating Thing 0:")
	t2.UpdateThing(0, "Zero", true)
	t2.printThing()
}

func GoLoopy(desired_num_things int) {
	fmt.Println("Let's Go Loopy!!! Let's make some Things!...")

	// A Slice/Array to store all the Things
	all_things := make([]*Thing, desired_num_things)

	fmt.Println("\nFirst, let's add some things...")
	// Creates Things based on the number requested
	for i := 0; i < desired_num_things; i++ {
		all_things[i] = makeThingSimple(i+1, false)
	}

	// Prints all the Things in the Table
	PrintThingTable(all_things)

	// Checks if a Thing is meant to be canon (based on whether it appeared in a Dr Seuss book)
	for _, th := range all_things {
		for _, vn := range true_thing_id_nums {
			if th.number == vn {
				th.makeThingCanon()
			}
		}
	}

	// Prints all the Things in the Table after Update
	fmt.Println("\nAfter correcting cannon Things")
	PrintThingTable(all_things)

	fmt.Println("\nVibe Check on all the Things...")
	for _, th := range all_things {
		th.checkVibe()
	}
}
