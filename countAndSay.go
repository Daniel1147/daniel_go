package main

import "fmt"

func countAndSay(n int) string {
    if (n == 1) {
        return "1"
    }

    str := "1"
    for i := 1; i < n; i++ {
        str = iterateByte(str)
    }

    return str
}

func iterateByte(s string) string {
    newStr := make([]byte, 0)
    lastChar := s[0]
    var counter byte
    counter = 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if (char == lastChar) {
            counter ++
        } else {
            newStr = append(newStr, counter + '0')
            newStr = append(newStr, lastChar)
            lastChar = char
            counter = 1
        }
    }
    newStr = append(newStr, counter + '0')
    newStr = append(newStr, lastChar)

    return string(newStr)
}

func main() {

    for i := 1; i <= 10; i++ {
        fmt.Println(countAndSay(i))
    }

	return
}

