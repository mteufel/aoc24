package util

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func RankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}

func StringToIntArray(src string, separator string) []int {
	arr := strings.Split(src, separator)
	var result []int
	for _, s := range arr {
		value, _ := strconv.Atoi(s)
		result = append(result, value)
	}
	return result
}
