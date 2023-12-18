package furtherexplored

import "fmt"

// This is ALSO a 'Package Block' scope
// This IS exported given the identifier 'Fascinating' is capitalized

func Fascinating() {

	fmt.Println("Planet Speed: ", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planet Spinning Speed: ", planetRotationSpeed)
}
