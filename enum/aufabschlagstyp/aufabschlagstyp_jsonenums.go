// Code generated by jsonenums --type AufAbschlagstyp; DO NOT EDIT.

package aufabschlagstyp

import (
	"encoding/json"
	"fmt"
)

var (
	_AufAbschlagstypNameToValue = map[string]AufAbschlagstyp{
		"RELATIV": RELATIV,
		"ABSOLUT": ABSOLUT,
	}

	_AufAbschlagstypValueToName = map[AufAbschlagstyp]string{
		RELATIV: "RELATIV",
		ABSOLUT: "ABSOLUT",
	}
)

func init() {
	var v AufAbschlagstyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_AufAbschlagstypNameToValue = map[string]AufAbschlagstyp{
			interface{}(RELATIV).(fmt.Stringer).String(): RELATIV,
			interface{}(ABSOLUT).(fmt.Stringer).String(): ABSOLUT,
		}
	}
}

// MarshalJSON is generated so AufAbschlagstyp satisfies json.Marshaler.
func (r AufAbschlagstyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _AufAbschlagstypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid AufAbschlagstyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AufAbschlagstyp satisfies json.Unmarshaler.
func (r *AufAbschlagstyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AufAbschlagstyp should be a string, got %s", data)
	}
	v, ok := _AufAbschlagstypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid AufAbschlagstyp %q", s)
	}
	*r = v
	return nil
}
