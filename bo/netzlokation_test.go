package bo_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
)

func (s *Suite) Test_Get_NeloId_Checksum() {
	actual := bo.GetNeLoIdCheckSum("E113735592")
	then.AssertThat(s.T(), actual, is.EqualTo(1))
}

func (s *Suite) Test_Get_NeloId_Doesnt_Panic() {
	actual := bo.GetNeLoIdCheckSum("E5345G7F6F")
	then.AssertThat(s.T(), actual, is.EqualTo(0))
}
