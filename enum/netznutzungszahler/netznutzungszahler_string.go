// Code generated by "stringer --type Netznutzungszahler"; DO NOT EDIT.

package netznutzungszahler

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KUNDE-1]
	_ = x[LIEFERANT-2]
}

const _Netznutzungszahler_name = "KUNDELIEFERANT"

var _Netznutzungszahler_index = [...]uint8{0, 5, 14}

func (i Netznutzungszahler) String() string {
	i -= 1
	if i < 0 || i >= Netznutzungszahler(len(_Netznutzungszahler_index)-1) {
		return "Netznutzungszahler(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Netznutzungszahler_name[_Netznutzungszahler_index[i]:_Netznutzungszahler_index[i+1]]
}
