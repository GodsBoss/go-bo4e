// Code generated by jsonenums --type Handelsunstimmigkeitsgrund; DO NOT EDIT.

package handelsunstimmigkeitsgrund

import (
	"encoding/json"
	"fmt"
)

var (
	_HandelsunstimmigkeitsgrundNameToValue = map[string]Handelsunstimmigkeitsgrund{
		"ANMELDUNG_BESTAETIGT":                                 ANMELDUNG_BESTAETIGT,
		"ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN": ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN,
		"ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE":     ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE,
		"NN_MSCONS_UEBERSENDET":                                NN_MSCONS_UEBERSENDET,
		"RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET":         RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET,
		"SONSTIGES_SIEHE_BEGRUENDUNG":                          SONSTIGES_SIEHE_BEGRUENDUNG,
	}

	_HandelsunstimmigkeitsgrundValueToName = map[Handelsunstimmigkeitsgrund]string{
		ANMELDUNG_BESTAETIGT: "ANMELDUNG_BESTAETIGT",
		ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN: "ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN",
		ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE:     "ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE",
		NN_MSCONS_UEBERSENDET:                                "NN_MSCONS_UEBERSENDET",
		RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET:         "RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET",
		SONSTIGES_SIEHE_BEGRUENDUNG:                          "SONSTIGES_SIEHE_BEGRUENDUNG",
	}
)

func init() {
	var v Handelsunstimmigkeitsgrund
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_HandelsunstimmigkeitsgrundNameToValue = map[string]Handelsunstimmigkeitsgrund{
			interface{}(ANMELDUNG_BESTAETIGT).(fmt.Stringer).String():                                 ANMELDUNG_BESTAETIGT,
			interface{}(ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN).(fmt.Stringer).String(): ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN,
			interface{}(ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE).(fmt.Stringer).String():     ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE,
			interface{}(NN_MSCONS_UEBERSENDET).(fmt.Stringer).String():                                NN_MSCONS_UEBERSENDET,
			interface{}(RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET).(fmt.Stringer).String():         RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET,
			interface{}(SONSTIGES_SIEHE_BEGRUENDUNG).(fmt.Stringer).String():                          SONSTIGES_SIEHE_BEGRUENDUNG,
		}
	}
}

// MarshalJSON is generated so Handelsunstimmigkeitsgrund satisfies json.Marshaler.
func (r Handelsunstimmigkeitsgrund) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _HandelsunstimmigkeitsgrundValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Handelsunstimmigkeitsgrund: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Handelsunstimmigkeitsgrund satisfies json.Unmarshaler.
func (r *Handelsunstimmigkeitsgrund) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Handelsunstimmigkeitsgrund should be a string, got %s", data)
	}
	v, ok := _HandelsunstimmigkeitsgrundNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Handelsunstimmigkeitsgrund %q", s)
	}
	*r = v
	return nil
}