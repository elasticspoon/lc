package main

const mod = 1e9 + 7

func countVowelPermutation(n int) int {
	vals := [5]int{1, 1, 1, 1, 1}

	for i := 2; i <= n; i++ {
		temp := [5]int{}
		temp[0] = (vals[2] + vals[1]) % mod
		temp[1] = (vals[0] + vals[2] + vals[4]) % mod
		temp[2] = (vals[0] + vals[3]) % mod
		temp[3] = vals[2]
		temp[4] = (vals[2] + vals[3]) % mod
		vals = temp
	}

	result := 0
	for _, v := range vals {
		result = (result + v) % mod
	}

	return result
}
