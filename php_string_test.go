package php



import (
	"testing"
)



func TestStrposRune(t *testing.T) {

	var haystack string
	var needle rune
	var offset int

	var pos int
	var err error

	var expectedPos int



	haystack = ""
	needle = 'c'

	if pos, err = StrposRune(haystack, needle, offset); nil == err {
		t.Errorf("Excepted error when calling StrposRune() but got pos = [%v] for haystack = [%v] and needle = [%v].", pos, haystack, needle)
	}


	haystack = "bdfhjlnprtvxz"
	needle = 'c'

	if pos, err = StrposRune(haystack, needle, offset); nil == err {
		t.Errorf("Excepted error when calling StrposRune() but got pos = [%v] for haystack = [%v] and needle = [%v].", pos, haystack, needle)
	}


	haystack = "c"
	needle = 'c'
	expectedPos = 0

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}


	haystack = "ac"
	needle = 'c'
	expectedPos = 1

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}


	haystack = "aci"
	needle = 'c'
	expectedPos = 1

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}


	haystack = "cic"
	needle = 'c'
	expectedPos = 0

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}


	haystack = "acic"
	needle = 'c'
	expectedPos = 1

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}

	haystack = "acicw"
	needle = 'c'
	expectedPos = 1

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}


	haystack = "BaCbDcEdFeGfHgIhJiKjLkMlNmOnPoQpRqSrTsUtVuWvXwYxZyAz"
	needle = 'c'
	expectedPos = 5

	if pos, err = StrposRune(haystack, needle, offset); nil != err {
		t.Errorf("Excepted no error when calling StrposRune() but got err = [%v] with pos = [%v] for haystack = [%v] and needle = [%v].", err, pos, haystack, needle)
	} else if expectedPos != pos {
		t.Errorf("When calling StrposRune() expected pos = [%v] but instead of pos = [%v] for haystack = [%v] and needle = [%v].", expectedPos, pos, haystack, needle)
	}

}
