package php



import (
	"strconv"
)



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
					result = isStringNumeric( xAsString )
				} else {
					result = false
				}
		} // switch


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
				result = false; return result

			case '-':
				if isFirstLoop {
					// Nothing here.
				} else {
					result = false; return result
				}

			case '.':
				if hasPeriod {
					result = false; return result
				}

				if isFirstLoop {
					result = false; return result
				}

				if len(x) == 1+i {
					result = false; return result
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



func Serialize(x interface{}) string {

	// Initialize.
		serialized := ""

		switch x.(type) {
			default:
				serialized = ""

			case uint8:
				serialized = "i:"+ strconv.FormatUint(uint64(x.(uint8)), 10) +";"
			case uint16:
				serialized = "i:"+ strconv.FormatUint(uint64(x.(uint16)), 10) +";"
			case uint32:
				serialized = "i:"+ strconv.FormatUint(uint64(x.(uint32)), 10) +";"
			case uint64:
				serialized = "i:"+ strconv.FormatUint(uint64(x.(uint64)), 10) +";"

			case int8:
				serialized = "i:"+ strconv.FormatInt(int64(x.(int8)), 10) +";"
			case int16:
				serialized = "i:"+ strconv.FormatInt(int64(x.(int16)), 10) +";"
			case int32:
				serialized = "i:"+ strconv.FormatInt(int64(x.(int32)), 10) +";"
			case int64:
				serialized = "i:"+ strconv.FormatInt(int64(x.(int64)), 10) +";"

			case float32:
				serialized = ""
			case float64:
				serialized = ""

			case complex64:
				serialized = ""
			case complex128:
				serialized = ""

			//case byte:
			//	serialized = "i:"+ strconv.FormatInt(x.(byte), 10) +";"
			//case rune:
			//	serialized = "i:"+ strconv.FormatInt(x.(rune), 10) +";"

			case uint:
				serialized = "i:"+ strconv.FormatUint(uint64(x.(uint)), 10) +";"
			case int:
				serialized = "i:"+ strconv.FormatInt(int64(x.(int)), 10) +";"

			case string:
				serialized = "s:"+ strconv.FormatInt(int64(len(x.(string))), 10) +":\""+ x.(string) +"\";"

//@TODO: Add support for maps, slices and arrays
		} // switch


	// Return.
		return serialized	
}
