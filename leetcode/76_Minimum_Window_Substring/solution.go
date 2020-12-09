package leetcode

// 速度比较慢，有优化空间

func minWindow(s string, t string) string {
	tail, head := -1, 0
	showCount := make(map[byte]int)
	result := ""
	for index := range t {
		if _, existed := showCount[t[index]]; !existed {
			showCount[t[index]] = 0
		}
		showCount[t[index]]++
	}

	for head = 0; head < len(s); head++ {
		if _, existed := showCount[s[head]]; existed {
			if tail == -1 {
				tail = head
			}
			showCount[s[head]]--
		}
		for tail >= 0 {
			val, existed := showCount[s[tail]]
			if existed && val >= 0 {
				break
			}
			if existed {
				showCount[s[tail]]++
			}
			tail++
		}
		valid := true
		for _, val := range showCount {
			if val > 0 {
				valid = false
				break
			}
		}
		if valid && (len(result) == 0 || len(result) > head-tail) {
			result = s[tail : head+1]
		}
	}

	return result
}
