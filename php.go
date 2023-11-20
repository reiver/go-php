package php

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
