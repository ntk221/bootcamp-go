package object

import (
	"database/sql/driver"
	"time"
)

// DateTime
// Wrapper of time.Time to implement custom method for JSON/DB interface
type DateTime struct{ time.Time }

const timeFormat = "2006-01-02T15:04:05Z07:00"

func (t DateTime) format() string {
	return t.Format(timeFormat)
}

// MarshalJSON
// encoding/json/Marshaler
func (t *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}

// UnmarshalJSON
// encoding/json/Unmarshaler
func (t *DateTime) UnmarshalJSON(b []byte) error {
	t.Time, _ = time.Parse(`"`+timeFormat+`"`, string(b))
	return nil
}

// Value
// database/sql/driver/Valuer
func (t DateTime) Value() (driver.Value, error) {
	return t.Time, nil

}

// Scan
// database/sql/driver/Valuer
func (t *DateTime) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	return nil
}
