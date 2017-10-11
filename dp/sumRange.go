package dp

type NumArray struct {
    sum []int
}


func Constructor(nums []int) NumArray {
    leftSum := make([]int, len(nums) + 1)
    sum := 0
    leftSum[0] = 0
    for i, v := range nums {
        sum += v
        leftSum[i + 1] = sum
    }
    result := NumArray{leftSum}
    return result
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.sum[j + 1] - this.sum[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
