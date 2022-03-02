package lettercombinations

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{}
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var dfs func(res string, digits string)
	dfs = func(res string, digits string) {
		if len(digits) == 0 {
			result = append(result, res)
			return
		}
		for _, v := range mp[string(digits[0])] {
			res = res + string(v)
			dfs(res, digits[1:])
			res = res[:len(res)-1] //回溯
		}
	}
	dfs("", digits)
	return result
}
