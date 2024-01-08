package line

import "fmt"

func UserWrites() (lines []string) {
	var userInput string
	for {
		fmt.Scanln(&userInput)
		if userInput == "." {
			return
		}
		lines = append(lines, userInput)
	}
}
