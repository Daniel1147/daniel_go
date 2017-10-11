package dp

func rob(nums []int) int {
    var withLastMax, withoutLastMax int
    if len(nums) == 0 {
        return 0
    }

    if len(nums) == 1 {
        return nums[0]
    }

    withLastMax = nums[0]

    for i := 1; i < len(nums); i++ {
        withoutLastMax += nums[i]
        if withoutLastMax < withLastMax {
            withoutLastMax = withLastMax
        }
        withLastMax, withoutLastMax = withoutLastMax, withLastMax
    }

    if withoutLastMax > withLastMax {
        return withoutLastMax
    }
    return withLastMax
}
