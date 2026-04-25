package main

func convert(s string, numRows int) string {
	// r := numRows
	mat := [][]string{}
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		mat = append(mat, []string{})
	}
	for i := range s {
		cntr := i % (numRows*2 - 2)
		if cntr < numRows {
			mat[cntr] = append(mat[cntr], string(s[i]))
		} else {
			mat[2*numRows-cntr-2] = append(mat[2*numRows-cntr-2], string(s[i]))
		}
	}
	result := ""
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			result += mat[i][j]
		}
	}
	return result
}
