package simpletrace

import (
	"wegod/internal/pkg/stringhelper"
)

func NewSpanId() string {
	return stringhelper.GenRandomStr(6, stringhelper.RandomStringModNumberPlusLetter)
}

func NewTraceId() string {
	return stringhelper.GenRandomStr(20, stringhelper.RandomStringModNumberPlusLetter)
}
