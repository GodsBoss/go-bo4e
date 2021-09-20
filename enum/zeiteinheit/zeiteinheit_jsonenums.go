// Code generated by jsonenums --type Zeiteinheit; DO NOT EDIT.

package zeiteinheit

import (
	"encoding/json"
	"fmt"
)

var (
	_ZeiteinheitNameToValue = map[string]Zeiteinheit{
		"Sekunde":       Sekunde,
		"Minute":        Minute,
		"Stunde":        Stunde,
		"ViertelStunde": ViertelStunde,
		"Tag":           Tag,
		"Woche":         Woche,
		"Monat":         Monat,
		"Quartal":       Quartal,
		"Halbjahr":      Halbjahr,
		"Jahr":          Jahr,
	}

	_ZeiteinheitValueToName = map[Zeiteinheit]string{
		Sekunde:       "Sekunde",
		Minute:        "Minute",
		Stunde:        "Stunde",
		ViertelStunde: "ViertelStunde",
		Tag:           "Tag",
		Woche:         "Woche",
		Monat:         "Monat",
		Quartal:       "Quartal",
		Halbjahr:      "Halbjahr",
		Jahr:          "Jahr",
	}
)

func init() {
	var v Zeiteinheit
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ZeiteinheitNameToValue = map[string]Zeiteinheit{
			interface{}(Sekunde).(fmt.Stringer).String():       Sekunde,
			interface{}(Minute).(fmt.Stringer).String():        Minute,
			interface{}(Stunde).(fmt.Stringer).String():        Stunde,
			interface{}(ViertelStunde).(fmt.Stringer).String(): ViertelStunde,
			interface{}(Tag).(fmt.Stringer).String():           Tag,
			interface{}(Woche).(fmt.Stringer).String():         Woche,
			interface{}(Monat).(fmt.Stringer).String():         Monat,
			interface{}(Quartal).(fmt.Stringer).String():       Quartal,
			interface{}(Halbjahr).(fmt.Stringer).String():      Halbjahr,
			interface{}(Jahr).(fmt.Stringer).String():          Jahr,
		}
	}
}

// MarshalJSON is generated so Zeiteinheit satisfies json.Marshaler.
func (r Zeiteinheit) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ZeiteinheitValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Zeiteinheit: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Zeiteinheit satisfies json.Unmarshaler.
func (r *Zeiteinheit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Zeiteinheit should be a string, got %s", data)
	}
	v, ok := _ZeiteinheitNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Zeiteinheit %q", s)
	}
	*r = v
	return nil
}
