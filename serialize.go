package php

import (
	"bytes"
	"sort"
	"strconv"
)

type myByteBuffer struct {
	bytes.Buffer
}

func Serialize(x interface{}) string {

	// Initialize.
	var serialized myByteBuffer

	// Serialize.
	serialized.serialize(x)

	// Return.
	return serialized.String()
}

func (serialized *myByteBuffer) serialize(x interface{}) {

	// Serialize
	switch xx := x.(type) {
	case uint8:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatUint() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatUint(uint64(x.(uint8)), 10))
		serialized.WriteString(";")
	case uint16:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatUint() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatUint(uint64(x.(uint16)), 10))
		serialized.WriteString(";")
	case uint32:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatUint() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatUint(uint64(x.(uint32)), 10))
		serialized.WriteString(";")
	case uint64:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatUint() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatUint(uint64(x.(uint64)), 10))
		serialized.WriteString(";")

	case int8:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(x.(int8)), 10))
		serialized.WriteString(";")
	case int16:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(x.(int16)), 10))
		serialized.WriteString(";")
	case int32:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(x.(int32)), 10))
		serialized.WriteString(";")
	case int64:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(x.(int64)), 10))
		serialized.WriteString(";")

	case float32:
		serialized.writeSerializedFloat32(x.(float32))
	case float64:
		serialized.writeSerializedFloat64(x.(float64))

	case complex64:
		serialized.WriteString("a:2:{s:4:\"real\";")
		serialized.writeSerializedFloat32(real(x.(complex64)))
		serialized.WriteString("s:4:\"imag\";")
		serialized.writeSerializedFloat32(imag(x.(complex64)))
		serialized.WriteString("}")
	case complex128:
		serialized.WriteString("a:2:{s:4:\"real\";")
		serialized.writeSerializedFloat64(real(x.(complex128)))
		serialized.WriteString("s:4:\"imag\";")
		serialized.writeSerializedFloat64(imag(x.(complex128)))
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
		//@TODO: Need a version of strconv.FormatUint() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatUint(uint64(x.(uint)), 10))
		serialized.WriteString(";")
	case int:
		serialized.WriteString("i:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(x.(int)), 10))
		serialized.WriteString(";")

	case string:
		serialized.WriteString("s:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(len(x.(string))), 10))
		serialized.WriteString(":\"")
		serialized.WriteString(x.(string))
		serialized.WriteString("\";")

	case []interface{}:
		serialized.WriteString("a:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(len(x.([]interface{}))), 10))
		serialized.WriteString(":{")
		for i := 0; i < len(x.([]interface{})); i++ {
			serialized.serialize(i)
			serialized.serialize(x.([]interface{})[i])
		}
		serialized.WriteString("}")

	case map[string]interface{}:

		// NOTE that this will serialize the map in a deterministic way. Golang normally gives you
		// the keys of the map in a random order. The code here gets the keys of the map, and
		// orders then (in alphabetical order) so that the serialized form of the map is
		// deterministic. (This also means that serializing maps is a bit slower than need be,
		// and may use a bit more memory than need be. It's a trade off and a choice to do it
		// this way.)

		mapKeys := make([]string, len(xx))
		i := 0
		for k, _ := range xx {
			mapKeys[i] = k
			i++
		} // for
		sort.Strings(mapKeys)

		serialized.WriteString("a:")
		//@TODO: Need a version of strconv.FormatInt() that writes to a buffer to make it more efficient.
		serialized.WriteString(strconv.FormatInt(int64(len(xx)), 10))
		serialized.WriteString(":{")
		for i = 0; i < len(xx); i++ {
			theMapKey := mapKeys[i]

			serialized.serialize(theMapKey)
			serialized.serialize(xx[theMapKey])
		}
		serialized.WriteString("}")

	default:
		serialized.WriteString("")

		//@TODO: Add support for arrays
	} // switch
}

func (me *myByteBuffer) writeSerializedFloat32(x float32) {
	//@TODO
	me.WriteString("")
}

func (me *myByteBuffer) writeSerializedFloat64(x float64) {
	//@TODO
	me.WriteString("")
}
