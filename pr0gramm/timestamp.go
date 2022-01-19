package pr0gramm

import (
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (ct *Timestamp) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	secs, err := strconv.ParseInt(string(b), 10, 0)
	if err != nil {
		return err
	}

	ct.Time = time.Unix(secs, 0)
	return
}

func (ct Timestamp) MarshalJSON() ([]byte, error) {
	formatted := strconv.FormatInt(int64(ct.Time.Unix()), 10)
	return []byte(formatted), nil
}

func (ct Timestamp) MarshalText() (text []byte, err error) {
	value := strconv.Itoa(int(ct.Unix()))
	return []byte(value), nil
}
