package php

// IsNumeric is meant to be a Go equivalent of the PHP function: is_numeric().
//
// https://www.php.net/is_numeric
func IsNumeric(x interface{}) bool {

	// Initialize
	result := false

	// Figure out result
	switch x.(type) {
	default:
		result = false

	case uint8:
		result = true
	case uint16:
		result = true
	case uint32:
		result = true
	case uint64:
		result = true

	case int8:
		result = true
	case int16:
		result = true
	case int32:
		result = true
	case int64:
		result = true

	case float32:
		result = true
	case float64:
		result = true

	case complex64:
		result = true
	case complex128:
		result = true

		//case byte:
		//	result = true
		//case rune:
		//	result = true

	case uint:
		result = true
	case int:
		result = true

	case string:
		if xAsString, ok := x.(string); ok {
			result = isStringNumeric(xAsString)
		} else {
			result = false
		}
	}

	// Return.
	return result
}

func isStringNumeric(x string) bool {

	result := true

	isFirstLoop := true

	hasPeriod := false

	for i, c := range x {

		switch c {
		default:
			result = false
			return result

		case '-':
			if isFirstLoop {
				// Nothing here.
			} else {
				result = false
				return result
			}

		case '.':
			if hasPeriod {
				result = false
				return result
			}

			if isFirstLoop {
				result = false
				return result
			}

			if len(x) == 1+i {
				result = false
				return result
			}

			hasPeriod = true

		case '0': // Nothing here.
		case '1': // Nothing here.
		case '2': // Nothing here.
		case '3': // Nothing here.
		case '4': // Nothing here.
		case '5': // Nothing here.
		case '6': // Nothing here.
		case '7': // Nothing here.
		case '8': // Nothing here.
		case '9': // Nothing here.
		} // switch

		isFirstLoop = false

	} // for

	return result
}
