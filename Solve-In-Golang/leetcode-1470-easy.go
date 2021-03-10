func shuffle(nums []int, n int) []int {
    Result := make([]int,0) 
    for i := 0; i < n; i++{
        Result = append(Result,nums[i],nums[i+n])
    }
    return Result
}
