package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-04-17

import (
	"testing"
)

func TestUniq(t *testing.T) {

	u := New()
	defer u.Close()

	for i := 0; i < 20; i++ {
		v := u.Next()
		if v == "" {
			t.Fatal("not work")
		}
	}
}
