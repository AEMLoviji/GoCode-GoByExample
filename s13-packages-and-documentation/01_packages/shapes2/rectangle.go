package shapes2

import (
	"fmt"
	"log"
)

// Length ... your comments
var Length = 2

// Width ... your comments
var Width = 2

// Area ... your comments
func Area(l, w int) int {
	return l * w
}

func init() {
	fmt.Println("=> init() from shapes2.rectangle package.")

	if (Length < 0) || (Width < 0) {
		log.Fatal("Length & Width cannot be less than zero.")
	}
}
