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

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}


	x = "123"
	expected = "s:3:\"123\";"

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}


	x = "-3"
	expected = "s:2:\"-3\";"

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}


	x = "12.34"
	expected = "s:5:\"12.34\";"

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}


	x = "-12.34"
	expected = "s:6:\"-12.34\";"

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}


	x = "apple \"ok\" banana"
	expected = "s:17:\"apple \"ok\" banana\";"

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}
}



func TestSerializeArray(t *testing.T) {

	var serialized string
	var expected string

	var x [3]string



	x = [3]string{"a", "bb", "ccc"}
	expected = "a:3:{i:0;s:1:\"a\";i:1;s:2:\"bb\";i:2;s:3:\"ccc\";}"

	if serialized = Serialize(x); expected != serialized  {
		t.Errorf("Serialize() did not serialize [%#v] as expected [%#v]. Instead got [%#v].", x, expected, serialized)
	}
}
