package welpen

import (
	"github.com/YamanaShop/hunde"
)

func Bark() string {
	return "WOOF !"
}

func Barks() string {
	return "WOOF! WOOF! WOOF !"
}

func BigBark() string {
	return hunde.WhenGrownUp(Bark())
}

func BigBarks() string {
	return hunde.WhenGrownUp(Barks())
}
