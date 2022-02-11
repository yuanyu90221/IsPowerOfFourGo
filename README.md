# IsPowerOfFourGo

Write a method to judge whether an input number is power of four

## 我的解法

如同 isPowerOfTwo 一樣

可以知道 如果轉成 base4

假設是 4 的次方數 會符合 10*$ 這樣的模式

```golang
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
```