package php

// IsNumeric is meant to be a Go equivalent of the PHP function: is_numeric().
//
// https://www.php.net/is_numeric
func IsNumeric(x interface{}) bool {

	// Figure out result
	switch casted := x.(type) {
	case
		complex64,complex128,
		float32,float64,
		int,int8,int16,int32,int64,
		uint,uint8,uint16,uint32,uint64:
		return true
	case string:
		return isStringNumeric(casted)
	default:
		return false
	}
}
