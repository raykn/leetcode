package money

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// TODO 买卖股票的最佳时机 ———— 最大值
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	r := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > r {
			r = prices[i]
		} else if prices[i] < r {
			max = Max(max, r-prices[i])
		}
	}
	return max
}

// TODO 买卖股票的最佳时机 II ———— 多次交易
func maxProfit2_(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	r := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > r {
			r = prices[i]
		} else if prices[i] < r {
			max += r - prices[i]
			r = prices[i]
		}
	}
	return max
}

// TODO 买卖股票的最佳时机 II ———— 多次交易 100%
func maxProfit2__(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		res += Max(0, prices[i+1]-prices[i])
	}
	return res
}
