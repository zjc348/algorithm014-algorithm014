func maxSumSubmatrix(matrix [][]int, k int) int {
	rowNum, colNum := len(matrix), len(matrix[0])
	result := math.MinInt32
	for left := 0; left < colNum; left ++ {
		rowSum := make([]int, rowNum)
		for right := left; right < colNum; right++ {
			for row := 0; row < rowNum; row++ {
				rowSum[row] += matrix[row][right]
			}
			result = max(result, maxSubArrBelowK(rowSum, k))
			if result == k {
				return k
			}
		}
	}
	return result
}

func maxSubArrBelowK(arr []int, k int) int {
	sum, max, l := arr[0], arr[0], len(arr)
	for i := 1; i < l; i++ {
		if sum > 0 {
			sum += arr[i]
		} else {
			sum = arr[i]
		}
		if sum > max {
			max = sum
		}
	}
	if max <= k {
		return max
	}
	max = math.MinInt32
	for left := 0; left < l; left++ {
		sum := 0
		for right := left; right < l; right++ {
			sum += arr[right]
			if sum > max && sum <= k {
				max = sum
			}
			if max == k {
				return k
			}
		}
	}
	return max
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}