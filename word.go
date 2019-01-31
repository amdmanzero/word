package word

import "strings"

func WordCount(s string) map[string]int {
	arr := WordSplit(s)

	dup := make(map[string]int)

	for _, x := range arr {

		_, exist := dup[x]

		if exist {
			dup[x] += 1
		} else {
			dup[x] = 1
		}

	}

	for k, v := range dup {
		if v == 1 {

			delete(dup, k)
		}
	}

	return dup
}
func WordSplit(str string) []string {
	return strings.Fields(str)
}
