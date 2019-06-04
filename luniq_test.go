package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-06-04

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

		if !u.Check(v, false) {
			t.Fatal("Check full=false failed")
		}

		if !u.Check(v, true) {
			t.Fatal("Check full=true failed")
		}
	}
}
