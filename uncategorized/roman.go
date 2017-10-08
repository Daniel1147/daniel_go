package main

import "fmt"

func romanToInt(s string) int {
	sum := 0

    // fmt.Println(statusList)

    for charIndex := len(s) - 1; charIndex >= 0; charIndex-- {
        lastChar := s[charIndex]
        switch lastChar {
            case 'I':
                if (sum >= 5) {
                    sum -= 1
                } else {
                    sum += 1;
                }

            case 'X':
                if (sum >= 50) {
                    sum -= 10
                } else {
                    sum += 10;
                }

            case 'C':
                if (sum >= 500) {
                    sum -= 100
                } else {
                    sum += 100;
                }

            case 'V':
                sum += 5;

            case 'L':
                sum += 50;

            case 'D':
                sum += 500;

            case 'M':
                sum += 1000;

        }
    }

    return sum
}

func main() {
	fmt.Println(romanToInt("IV"))
	// fmt.Println(romanToInt("MCMIV"))
	// fmt.Println(romanToInt("MCMLIV"))
	// fmt.Println(romanToInt("MCMXC"))
	// fmt.Println(romanToInt("MMXIV"))
	// fmt.Println(romanToInt("XXII"))
}
