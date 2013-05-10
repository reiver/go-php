package php



import (
	"testing"
)



func TestSerializeString(t *testing.T) {

        var s string
	var expected string



	s = "apple"
	expected = "s:5:\"apple\";"

	if expected != Serialize(s)  {
	        t.Errorf("Serialize() did not serialize [%v] as expected: [%v]", s, expected)
	}


	s = "123"
	expected = "s:3:\"123\";"

	if expected != Serialize(s)  {
	        t.Errorf("Serialize() did not serialize [%v] as expected: [%v]", s, expected)
	}


	s = "-3"
	expected = "s:2:\"-3\";"

	if expected != Serialize(s)  {
	        t.Errorf("Serialize() did not serialize [%v] as expected: [%v]", s, expected)
	}


	s = "12.34"
	expected = "s:5:\"12.34\";"

	if expected != Serialize(s)  {
	        t.Errorf("Serialize() did not serialize [%v] as expected: [%v]", s, expected)
	}


	s = "-12.34"
	expected = "s:6:\"-12.34\";"

	if expected != Serialize(s)  {
	        t.Errorf("Serialize() did not serialize [%v] as expected: [%v]", s, expected)
	}


	s = "apple \"ok\" banana"
	expected = "s:17:\"apple \"ok\" banana\";"

	if expected != Serialize(s)  {
	        t.Errorf("Serialize() did not serialize [%v] as expected: [%v]", s, expected)
	}
}

