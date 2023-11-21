package php

import (
	"testing"
)

func TestSerializeString(t *testing.T) {

	var serialized string
	var expected string

	var x string

	x = "apple"
	expected = "s:5:\"apple\";"

	if serialized = Serialize(x); expected != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}

	x = "123"
	expected = "s:3:\"123\";"

	if serialized = Serialize(x); expected != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}

	x = "-3"
	expected = "s:2:\"-3\";"

	if serialized = Serialize(x); expected != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}

	x = "12.34"
	expected = "s:5:\"12.34\";"

	if serialized = Serialize(x); expected != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}

	x = "-12.34"
	expected = "s:6:\"-12.34\";"

	if serialized = Serialize(x); expected != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}

	x = "apple \"ok\" banana"
	expected = "s:17:\"apple \"ok\" banana\";"

	if serialized = Serialize(x); expected != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}
}

func TestSerializeMapWithStringKeys(t *testing.T) {

	var serialized string
	var expected1, expected2, expected3, expected4, expected5, expected6 string

	var x map[string]interface{}

	x = map[string]interface{}{"one": "a", "two": "bb", "three": "ccc"}
	expected1 = "a:3:{s:3:\"one\";s:1:\"a\";s:3:\"two\";s:2:\"bb\";s:5:\"three\";s:3:\"ccc\";}"
	expected2 = "a:3:{s:3:\"one\";s:1:\"a\";s:5:\"three\";s:3:\"ccc\";s:3:\"two\";s:2:\"bb\";}"
	expected3 = "a:3:{s:3:\"two\";s:2:\"bb\";s:3:\"one\";s:1:\"a\";s:5:\"three\";s:3:\"ccc\";}"
	expected4 = "a:3:{s:5:\"three\";s:3:\"ccc\";s:3:\"one\";s:1:\"a\";s:3:\"two\";s:2:\"bb\";}"
	expected5 = "a:3:{s:3:\"two\";s:2:\"bb\";s:5:\"three\";s:3:\"ccc\";s:3:\"one\";s:1:\"a\";}"
	expected6 = "a:3:{s:5:\"three\";s:3:\"ccc\";s:3:\"two\";s:2:\"bb\";s:3:\"one\";s:1:\"a\";}"

	if serialized = Serialize(x); expected1 != serialized && expected2 != serialized && expected3 != serialized && expected4 != serialized && expected5 != serialized && expected6 != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v] | [%#v] | [%#v] | [%#v] | [%#v] | [%#v]. Instead got [%#v].", x, expected1, expected2, expected3, expected4, expected5, expected6, serialized)
	}

	x = map[string]interface{}{"apple": "banana", "cherry": "watermelon"}
	expected1 = "a:2:{s:5:\"apple\";s:6:\"banana\";s:6:\"cherry\";s:10:\"watermelon\";}"
	expected2 = "a:2:{s:6:\"cherry\";s:10:\"watermelon\";s:5:\"apple\";s:6:\"banana\";}"

	if serialized = Serialize(x); expected1 != serialized && expected2 != serialized {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v] | [%#v]. Instead got [%#v].", x, expected1, expected2, serialized)
	}
}

func TestSerializeArray(t *testing.T) {

	var serialized string
	var expected string

	var x [3]string

	x = [3]string{"a", "bb", "ccc"}
	expected = "a:3:{i:0;s:1:\"a\";i:1;s:2:\"bb\";i:2;s:3:\"ccc\";}"

	if serialized = Serialize(x); expected != serialized {
		//@TODO
		//		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}
}
