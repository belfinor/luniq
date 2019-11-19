package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-11-19

var (
	global *Uniq
)

func init() {
	global = New()
}

func Next() string {
	return global.Next()
}

func Check(val string, fullCheck bool) bool {
	return global.Check(val, fullCheck)
}
