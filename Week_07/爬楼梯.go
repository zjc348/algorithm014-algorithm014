var cache map[int]int
func climbStairs(n int) int {
    if cache == nil{
        cache = make(map[int]int, n)
    }
    if n == 0{
        return 0
    }
    if n == 1{
        return 1
    }
    if n == 2{
        return 2
    }
    if _, ok := cache[n];ok{
        return cache[n]
    }
    temp := climbStairs(n - 1) + climbStairs(n - 2)
    cache[n] = temp
    return cache[n]
}