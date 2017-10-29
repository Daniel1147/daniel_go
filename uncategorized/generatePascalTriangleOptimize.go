package uncategorized

func getRow(rowIndex int) []int {
    result := []int{1}

    if rowIndex == 0 {
        return result
    }

    for i := 1; i <= rowIndex; i++ {
        var newRow []int
        newRow = make([]int, i + 1)
        for index, v := range(result) {
            newRow[index] += v
            newRow[index + 1] += v
        }
        result = newRow
    }

    return result
}
