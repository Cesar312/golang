package main

// Function with a "naked" return

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	println(split(100))
}
