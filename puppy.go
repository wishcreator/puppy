package puppy

import (
	"github.com/wishcreator/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
