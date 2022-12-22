package loopy

import "fmt"

type Thing struct {
	number int
	name   string
	fake   bool
}

func (th Thing) String() string {
	return fmt.Sprintf("Name: %s | Number: %v | Fake?: %v", th.name, th.number, th.fake)
}

func makeThing(number int, name string, fake bool) Thing {
	return Thing{number, name, fake}
}

func GoLoopy() {
	fmt.Println("Let's Go Loopy!!!")
	t1 := Thing{0, "Zero", false}
	fmt.Println(t1)
}
