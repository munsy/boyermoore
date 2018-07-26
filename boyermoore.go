package boyermoore

type BoyerMoore struct {
	ToSearch string
}

func NewBoyerMoore(searchString string) *BoyerMoore {
	return &BoyerMoore{
		searchString,
	}
}

func (s *BoyerMoore) Search(patternString string) bool {
	return bm(s.ToSearch, patternString)
}

func makeBadCharacterRuleTable(pattern string) map[byte]int {
	badCharacterTable := make(map[byte]int)

	for i := 0; i < len(pattern)-1; i++ {
		badCharacterTable[pattern[i]] = len(pattern) - i - 1
	}

	return badCharacterTable
}

func badCharacterRule(pattern string, table map[byte]int, letter byte) int {
	if val, found := table[letter]; found {
		return val
	}

	return len(pattern)
}

func bm(s, p string) bool {
	m := len(s)
	n := len(p)
	t := makeBadCharacterRuleTable(p)
	x := 0 // Match counter. Reset when pattern shifts.
	j := n - 1

	if m == 0 && n == 0 {
		return true
	}
	if m == 0 && n > 0 {
		return false
	}
	if n == 0 && m > 0 {
		return false
	}

	for i := n - 1; i < m; i++ {
		if i > m {
			return false
		}
		if x == n {
			return true
		}

		switch s[i] == p[j] {
		case true:
			x++
			i -= 2
			j--
		case false:
			i += badCharacterRule(p, t, s[i]) - 1
			j = n - 1
			x = 0
		}
	}
	return false
}
