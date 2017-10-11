package uncategorized

func maxProfit(prices []int) int {
    var max int
    for i, p := range prices {
        for _, n := range prices[i:] {
            if n - p > max {
                max = n - p
            }
        }
    }
    return max
}
