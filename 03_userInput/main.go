package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user Input"
	fmt.Println(welcome)

	// pkg.go.dev (All packages website)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our Pizza")

	// comma ok || err err
	// input, err := reader.ReadString('\n') // same as try and catch, first variable is try and second is catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of this rating is %T: ", input)
}
