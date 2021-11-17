// Code generated by jsonenums --type Themengebiet; DO NOT EDIT.

package themengebiet

import (
	"encoding/json"
	"fmt"
)

var (
	_ThemengebietNameToValue = map[string]Themengebiet{
		"ALLGEMEINER_INFORMATIONSAUSTAUSCH":           ALLGEMEINER_INFORMATIONSAUSTAUSCH,
		"AN_UND_ABMELDUNG":                            AN_UND_ABMELDUNG,
		"ANSPRECHPARTNER_ALLGEMEIN":                   ANSPRECHPARTNER_ALLGEMEIN,
		"ANSPRECHPARTNER_BDEW_DVGW":                   ANSPRECHPARTNER_BDEW_DVGW,
		"ANSPRECHPARTNER_IT_TECHNIK":                  ANSPRECHPARTNER_IT_TECHNIK,
		"BILANZIERUNG":                                BILANZIERUNG,
		"BILANZKREISKOORDINATOR":                      BILANZKREISKOORDINATOR,
		"BILANZKREISVERANTWORTLICHER":                 BILANZKREISVERANTWORTLICHER,
		"DATENFORMATE_ZERTIFIKATE_VERSCHLUESSELUNGEN": DATENFORMATE_ZERTIFIKATE_VERSCHLUESSELUNGEN,
		"DEBITORENMANAGEMENT":                         DEBITORENMANAGEMENT,
		"DEMAND_SIDE_MANAGEMENT":                      DEMAND_SIDE_MANAGEMENT,
		"EDI_VEREINBARUNG":                            EDI_VEREINBARUNG,
		"EDIFACT":                                     EDIFACT,
		"ENERGIEDATENMANAGEMENT":                      ENERGIEDATENMANAGEMENT,
		"FAHRPLANMANAGEMENT":                          FAHRPLANMANAGEMENT,
		"ALOCAT":                                      ALOCAT,
		"APERAK":                                      APERAK,
		"CONTRL":                                      CONTRL,
		"INVOIC":                                      INVOIC,
		"MSCONS":                                      MSCONS,
		"ORDERS":                                      ORDERS,
		"ORDERSP":                                     ORDERSP,
		"REMADV":                                      REMADV,
		"UTILMD":                                      UTILMD,
		"GABI":                                        GABI,
		"GELI":                                        GELI,
		"GERAETERUECKGABE":                            GERAETERUECKGABE,
		"GERAETEWECHSEL":                              GERAETEWECHSEL,
		"GPKE":                                        GPKE,
		"INBETRIEBNAHME":                              INBETRIEBNAHME,
		"KAPAZITAETSMANAGEMENT":                       KAPAZITAETSMANAGEMENT,
		"KLAERFAELLE":                                 KLAERFAELLE,
		"LASTGAENGE_RLM":                              LASTGAENGE_RLM,
		"LIEFERANTENRAHMENVERTRAG":                    LIEFERANTENRAHMENVERTRAG,
		"LIEFERANTENWECHSEL":                          LIEFERANTENWECHSEL,
		"MABIS":                                       MABIS,
		"MAHNWESEN":                                   MAHNWESEN,
		"MARKTGEBIETSVERANTWORTLICHER":                MARKTGEBIETSVERANTWORTLICHER,
		"MARKTKOMMUNIKATION":                          MARKTKOMMUNIKATION,
		"MEHR_MINDERMENGEN":                           MEHR_MINDERMENGEN,
		"MSB_MDL":                                     MSB_MDL,
		"NETZABRECHNUNG":                              NETZABRECHNUNG,
		"NETZENTGELTE":                                NETZENTGELTE,
		"NETZMANAGEMENT":                              NETZMANAGEMENT,
		"RECHT":                                       RECHT,
		"REGULIERUNGSMANAGEMENT":                      REGULIERUNGSMANAGEMENT,
		"REKLAMATIONEN":                               REKLAMATIONEN,
		"SPERREN_ENTSPERREN_INKASSO":                  SPERREN_ENTSPERREN_INKASSO,
		"STAMMDATEN":                                  STAMMDATEN,
		"STOERUNGSFAELLE":                             STOERUNGSFAELLE,
		"TECHNISCHE_FRAGEN":                           TECHNISCHE_FRAGEN,
		"UMSTELLUNG_INVOIC":                           UMSTELLUNG_INVOIC,
		"VERSCHLUESSELUNG_SIGNATUR":                   VERSCHLUESSELUNG_SIGNATUR,
		"VERTRAGSMANAGEMENT":                          VERTRAGSMANAGEMENT,
		"VERTRIEB":                                    VERTRIEB,
		"WIM":                                         WIM,
		"ZAEHLERSTAENDE_SLP":                          ZAEHLERSTAENDE_SLP,
		"ZAHLUNGSVERKEHR":                             ZAHLUNGSVERKEHR,
		"ZUORDNUNGSVEREINBARUNG":                      ZUORDNUNGSVEREINBARUNG,
	}

	_ThemengebietValueToName = map[Themengebiet]string{
		ALLGEMEINER_INFORMATIONSAUSTAUSCH:           "ALLGEMEINER_INFORMATIONSAUSTAUSCH",
		AN_UND_ABMELDUNG:                            "AN_UND_ABMELDUNG",
		ANSPRECHPARTNER_ALLGEMEIN:                   "ANSPRECHPARTNER_ALLGEMEIN",
		ANSPRECHPARTNER_BDEW_DVGW:                   "ANSPRECHPARTNER_BDEW_DVGW",
		ANSPRECHPARTNER_IT_TECHNIK:                  "ANSPRECHPARTNER_IT_TECHNIK",
		BILANZIERUNG:                                "BILANZIERUNG",
		BILANZKREISKOORDINATOR:                      "BILANZKREISKOORDINATOR",
		BILANZKREISVERANTWORTLICHER:                 "BILANZKREISVERANTWORTLICHER",
		DATENFORMATE_ZERTIFIKATE_VERSCHLUESSELUNGEN: "DATENFORMATE_ZERTIFIKATE_VERSCHLUESSELUNGEN",
		DEBITORENMANAGEMENT:                         "DEBITORENMANAGEMENT",
		DEMAND_SIDE_MANAGEMENT:                      "DEMAND_SIDE_MANAGEMENT",
		EDI_VEREINBARUNG:                            "EDI_VEREINBARUNG",
		EDIFACT:                                     "EDIFACT",
		ENERGIEDATENMANAGEMENT:                      "ENERGIEDATENMANAGEMENT",
		FAHRPLANMANAGEMENT:                          "FAHRPLANMANAGEMENT",
		ALOCAT:                                      "ALOCAT",
		APERAK:                                      "APERAK",
		CONTRL:                                      "CONTRL",
		INVOIC:                                      "INVOIC",
		MSCONS:                                      "MSCONS",
		ORDERS:                                      "ORDERS",
		ORDERSP:                                     "ORDERSP",
		REMADV:                                      "REMADV",
		UTILMD:                                      "UTILMD",
		GABI:                                        "GABI",
		GELI:                                        "GELI",
		GERAETERUECKGABE:                            "GERAETERUECKGABE",
		GERAETEWECHSEL:                              "GERAETEWECHSEL",
		GPKE:                                        "GPKE",
		INBETRIEBNAHME:                              "INBETRIEBNAHME",
		KAPAZITAETSMANAGEMENT:                       "KAPAZITAETSMANAGEMENT",
		KLAERFAELLE:                                 "KLAERFAELLE",
		LASTGAENGE_RLM:                              "LASTGAENGE_RLM",
		LIEFERANTENRAHMENVERTRAG:                    "LIEFERANTENRAHMENVERTRAG",
		LIEFERANTENWECHSEL:                          "LIEFERANTENWECHSEL",
		MABIS:                                       "MABIS",
		MAHNWESEN:                                   "MAHNWESEN",
		MARKTGEBIETSVERANTWORTLICHER:                "MARKTGEBIETSVERANTWORTLICHER",
		MARKTKOMMUNIKATION:                          "MARKTKOMMUNIKATION",
		MEHR_MINDERMENGEN:                           "MEHR_MINDERMENGEN",
		MSB_MDL:                                     "MSB_MDL",
		NETZABRECHNUNG:                              "NETZABRECHNUNG",
		NETZENTGELTE:                                "NETZENTGELTE",
		NETZMANAGEMENT:                              "NETZMANAGEMENT",
		RECHT:                                       "RECHT",
		REGULIERUNGSMANAGEMENT:                      "REGULIERUNGSMANAGEMENT",
		REKLAMATIONEN:                               "REKLAMATIONEN",
		SPERREN_ENTSPERREN_INKASSO:                  "SPERREN_ENTSPERREN_INKASSO",
		STAMMDATEN:                                  "STAMMDATEN",
		STOERUNGSFAELLE:                             "STOERUNGSFAELLE",
		TECHNISCHE_FRAGEN:                           "TECHNISCHE_FRAGEN",
		UMSTELLUNG_INVOIC:                           "UMSTELLUNG_INVOIC",
		VERSCHLUESSELUNG_SIGNATUR:                   "VERSCHLUESSELUNG_SIGNATUR",
		VERTRAGSMANAGEMENT:                          "VERTRAGSMANAGEMENT",
		VERTRIEB:                                    "VERTRIEB",
		WIM:                                         "WIM",
		ZAEHLERSTAENDE_SLP:                          "ZAEHLERSTAENDE_SLP",
		ZAHLUNGSVERKEHR:                             "ZAHLUNGSVERKEHR",
		ZUORDNUNGSVEREINBARUNG:                      "ZUORDNUNGSVEREINBARUNG",
	}
)

