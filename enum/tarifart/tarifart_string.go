// Code generated by "stringer --type Tarifart"; DO NOT EDIT.

package tarifart

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Eintarif-1]
	_ = x[Zweitarif-2]
	_ = x[Mehrtarif-3]
	_ = x[SmartMeter-4]
	_ = x[Leistungsgemessen-5]
}

const _Tarifart_name = "EintarifZweitarifMehrtarifSmartMeterLeistungsgemessen"

var _Tarifart_index = [...]uint8{0, 8, 17, 26, 36, 53}

func (i Tarifart) String() string {
	i -= 1
	if i < 0 || i >= Tarifart(len(_Tarifart_index)-1) {
		return "Tarifart(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Tarifart_name[_Tarifart_index[i]:_Tarifart_index[i+1]]
}