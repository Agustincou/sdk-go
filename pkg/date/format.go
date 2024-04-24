package date

import (
	"time"
)

const (
	apiTimeFormat = "2006-01-02T15:04:05.000Z07:00"
)

type ApiTimeFormat struct {
	time.Time
}

func (ct ApiTimeFormat) String() string {
	return ct.Format(apiTimeFormat)
}

func (ct *ApiTimeFormat) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(apiTimeFormat, string(data))
	if err != nil {
		return err
	}
	ct.Time = t

	return nil
}

func (ct ApiTimeFormat) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Format(apiTimeFormat) + `"`), nil
}
