package php



import (
	"strconv"
)



func Serialize(x interface{}) string {

	// Initialize.
		serialized := ""

		switch x.(type) {
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
				serialized = serializeFloat32(x.(float32))
			case float64:
				serialized = serializeFloat64(x.(float64))

			case complex64:
				serialized = "a:2:{s:4:\"real\";"+ serializeFloat32(real(x.(complex64))) +"s:4:\"imag\";"+ serializeFloat32(imag(x.(complex64))) +"}"
			case complex128:
				serialized = "a:2:{s:4:\"real\";"+ serializeFloat64(real(x.(complex128))) +"s:4:\"imag\";"+ serializeFloat64(imag(x.(complex128))) +"}"

			//case byte:
			//	serialized = "i:"+ strconv.FormatInt(int64(x.(byte)), 10) +";"
			//case rune:
			//	serialized = "i:"+ strconv.FormatInt(int64(x.(rune)), 10) +";"

			case uint:
				serialized = "i:"+ strconv.FormatUint(uint64(x.(uint)), 10) +";"
			case int:
				serialized = "i:"+ strconv.FormatInt(int64(x.(int)), 10) +";"

			case string:
				serialized = "s:"+ strconv.FormatInt(int64(len(x.(string))), 10) +":\""+ x.(string) +"\";"


			case []interface{}:
				serialized = "a:"+ strconv.FormatInt(int64(len(x.([]interface{}))), 10) +":{"
				for i := 0; i < len(x.([]interface{})); i++ {
					serialized += Serialize(i)
					serialized += Serialize(x.([]interface{})[i])
				}
				serialized += "}"

			case map[string]interface{}:
				serialized = "a:"+ strconv.FormatInt(int64(len(x.(map[string]interface{}))), 10) +":{"
				for key, value := range x.(map[string]interface{}) {
					serialized += Serialize(key)
					serialized += Serialize(value)
				}
				serialized += "}"

			default:
				serialized = ""


//@TODO: Add support for arrays
		} // switch


	// Return.
		return serialized	
}

func serializeFloat32(x float32) string {
//@TODO
	return ""
}

func serializeFloat64(x float64) string {
//@TODO
	return ""
}
