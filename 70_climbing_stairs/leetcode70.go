package leetcode70


// https://leetcode-cn.com/problems/climbing-stairs/

// 递归做法
func climbStairs(n int) int {
	if n == 1  {
		return 1
	}
	if n == 2{
		return 2
	}

	return climbStairs(n - 1) + climbStairs(n - 2)

}

// 迭代做法 
func climbStairs2(n int) int {
	if n <=2 {
		return n
	}
	result := make([]int,n + 1,n + 1)
	result[1] = 1
	result[2] = 2
	for i := 3; i <= n; i++ {
		result[i] = result[i - 1] + result[i - 2]
	}
	return result[n]
}