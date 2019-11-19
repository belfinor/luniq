package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.007
// @date    2019-11-19

import (
	"context"
	"fmt"
	"hash/crc32"
	"os"
	"strconv"
	"time"

	"github.com/belfinor/lrand"
)

type Uniq struct {
	next   chan string
	cancel context.CancelFunc
}

func New(pref ...string) *Uniq {
	obj := &Uniq{
		next: make(chan string, 100),
	}

	ctx, cancel := context.WithCancel(context.Background())

	obj.cancel = cancel

	go maker(ctx, obj.next)

	return obj
}

func maker(ctx context.Context, stream chan string) {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = strconv.FormatInt(time.Now().Unix(), 10)
	}

	hc := crc32.ChecksumIEEE([]byte(hostname))

	tact := lrand.Next() & 0xffff

	calc := func() string {

		tm := time.Now()

		mod := tm.UnixNano() & 0xffff

		str := fmt.Sprintf("f%016x%08x%08x%04x%04x", lrand.Next(), tm.Unix(), hc, tact, mod) // 41

		tact = (tact + 1) & 0xffff

		return fmt.Sprintf("%s%08x", str, crc32.ChecksumIEEE([]byte(str))) // 49
	}

	for {
		select {
		case stream <- calc():
		case <-ctx.Done():
			close(stream)
			return
		}
	}
}

func (u *Uniq) Next() string {
	return <-u.next
}

func (u *Uniq) Close() {
	u.cancel()
}

func (u *Uniq) Check(val string, fullCheck bool) bool {
	waitLen := 49

	if len(val) != waitLen {
		return false
	}

	thex := val[17:25]
	ts, e := strconv.ParseInt(thex, 16, 64)
	if e != nil {
		return false
	}

	if ts > time.Now().Unix()+300 {
		return false
	}

	if fullCheck {
		sig := val[waitLen-8:]
		waitSig := fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(val[:waitLen-8])))
		return sig == waitSig
	}

	return true
}
