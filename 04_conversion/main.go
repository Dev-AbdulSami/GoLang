package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	// Trim the newline characters from the input
	input = strings.TrimSpace(input)

	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Rating after add: ", numRating+1)

	}

}
