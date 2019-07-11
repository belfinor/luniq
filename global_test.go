package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-07-11

import (
	"testing"
)

func TestGlobal(t *testing.T) {

	for i := 0; i < 10; i++ {

		val := Next()
		if val == "" {
			t.Fatal("Next not work")
		}

		if !Check(val, true) {
			t.Fatal("Check not work")
		}

	}

}
