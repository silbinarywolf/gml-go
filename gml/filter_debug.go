// +build debug

package gml

import "strings"

func hasFilterMatch(s string, filterBy string) bool {
	return filterBy == "" ||
		strings.Index(strings.ToLower(s), strings.ToLower(filterBy)) >= 0
}
