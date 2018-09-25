package stuff

import (
	"fmt"

	"github.com/jamesward/godiamond/things"
)

func Print(something things.Something) {
	fmt.Println(something.Name)
}
