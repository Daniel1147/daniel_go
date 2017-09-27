package main

import "fmt"

func lengthOfLastWord(s string) int {
    lastLength := 0
    counter := 0
    for i:= 0; i < len(s); i++ {
        char := s[i]
        if (char != ' ') {
            counter++
            lastLength = counter
        } else {
            counter = 0
        }
    }
    return lastLength
}

func main() {
    input := "Hello World"
    fmt.Println(lengthOfLastWord(input))

    input = ""
    fmt.Println(lengthOfLastWord(input))

    input = "a "
    fmt.Println(lengthOfLastWord(input))

	return
}

