package welpen

import (
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

func From12() string {
	return "I am from version 1.2.0"
}
