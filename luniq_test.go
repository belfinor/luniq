package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2019-11-19

import (
	"testing"
)

func TestUniq(t *testing.T) {

	u := New()
	defer u.Close()

	if !u.Check("c5b2ec53b1f7e48275dbb35aad333655a9fad1400b05a4781", true) {
		t.Fatal("Check old failed")
	}

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

	if u.Check("f10cd51ed23f2be9affffffff22eb4869ca6ea933e25471bb", false) {
		t.Fatal("Invalid time success")
	}

	if !u.Check("f10cd51ed23f2be9a0000000022eb4869ca6ea933e25471bb", false) {
		t.Fatal("Valid time unsuccess")
	}

	if u.Check("f10cd51ed23f2be9a0000000022eb4869ca6ea933e25471bb", true) {
		t.Fatal("Invalid sig success")
	}
}
