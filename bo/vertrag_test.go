package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsart"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsstatus"
	"time"
)

// Test_Vertragspartner_Deserialization deserializes an Vertrag json
func (s *Suite) Test_Vertrag_Deserialization() {
	var contract = Vertrag{
		BusinessObject: BusinessObject{
			BoTyp:             botyp.Vertrag,
			VersionStruktur:   "2",
			ExterneReferenzen: nil,
		},
		Vertragsnummer: "",
		Beschreibung:   "",
		Vertragsstatus: vertragsstatus.Abgelehnt,
		Vertragsart:    vertragsart.Buendelvertrag,
		Sparte:         sparte.Strom,
		Vertragsbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Vertragsende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		Vertragspartner1: Geschaeftspartner{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Geschaeftspartner,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Name1:                "Mustermann",
			Anrede:               anrede.Herr,
			Gewerbekennzeichnung: false,
			Partneradresse: com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "28A",
				Landescode:   landescode.DE,
			},
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.Dienstleister,
			},
		},
		Vertragspartner2: Geschaeftspartner{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Geschaeftspartner,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Anrede:               anrede.Frau,
			Name1:                "Musterfrau",
			Gewerbekennzeichnung: false,
			Partneradresse: com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.Dienstleister,
			},
		},
		UnterzeichnerVp1:    nil,
		UnterzeichnerVp2:    nil,
		Vertragskonditionen: nil,
		Vertragsteile: []com.Vertragsteil{
			{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "DE0123456789012345678901234567890",
			},
		},
	}
	serializedContract, err := json.Marshal(contract)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedContract, is.Not(is.Nil()))
	var deserializedVertrag Vertrag
	err = json.Unmarshal(serializedContract, &deserializedVertrag)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVertrag, is.EqualTo(contract))
}

// TestFailedGVertragValidation verifies that the validators of a Vertrag work
func (s *Suite) TestFailedVertragValidation() {
	validate := validator.New()
	invalidVertrags := map[string][]interface{}{
		"required": {
			Vertrag{},
		},
		"gtfield": {
			Vertrag{
				Vertragsbeginn: time.Now(),
				Vertragsende:   time.Now().Add(time.Hour * -12),
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertrags)
}

//  Test_Successful_Vertrag_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Vertrag_Validation() {

	validate := validator.New()
	validVertrag := []interface{}{
		Vertrag{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Vertrag,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Vertragsnummer: "asd",
			Beschreibung:   "",
			Vertragsstatus: vertragsstatus.Angenommen,
			Vertragsart:    vertragsart.Buendelvertrag,
			Sparte:         sparte.Strom,
			Vertragsbeginn: time.Now(),
			Vertragsende:   time.Now().Add(time.Hour * 24),
			Vertragspartner1: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Geschaeftspartner,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Name1:                "Musterfrau",
				Gewerbekennzeichnung: false,
				Partneradresse: com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.Dienstleister,
				},
			},
			Vertragspartner2: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Geschaeftspartner,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Name1:                "Musterfrau",
				Gewerbekennzeichnung: false,
				Partneradresse: com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.Dienstleister,
				},
			},
			UnterzeichnerVp1:    nil,
			UnterzeichnerVp2:    nil,
			Vertragskonditionen: nil,
			Vertragsteile: []com.Vertragsteil{
				{
					Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
					Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
					Lokation:           "DE0123456789012345678901234567890",
				},
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validVertrag)
}