func init() {
	var v Themengebiet
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ThemengebietNameToValue = map[string]Themengebiet{
			interface{}(ALLGEMEINER_INFORMATIONSAUSTAUSCH).(fmt.Stringer).String():           ALLGEMEINER_INFORMATIONSAUSTAUSCH,
			interface{}(AN_UND_ABMELDUNG).(fmt.Stringer).String():                            AN_UND_ABMELDUNG,
			interface{}(ANSPRECHPARTNER_ALLGEMEIN).(fmt.Stringer).String():                   ANSPRECHPARTNER_ALLGEMEIN,
			interface{}(ANSPRECHPARTNER_BDEW_DVGW).(fmt.Stringer).String():                   ANSPRECHPARTNER_BDEW_DVGW,
			interface{}(ANSPRECHPARTNER_IT_TECHNIK).(fmt.Stringer).String():                  ANSPRECHPARTNER_IT_TECHNIK,
			interface{}(BILANZIERUNG).(fmt.Stringer).String():                                BILANZIERUNG,
			interface{}(BILANZKREISKOORDINATOR).(fmt.Stringer).String():                      BILANZKREISKOORDINATOR,
			interface{}(BILANZKREISVERANTWORTLICHER).(fmt.Stringer).String():                 BILANZKREISVERANTWORTLICHER,
			interface{}(DATENFORMATE_ZERTIFIKATE_VERSCHLUESSELUNGEN).(fmt.Stringer).String(): DATENFORMATE_ZERTIFIKATE_VERSCHLUESSELUNGEN,
			interface{}(DEBITORENMANAGEMENT).(fmt.Stringer).String():                         DEBITORENMANAGEMENT,
			interface{}(DEMAND_SIDE_MANAGEMENT).(fmt.Stringer).String():                      DEMAND_SIDE_MANAGEMENT,
			interface{}(EDI_VEREINBARUNG).(fmt.Stringer).String():                            EDI_VEREINBARUNG,
			interface{}(EDIFACT).(fmt.Stringer).String():                                     EDIFACT,
			interface{}(ENERGIEDATENMANAGEMENT).(fmt.Stringer).String():                      ENERGIEDATENMANAGEMENT,
			interface{}(FAHRPLANMANAGEMENT).(fmt.Stringer).String():                          FAHRPLANMANAGEMENT,
			interface{}(ALOCAT).(fmt.Stringer).String():                                      ALOCAT,
			interface{}(APERAK).(fmt.Stringer).String():                                      APERAK,
			interface{}(CONTRL).(fmt.Stringer).String():                                      CONTRL,
			interface{}(INVOIC).(fmt.Stringer).String():                                      INVOIC,
			interface{}(MSCONS).(fmt.Stringer).String():                                      MSCONS,
			interface{}(ORDERS).(fmt.Stringer).String():                                      ORDERS,
			interface{}(ORDERSP).(fmt.Stringer).String():                                     ORDERSP,
			interface{}(REMADV).(fmt.Stringer).String():                                      REMADV,
			interface{}(UTILMD).(fmt.Stringer).String():                                      UTILMD,
			interface{}(GABI).(fmt.Stringer).String():                                        GABI,
			interface{}(GELI).(fmt.Stringer).String():                                        GELI,
			interface{}(GERAETERUECKGABE).(fmt.Stringer).String():                            GERAETERUECKGABE,
			interface{}(GERAETEWECHSEL).(fmt.Stringer).String():                              GERAETEWECHSEL,
			interface{}(GPKE).(fmt.Stringer).String():                                        GPKE,
			interface{}(INBETRIEBNAHME).(fmt.Stringer).String():                              INBETRIEBNAHME,
			interface{}(KAPAZITAETSMANAGEMENT).(fmt.Stringer).String():                       KAPAZITAETSMANAGEMENT,
			interface{}(KLAERFAELLE).(fmt.Stringer).String():                                 KLAERFAELLE,
			interface{}(LASTGAENGE_RLM).(fmt.Stringer).String():                              LASTGAENGE_RLM,
			interface{}(LIEFERANTENRAHMENVERTRAG).(fmt.Stringer).String():                    LIEFERANTENRAHMENVERTRAG,
			interface{}(LIEFERANTENWECHSEL).(fmt.Stringer).String():                          LIEFERANTENWECHSEL,
			interface{}(MABIS).(fmt.Stringer).String():                                       MABIS,
			interface{}(MAHNWESEN).(fmt.Stringer).String():                                   MAHNWESEN,
			interface{}(MARKTGEBIETSVERANTWORTLICHER).(fmt.Stringer).String():                MARKTGEBIETSVERANTWORTLICHER,
			interface{}(MARKTKOMMUNIKATION).(fmt.Stringer).String():                          MARKTKOMMUNIKATION,
			interface{}(MEHR_MINDERMENGEN).(fmt.Stringer).String():                           MEHR_MINDERMENGEN,
			interface{}(MSB_MDL).(fmt.Stringer).String():                                     MSB_MDL,
			interface{}(NETZABRECHNUNG).(fmt.Stringer).String():                              NETZABRECHNUNG,
			interface{}(NETZENTGELTE).(fmt.Stringer).String():                                NETZENTGELTE,
			interface{}(NETZMANAGEMENT).(fmt.Stringer).String():                              NETZMANAGEMENT,
			interface{}(RECHT).(fmt.Stringer).String():                                       RECHT,
			interface{}(REGULIERUNGSMANAGEMENT).(fmt.Stringer).String():                      REGULIERUNGSMANAGEMENT,
			interface{}(REKLAMATIONEN).(fmt.Stringer).String():                               REKLAMATIONEN,
			interface{}(SPERREN_ENTSPERREN_INKASSO).(fmt.Stringer).String():                  SPERREN_ENTSPERREN_INKASSO,
			interface{}(STAMMDATEN).(fmt.Stringer).String():                                  STAMMDATEN,
			interface{}(STOERUNGSFAELLE).(fmt.Stringer).String():                             STOERUNGSFAELLE,
			interface{}(TECHNISCHE_FRAGEN).(fmt.Stringer).String():                           TECHNISCHE_FRAGEN,
			interface{}(UMSTELLUNG_INVOIC).(fmt.Stringer).String():                           UMSTELLUNG_INVOIC,
			interface{}(VERSCHLUESSELUNG_SIGNATUR).(fmt.Stringer).String():                   VERSCHLUESSELUNG_SIGNATUR,
			interface{}(VERTRAGSMANAGEMENT).(fmt.Stringer).String():                          VERTRAGSMANAGEMENT,
			interface{}(VERTRIEB).(fmt.Stringer).String():                                    VERTRIEB,
			interface{}(WIM).(fmt.Stringer).String():                                         WIM,
			interface{}(ZAEHLERSTAENDE_SLP).(fmt.Stringer).String():                          ZAEHLERSTAENDE_SLP,
			interface{}(ZAHLUNGSVERKEHR).(fmt.Stringer).String():                             ZAHLUNGSVERKEHR,
			interface{}(ZUORDNUNGSVEREINBARUNG).(fmt.Stringer).String():                      ZUORDNUNGSVEREINBARUNG,
		}
	}
}

// MarshalJSON is generated so Themengebiet satisfies json.Marshaler.
func (r Themengebiet) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ThemengebietValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Themengebiet: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Themengebiet satisfies json.Unmarshaler.
func (r *Themengebiet) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Themengebiet should be a string, got %s", data)
	}
	v, ok := _ThemengebietNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Themengebiet %q", s)
	}
	*r = v
	return nil
}