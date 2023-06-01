package testhelper

import (
	"fmt"
	"regexp"
	"strings"
)

func NormalizeRegexpQuery(q string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(RemoveSpace(q)))
}

// RemoveSpace, normalize query by removing space and tab character
func RemoveSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
