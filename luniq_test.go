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

	for i := 0; i < 300000; i++ {
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

	if u.Check("dwqedwedqwedeqwdqw", false) {
		t.Fatal("Check invalid value success")
	}

	if u.Check("10cd51ed23f2be9affffffff22eb4869ca6ea933e25471bb", false) {
		t.Fatal("Invalid time success")
	}

	if !u.Check("10cd51ed23f2be9a0000000022eb4869ca6ea933e25471bb", false) {
		t.Fatal("Valid time unsuccess")
	}

	if u.Check("10cd51ed23f2be9a0000000022eb4869ca6ea933e25471bb", true) {
		t.Fatal("Invalid sig success")
	}
}
