package luniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.006
// @date    2019-06-04

import (
	"context"
	"fmt"
	"hash/crc32"
	"strings"
	"time"

	"github.com/belfinor/lrand"
)

var crctab *crc32.Table

func init() {
	crctab = crc32.MakeTable(crc32.IEEE)
}

type Uniq struct {
	next   chan string
	pref   string
	cancel context.CancelFunc
}

func New(pref ...string) *Uniq {
	obj := &Uniq{
		next: make(chan string, 10),
	}

	ctx, cancel := context.WithCancel(context.Background())

	obj.cancel = cancel
	obj.pref = strings.Join(pref, "")

	go maker(ctx, obj.next, obj.pref)

	return obj
}

func maker(ctx context.Context, stream chan string, prefix string) {

	fb1 := int64(1)
	fb2 := int64(1)

	tact := lrand.Next() & 0xffffff

	calc := func() string {
		ts := time.Now()
		mod := ts.UnixNano() & 0xffff
		epoch := ts.Unix()
		str := fmt.Sprintf("%s%016x%08x%06x%04x%06x", prefix, lrand.Next(), epoch, tact, mod, fb1&0xffffffff)
		fb3 := fb1 + fb2
		fb1 = fb2
		fb2 = fb3
		tact = (tact + 1) & 0xffffff
		return fmt.Sprintf("%s%08x", str, crc32.Checksum([]byte(str), crctab))
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
	waitLen := len(u.pref) + 48

	if len(val) != waitLen {
		return false
	}

	if fullCheck {
		sig := val[waitLen-8:]
		waitSig := fmt.Sprintf("%08x", crc32.Checksum([]byte(val[:waitLen-8]), crctab))
		return sig == waitSig
	}

	return true
}
