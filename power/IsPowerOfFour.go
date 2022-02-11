package power

import (
	"regexp"
	"strconv"
)

func IsPowerOfFour(n int64) bool {
	base4Pattern := "^10*$"
	input := strconv.FormatInt(n, 4)
	match, _ := regexp.MatchString(base4Pattern, input)
	return match
}
