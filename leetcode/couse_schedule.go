package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	seen := map[int]int{} // 0 not visited, 1 visiting, and 2 visited
	preMap := map[int][]int{}
	for _, pre := range prerequisites {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}
	var dfs func(curr int) bool
	dfs = func(curr int) bool {
		if seen[curr] == 2 {
			return true
		}
		if seen[curr] == 1 {
			return false
		}
		seen[curr] = 1
		if len(preMap[curr]) < 1 {
			seen[curr] = 2
			return true
		}
		for _, next := range preMap[curr] {
			flag := dfs(next)
			if flag == false {
				return false
			}
		}
		seen[curr] = 2
		return true
	}
	for i := range numCourses {
		if seen[i] == 2 {
			continue
		}
		flag := dfs(i)
		if flag == false {
			return false
		}
	}
	return true
}
