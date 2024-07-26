package reducer

import (
	"bytes"
	"strings"
	"testing"
)

func TestReduceMappedElements(t *testing.T) {

	sourceTxt := "abc\t1\ndef\t1\nghi\t1\nabc\t1"
	expectedElems := []string{"abc\t2", "def\t1", "ghi\t1"}

	source := strings.NewReader(sourceTxt)
	errors := bytes.NewBuffer(nil)
	destination := bytes.NewBuffer(nil)

	Reduce(source, errors, destination)

	result := destination.String()
	for _, elem := range expectedElems {
		if !strings.Contains(result, elem) {
			t.Errorf("Result doesn't contain '%v'", elem)
		}
	}
}

func TestReduceShouldIgnoreEmptyLines(t *testing.T) {

	sourceTxt := "  \t  "

	source := strings.NewReader(sourceTxt)
	errors := bytes.NewBuffer(nil)
	destination := bytes.NewBuffer(nil)

	Reduce(source, errors, destination)

	result := destination.String()
	errs := errors.String()

	if result != "" {
		t.Errorf("Result should contain only white characters (runes), but contains: '%v'", result)
	}

	if errs != "" {
		t.Errorf("No errors should be returned, but received: '%v'", errs)
	}
}

func TestReturnErrorForIncorrectInput(t *testing.T) {

	sourceTxt := "abc\t1\ndef\t_invalid_\nghi\t1\nabc\t1"

	source := strings.NewReader(sourceTxt)
	errors := bytes.NewBuffer(nil)
	destination := bytes.NewBuffer(nil)

	Reduce(source, errors, destination)

	errs := errors.String()

	if errs == "" {
		t.Errorf("Expected error but none was returned")
	}
}
