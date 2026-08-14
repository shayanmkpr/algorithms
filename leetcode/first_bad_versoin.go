package main

func isBadVersion(v int) bool {
	return false
}

// everthing after the first bad version is a bad version.
func firstBadVersion(n int) int {
	l, h := 1, n
	for l < h {
		curr := l + (h-l)/2
		if isBadVersion(curr) {
			h = curr
		} else {
			l = curr + 1
		}
	}
	return h
}
