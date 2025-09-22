package utils

import (
	"editor/app/types"
	"time"
)

type timeNowDefault struct{}

func NewTimeNowDefault() types.TimeNow {
	return &timeNowDefault{}
}

func (t *timeNowDefault) Now() time.Time {
	return time.Now()
}

type timeNowTest struct {
	time time.Time
}

func NewTimeNowTest(timeNow string) types.TimeNow {
	timeParsed, err := time.Parse(time.DateTime, timeNow)
	if err != nil {
		panic(err)
	}
	return &timeNowTest{
		time: timeParsed,
	}
}

func (t *timeNowTest) Now() time.Time {
	return t.time
}
