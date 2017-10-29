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

    lastPrice := prices[0]
    minPrice := prices[0]
    for i := 1; i < priceNum; i++ {
        currentPrice := prices[i]
        if (currentPrice < lastPrice) {
            profitSum += lastPrice - minPrice
            minPrice = currentPrice
        } else {
        }

        lastPrice = currentPrice
    }

    profitSum += lastPrice - minPrice

    return profitSum
}
