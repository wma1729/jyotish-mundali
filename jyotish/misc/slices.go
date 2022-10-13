package misc

func StringSliceContains(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

func IntSliceContains(slice []int, element int) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}
