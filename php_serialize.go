package php



import (
	"bytes"
	"strconv"
)



func Serialize(x interface{}) string {

	// Initialize.
		var serialized bytes.Buffer

		switch x.(type) {
			case uint8:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatUint(uint64(x.(uint8)), 10)   )
				serialized.WriteString(";")
			case uint16:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatUint(uint64(x.(uint16)), 10)   )
				serialized.WriteString(";")
			case uint32:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatUint(uint64(x.(uint32)), 10)   )
				serialized.WriteString(";")
			case uint64:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatUint(uint64(x.(uint64)), 10)   )
				serialized.WriteString(";")

			case int8:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatInt(int64(x.(int8)), 10)   )
				serialized.WriteString(";")
			case int16:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatInt(int64(x.(int16)), 10)   )
				serialized.WriteString(";")
			case int32:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatInt(int64(x.(int32)), 10)   )
				serialized.WriteString(";")
			case int64:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatInt(int64(x.(int64)), 10)   )
				serialized.WriteString(";")

			case float32:
				serialized.WriteString(   serializeFloat32(x.(float32))   )
			case float64:
				serialized.WriteString(   serializeFloat64(x.(float64))   )

			case complex64:
				serialized.WriteString("a:2:{s:4:\"real\";")
				serialized.WriteString(   serializeFloat32(real(x.(complex64)))   )
				serialized.WriteString("s:4:\"imag\";")
				serialized.WriteString(   serializeFloat32(imag(x.(complex64)))   )
				serialized.WriteString("}")
			case complex128:
				serialized.WriteString("a:2:{s:4:\"real\";")
				serialized.WriteString(   serializeFloat64(real(x.(complex128)))   )
				serialized.WriteString("s:4:\"imag\";")
				serialized.WriteString(   serializeFloat64(imag(x.(complex128)))   )
				serialized.WriteString("}")

			//case byte:
			//	serialized.WriteString("i:")
			//	serialized.WriteString(   strconv.FormatInt(int64(x.(byte)), 10)   )
			//	serialized.WriteString(";")
			//case rune:
			//	serialized.WriteString("i:")
			//	serialized.WriteString(   strconv.FormatInt(int64(x.(rune)), 10)   )
			//	serialized.WriteString(";")

			case uint:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatUint(uint64(x.(uint)), 10)   )
				serialized.WriteString(";")
			case int:
				serialized.WriteString("i:")
				serialized.WriteString(   strconv.FormatInt(int64(x.(int)), 10)   )
				serialized.WriteString(";")

			case string:
				serialized.WriteString("s:")
				serialized.WriteString(   strconv.FormatInt(int64(len(x.(string))), 10)   )
				serialized.WriteString(":\"")
				serialized.WriteString(   x.(string)   )
				serialized.WriteString("\";")


			case []interface{}:
				serialized.WriteString("a:")
				serialized.WriteString(   strconv.FormatInt(int64(len(x.([]interface{}))), 10)   )
				serialized.WriteString(":{")
				for i := 0; i < len(x.([]interface{})); i++ {
					serialized.WriteString(   Serialize(i)   )
					serialized.WriteString(   Serialize(x.([]interface{})[i])   )
				}
				serialized.WriteString("}")

			case map[string]interface{}:
				serialized.WriteString("a:")
				serialized.WriteString(   strconv.FormatInt(int64(len(x.(map[string]interface{}))), 10)   )
				serialized.WriteString(":{")
				for key, value := range x.(map[string]interface{}) {
					serialized.WriteString(   Serialize(key)   )
					serialized.WriteString(   Serialize(value)   )
				}
				serialized.WriteString("}")

			default:
				serialized.WriteString("")


//@TODO: Add support for arrays
		} // switch


	// Return.
		return serialized.String()
}

func serializeFloat32(x float32) string {
//@TODO
	return ""
}

func serializeFloat64(x float64) string {
//@TODO
	return ""
}
