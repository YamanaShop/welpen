package welpen

import (
	"fmt"

	"github.com/YamanaShop/hunde"
)

func Bark() string {
	return "Woof !"
}

func Barks() string {
	return "Woof! Woof! Woof !"
}

func BigBark() string {
	return hunde.WhenGrownUp(Bark())
}

func BigBarks() string {
	return hunde.WhenGrownUp(Barks())
}

func From11() string {
	fmt.Println("I am from version 1.1.0")
}
