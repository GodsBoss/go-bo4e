// Code generated by jsonenums --type Energierichtung; DO NOT EDIT.

package energierichtung

import (
	"encoding/json"
	"fmt"
)

var (
	_EnergierichtungNameToValue = map[string]Energierichtung{
		"Aussp": Aussp,
		"Einsp": Einsp,
	}

	_EnergierichtungValueToName = map[Energierichtung]string{
		Aussp: "Aussp",
		Einsp: "Einsp",
	}
)

func init() {
	var v Energierichtung
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_EnergierichtungNameToValue = map[string]Energierichtung{
			interface{}(Aussp).(fmt.Stringer).String(): Aussp,
			interface{}(Einsp).(fmt.Stringer).String(): Einsp,
		}
	}
}

// MarshalJSON is generated so Energierichtung satisfies json.Marshaler.
func (r Energierichtung) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _EnergierichtungValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Energierichtung: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Energierichtung satisfies json.Unmarshaler.
func (r *Energierichtung) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Energierichtung should be a string, got %s", data)
	}
	v, ok := _EnergierichtungNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Energierichtung %q", s)
	}
	*r = v
	return nil
}