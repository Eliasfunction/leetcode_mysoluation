func maximumWealth(accounts [][]int) int {
    max :=0
    for i := 0; i < len(accounts); i++ {
        tmp := 0
        for _, v := range accounts[i] { tmp += v }

        if tmp > max {
            max = tmp
        }
        
    }
    return max
}
