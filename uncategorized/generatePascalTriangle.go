package uncategorized

func generatePascalTriangle(numRows int) [][]int {
    var result [][]int
    var lastSlice, currentSlice []int
    if numRows == 0 {
        return result
    }

    lastSlice = []int{1}
    result = append(result, lastSlice)

    for i := 1; i < numRows; i++{
        currentSlice = make([]int, i + 1)
        for j := 0; j < i; j++{
            currentSlice[j] += lastSlice[j]
            currentSlice[j + 1] += lastSlice[j]
        }
        result = append(result, currentSlice)
        lastSlice = currentSlice
    }

    return result
}
