package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-07-11

import (
	"hash/crc32"
)

var (
	crctab *crc32.Table
	global *Uniq
)

func init() {
	crctab = crc32.MakeTable(crc32.IEEE)
	global = New()
}

func Next() string {
	return global.Next()
}

func Check(val string, fullCheck bool) bool {
	return global.Check(val, fullCheck)
}
