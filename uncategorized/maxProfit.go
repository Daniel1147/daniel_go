package uncategorized

func maxProfit(prices []int) int {
    var max, lessNum int

    if len(prices) == 0 {
        return 0
    }

    lessNum = prices[0]
    for _, v := range prices {
        if v - lessNum > max {
            max = v - lessNum
        }
        if v < lessNum {
            lessNum = v
        }
    }

    return max
}
