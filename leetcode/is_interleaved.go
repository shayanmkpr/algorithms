package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	fmt.Println(dp)

	dp[0][0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 && s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			}
			if j > 0 && s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}

	fmt.Println(dp)
	return dp[len(s1)][len(s2)]
}
