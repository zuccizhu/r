package rutil

import "slices"

func Any[T any](slice []T, condition func(T) bool) bool {
	return slices.ContainsFunc(slice, condition) // 遍历完未找到，返回false
}

// 找到匹配的并支持处理匹配到的
func AnyHandle[T any](slice []T, condition func(T) bool, handle func(T)) {
	for _, item := range slice {
		if condition(item) {
			handle(item) // 找到满足条件的元素，立即返回true
		}
	}
}
