package uncategorized

func maxProfit(prices []int) int {
    var max, lessNum int

    if len(prices) == 0 {
        return 0
    }

    lessNum = prices[0]
    for _, v := range prices {
        if v < lessNum {
            lessNum = v
        } else if v - lessNum > max {
            max = v - lessNum
        }
    }

    return max
}

func maxProfit2(prices []int) int {
    profitSum := 0
    priceNum := len(prices)

    if (priceNum <= 1) {
        return profitSum
    }

    minPrice := prices[0]
    for i := 1; i < priceNum; i++ {
        // prices[i] as current price
        // prices[i - 1] as last price
        if (prices[i] < prices[i - 1]) {
            profitSum += prices[i - 1] - minPrice
            minPrice = prices[i]
        } else {
        }
    }

    profitSum += prices[len(prices) - 1] - minPrice

    return profitSum
}
