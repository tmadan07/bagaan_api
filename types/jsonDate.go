package types

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

// First create a type alias
type JsonDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (u *JsonDate) Scan(value interface{}) error {
	*u = JsonDate(value.(time.Time))
	return nil
}

// func (u *JsonDate) AddDate(value interface{}) error {
// 	*u = JsonDate(value.(time.Time))
// 	return nil
// }

func (u JsonDate) Value() (driver.Value, error) {
	return time.Time(u), nil
}
