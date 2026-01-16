package utils

import (
	"fmt"
	"strings"
)

func JoinAdvanced[T any](
	items []T,
	separator, lastSeparator string,
	format func(int, T) string,
) string {
	var result strings.Builder

	for i, item := range items {
		if i > 0 {
			if i == len(items)-1 {
				result.WriteString(lastSeparator)
			} else {
				result.WriteString(separator)
			}
		}
		result.WriteString(format(i, item))
	}

	return result.String()
}

func JoinWithLast[T any](items []T, sep, lastSep string) string {
	return JoinAdvanced(
		items,
		sep, lastSep,
		func(_ int, s T) string { return fmt.Sprint(s) },
	)
}

func JoinWithMapFunc[T any](items []T, sep string, format func(int, T) string) string {
	return JoinAdvanced(items, sep, sep, format)
}

func Join(items []string, sep string) string {
	return JoinWithMapFunc(items, sep, func(_ int, s string) string { return s })
}
