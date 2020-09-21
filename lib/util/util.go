package util

func UniqSlice(slice []string) []string {
	var uniqSlice []string
	checkerMap := make(map[string]bool)

	for _, v := range slice {
		if !checkerMap[v] {
			checkerMap[v] = true
			uniqSlice = append(uniqSlice, v)
		}
	}
	return uniqSlice
}