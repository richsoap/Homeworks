package leetcode

func isMatch(s string, p string) bool {
	return tryMatch(s+"x", p+"x", 0, 0)
}

func tryMatch(s, p string, sp, pp int) bool {
	if sp >= len(s) {
		return pp >= len(p)
	}
	if pp >= len(p) {
		return sp >= len(s)
	}
	curr := p[pp]
	next := curr

	if pp < len(p)-1 {
		next = p[pp+1]
	}

	if next == '*' {
		for i := 0; i <= len(s)-sp && (i == 0 || matchChar(s[sp+i-1], curr)); i++ {
			if tryMatch(s, p, sp+i, pp+2) {
				return true
			}
		}
	} else {
		return matchChar(s[sp], curr) && tryMatch(s, p, sp+1, pp+1)
	}
	return false
}

func matchChar(a, b byte) bool {
	return b == '.' || a == b
}
