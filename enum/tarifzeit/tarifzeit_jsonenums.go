// Code generated by jsonenums --type Tarifzeit; DO NOT EDIT.

package tarifzeit

import (
	"encoding/json"
	"fmt"
)

var (
	_TarifzeitNameToValue = map[string]Tarifzeit{
		"TZ_STANDARD": TZ_STANDARD,
		"TZ_HT":       TZ_HT,
		"TZ_NT":       TZ_NT,
	}

	_TarifzeitValueToName = map[Tarifzeit]string{
		TZ_STANDARD: "TZ_STANDARD",
		TZ_HT:       "TZ_HT",
		TZ_NT:       "TZ_NT",
	}
)

func init() {
	var v Tarifzeit
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TarifzeitNameToValue = map[string]Tarifzeit{
			interface{}(TZ_STANDARD).(fmt.Stringer).String(): TZ_STANDARD,
			interface{}(TZ_HT).(fmt.Stringer).String():       TZ_HT,
			interface{}(TZ_NT).(fmt.Stringer).String():       TZ_NT,
		}
	}
}

// MarshalJSON is generated so Tarifzeit satisfies json.Marshaler.
func (r Tarifzeit) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TarifzeitValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Tarifzeit: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Tarifzeit satisfies json.Unmarshaler.
func (r *Tarifzeit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Tarifzeit should be a string, got %s", data)
	}
	v, ok := _TarifzeitNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Tarifzeit %q", s)
	}
	*r = v
	return nil
}
