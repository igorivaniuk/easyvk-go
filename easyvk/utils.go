package easyvk

import "strconv"

func boolConverter(itIs bool) string {
	if itIs {
		return "1"
	}
	return "0"
}

func intIdsToString(ids []int) []string {
	var str []string
	for _, gid := range ids {
		text := strconv.Itoa(gid)
		str = append(str, text)
	}
	return str
}
