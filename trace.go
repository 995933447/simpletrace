package simpletrace

import (
	"github.com/995933447/stringhelper-go"
)

func NewSpanId() string {
	return stringhelper.GenRandomStr(6, stringhelper.RandomStringModNumberPlusLetter)
}

func NewTraceId() string {
	return stringhelper.GenRandomStr(20, stringhelper.RandomStringModNumberPlusLetter)
}
