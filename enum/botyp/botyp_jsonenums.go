// Code generated by jsonenums --type BOTyp; DO NOT EDIT.

package botyp

import (
	"encoding/json"
	"fmt"
)

var (
	_BOTypNameToValue = map[string]BOTyp{
		"Angebot":                     Angebot,
		"Ansprechpartner":             Ansprechpartner,
		"Ausschreibung":               Ausschreibung,
		"Energiemenge":                Energiemenge,
		"Geschaeftspartner":           Geschaeftspartner,
		"Kosten":                      Kosten,
		"Lastgang":                    Lastgang,
		"Marktlokation":               Marktlokation,
		"Messlokation":                Messlokation,
		"Marktteilnehmer":             Marktteilnehmer,
		"Netznutzungsrechnung":        Netznutzungsrechnung,
		"Preisblatt":                  Preisblatt,
		"PreisblattDienstleistung":    PreisblattDienstleistung,
		"PreisblattKonzessionsabgabe": PreisblattKonzessionsabgabe,
		"PreisblattMessung":           PreisblattMessung,
		"PreisblattUmlagen":           PreisblattUmlagen,
		"Rechnung":                    Rechnung,
		"Tarifinfo":                   Tarifinfo,
		"TarifPreisblatt":             TarifPreisblatt,
		"Vertrag":                     Vertrag,
		"Zaehler":                     Zaehler,
		"Zeitreihe":                   Zeitreihe,
	}

	_BOTypValueToName = map[BOTyp]string{
		Angebot:                     "Angebot",
		Ansprechpartner:             "Ansprechpartner",
		Ausschreibung:               "Ausschreibung",
		Energiemenge:                "Energiemenge",
		Geschaeftspartner:           "Geschaeftspartner",
		Kosten:                      "Kosten",
		Lastgang:                    "Lastgang",
		Marktlokation:               "Marktlokation",
		Messlokation:                "Messlokation",
		Marktteilnehmer:             "Marktteilnehmer",
		Netznutzungsrechnung:        "Netznutzungsrechnung",
		Preisblatt:                  "Preisblatt",
		PreisblattDienstleistung:    "PreisblattDienstleistung",
		PreisblattKonzessionsabgabe: "PreisblattKonzessionsabgabe",
		PreisblattMessung:           "PreisblattMessung",
		PreisblattUmlagen:           "PreisblattUmlagen",
		Rechnung:                    "Rechnung",
		Tarifinfo:                   "Tarifinfo",
		TarifPreisblatt:             "TarifPreisblatt",
		Vertrag:                     "Vertrag",
		Zaehler:                     "Zaehler",
		Zeitreihe:                   "Zeitreihe",
	}
)

func init() {
	var v BOTyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BOTypNameToValue = map[string]BOTyp{
			interface{}(Angebot).(fmt.Stringer).String():                     Angebot,
			interface{}(Ansprechpartner).(fmt.Stringer).String():             Ansprechpartner,
			interface{}(Ausschreibung).(fmt.Stringer).String():               Ausschreibung,
			interface{}(Energiemenge).(fmt.Stringer).String():                Energiemenge,
			interface{}(Geschaeftspartner).(fmt.Stringer).String():           Geschaeftspartner,
			interface{}(Kosten).(fmt.Stringer).String():                      Kosten,
			interface{}(Lastgang).(fmt.Stringer).String():                    Lastgang,
			interface{}(Marktlokation).(fmt.Stringer).String():               Marktlokation,
			interface{}(Messlokation).(fmt.Stringer).String():                Messlokation,
			interface{}(Marktteilnehmer).(fmt.Stringer).String():             Marktteilnehmer,
			interface{}(Netznutzungsrechnung).(fmt.Stringer).String():        Netznutzungsrechnung,
			interface{}(Preisblatt).(fmt.Stringer).String():                  Preisblatt,
			interface{}(PreisblattDienstleistung).(fmt.Stringer).String():    PreisblattDienstleistung,
			interface{}(PreisblattKonzessionsabgabe).(fmt.Stringer).String(): PreisblattKonzessionsabgabe,
			interface{}(PreisblattMessung).(fmt.Stringer).String():           PreisblattMessung,
			interface{}(PreisblattUmlagen).(fmt.Stringer).String():           PreisblattUmlagen,
			interface{}(Rechnung).(fmt.Stringer).String():                    Rechnung,
			interface{}(Tarifinfo).(fmt.Stringer).String():                   Tarifinfo,
			interface{}(TarifPreisblatt).(fmt.Stringer).String():             TarifPreisblatt,
			interface{}(Vertrag).(fmt.Stringer).String():                     Vertrag,
			interface{}(Zaehler).(fmt.Stringer).String():                     Zaehler,
			interface{}(Zeitreihe).(fmt.Stringer).String():                   Zeitreihe,
		}
	}
}

// MarshalJSON is generated so BOTyp satisfies json.Marshaler.
func (r BOTyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BOTypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BOTyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BOTyp satisfies json.Unmarshaler.
func (r *BOTyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BOTyp should be a string, got %s", data)
	}
	v, ok := _BOTypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BOTyp %q", s)
	}
	*r = v
	return nil
}