// Random choice generator
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	choices := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your choices. Enter 'done' to select random item.)")
	for {
		fmt.Print("Enter choice: ")
		text, _ := reader.ReadString('\n')

		text = strings.TrimRight(text, "\n")
		if strings.Compare(text, "done") == 0 {
			break
		}

		choices = append(choices, text)
	}

	if len(choices) > 1 {
		rand.Seed(time.Now().Unix())
		fmt.Println("Your random choice is: ", choices[rand.Intn(len(choices))])
	} else {
		fmt.Println("You did not add any choices.")
	}
}
