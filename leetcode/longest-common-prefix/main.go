package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 0; i <= len(prefix); i++ {
		char := prefix[:i]
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j]) || char != strs[j][:i] {
				return prefix[:i-1]
			}
		}
	}

	return prefix
}
