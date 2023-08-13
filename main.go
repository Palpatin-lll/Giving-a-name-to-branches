package main

import (
	"fmt"
	"strings"
)

func main() {
	stroka := "Giving a name to branches" /*Change the text to the desired branch name*/
	countSim := len(stroka)
	if countSim <= 30 {
		newGenName := []string{"no"}
		newGenName = strings.Fields(stroka)
		fmt.Println(strings.Join(newGenName, "-"))
	} else {
		fmt.Println("Incorrect branch name, you need to reduce the number of characters")
	}
}
