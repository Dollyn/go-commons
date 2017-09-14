package utils

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// ContainsArray checks whether the specific string array exists in the Two-dimensional string array.
//
func ContainsArray(slices [][]string, slice []string) bool {
	for _, sub := range slices {
		if len(sub) == len(slice) {
			if ArrayEquals(sub, slice) {
				return true
			}
		}
	}
	return false
}

func ArrayEquals(one []string, another []string) bool {
	if len(one) != len(another) {
		return false
	}
	for index, element := range one {
		if element != another[index] {
			return false
		}
	}
	return true
}
