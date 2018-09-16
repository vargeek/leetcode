package treesandgraphs

// Course Schedule II
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/848/

func findOrder(numCourses int, prerequisites [][]int) []int {
	degree := make([]int, numCourses)
	follows := make([][]int, numCourses)
	for _, pre := range prerequisites {
		follows[pre[1]] = append(follows[pre[1]], pre[0])
		degree[pre[0]]++
	}

	stack := make([]int, 0, numCourses)
	for c, d := range degree {
		if d == 0 {
			stack = append(stack, c)
		}
	}

	courses := make([]int, 0, numCourses)
	for i := 0; i < len(stack); i++ {
		c := stack[i]
		courses = append(courses, c)
		for _, follow := range follows[c] {
			degree[follow]--
			if degree[follow] == 0 {
				stack = append(stack, follow)
			}
		}
	}
	if len(courses) != numCourses {
		return []int{}
	}
	return courses
}
