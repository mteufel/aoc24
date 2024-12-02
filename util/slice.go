package util

func RemoveValueFromInt64Slice(s []int64, value int64) []int64 {
	var new []int64
	for _, v := range s {
		if v != value {
			new = append(new, v)
		}
	}
	return new
}

func RemoveIndexFromSlice(s []int, idx int) []int {
	var new []int
	for i, v := range s {
		if i != idx {
			new = append(new, v)
		}
	}
	return new
}
