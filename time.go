package httpcord

import (
	"fmt"
	"time"
)

type (
	Time struct {
		time.Time
	}
	TimestampStyle string
)

const (
	ShortTimeStyle     TimestampStyle = "t"
	LongTimeStyle      TimestampStyle = "T"
	ShortDateStyle     TimestampStyle = "d"
	LongDateStyle      TimestampStyle = "D"
	ShortDateTimeStyle TimestampStyle = "f"
	LongDateTimeStyle  TimestampStyle = "F"
	RelativeStyle      TimestampStyle = "R"
)

func (t *Time) Format(style TimestampStyle) string {
	return fmt.Sprintf("<t:%d:%s>", t.Unix(), style)
}

func (t Time) ShortTime() string {
	return t.Format(ShortTimeStyle)
}

func (t Time) LongTime() string {
	return t.Format(LongTimeStyle)
}

func (t Time) ShortDate() string {
	return t.Format(ShortDateStyle)
}

func (t Time) LongDate() string {
	return t.Format(LongDateStyle)
}

func (t Time) ShortDateTime() string {
	return t.Format(ShortDateTimeStyle)
}

func (t Time) LongDateTime() string {
	return t.Format(LongDateTimeStyle)
}

func (t Time) Relative() string {
	return t.Format(RelativeStyle)
}
