package txnno

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_TRANS         = 20
	MAX_NANO          = 10000
	MAX_RAND          = 9000
	MAX_NANO_DIGIT    = 10000000000
	MAX_NANO_DIGIT_1  = 100000000000
	MAX_TRANS_DIGIT   = 13
	MAX_TRANS_12DIGIT = 12
)

func getSeed(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func Generate() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS, et)
}

func Generate13Digit() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO_DIGIT
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS_DIGIT, et)
}

func Generate12Digit() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO_DIGIT_1
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS_12DIGIT, et)
}

func Generate10Digit() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO_DIGIT
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS_DIGIT, et)
}
