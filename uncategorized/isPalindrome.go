package uncategorized

// import "fmt"

func isPalindrome(s string) bool {
	alphaStr := toAlpha(s)

    // fmt.Println(alphaStr)

    alphaCnt := len(alphaStr)
    for i := 0; i < alphaCnt / 2; i++ {
        if (alphaStr[i] != alphaStr[alphaCnt - i - 1]) {
            return false
        }
    }

    return true
}

func toAlpha(s string) string {
    strlen := len(s)
    newStr := make([]byte, strlen)
    newCnt := 0

    for _, v := range s {
        if (v >= 'a' && v <= 'z') {
            newStr[newCnt] = byte(v)
            newCnt++
            continue
        }

        if (v >= 'A' && v <= 'Z') {
            newStr[newCnt] = byte(v) - 'A' + 'a'
            newCnt++
            continue
        }

        if (v >= '0' && v <= '9') {
            newStr[newCnt] = byte(v)
            newCnt++
            continue
        }
    }

    return string(newStr[:newCnt])
}
