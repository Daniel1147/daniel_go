package main

import "fmt"
import "strings"

func addBinary(a string, b string) string {
    if (len(a) > len(b)) {
        b = strings.Repeat("0", len(a) - len(b)) + b
    } else {
        a = strings.Repeat("0", len(b) - len(a)) + a
    }

    result := make([]byte, 0)

    over := byte('0')
    for i := len(a) - 1; i >= 0; i-- {
        count := over - '0' + a[i] - '0' + b[i] - '0'
        switch count {
            case 0:
                fallthrough
            case 1:
                over = '0'
                result = append([]byte{count + '0'}, result...)
            case 2:
                fallthrough
            case 3:
                over = '1'
                result = append([]byte{count - 2 + '0'}, result...)
        }
        // fmt.Println(count, result)
    }

    if (over == '1') {
        return "1" + string(result)
    }

    return string(result)
}

func main() {

    strA := "1010"
    strB := "1011"

    fmt.Println(addBinary(strA, strB))

	return
}

