package mapper

import (
	"bytes"
	"strings"
	"testing"
)

func TestMappingSimpleText(t *testing.T) {

	sourceTxt := " abc def     ghi  ddd  \n jjj  lll  "
	expectedElems := []string{"abc\t1", "def\t1", "ghi\t1", "ddd\t1", "jjj\t1", "lll\t1"}

	source := strings.NewReader(sourceTxt)
	destination := bytes.NewBuffer(nil)

	Map(source, destination)

	result := destination.String()
	for _, elem := range expectedElems {
		if !strings.Contains(result, elem) {
			t.Errorf("Result doesn't contain '%v'", elem)
		}
	}
}

func TestMappingTextWithElementsToIgnore(t *testing.T) {

	sourceTxt := " .abc.. ,ghi,, :ddd:: \n ;jjj;;  \\lll\\\\  'klk''   ?gdg??  !opo!! (kqw(( )lkj)) \n  -bvb--  \tkkk\t _ldl__  \"bgb\"\""
	expectedElems := []string{"abc\t1", "ghi\t1", "ddd\t1", "jjj\t1", "lll\t1", "klk\t1", "gdg\t1",
		"opo\t1", "kqw\t1", "lkj\t1", "bvb\t1", "kkk\t1", "ldl\t1", "bgb\t1"}

	source := strings.NewReader(sourceTxt)
	destination := bytes.NewBuffer(nil)

	Map(source, destination)

	result := destination.String()
	for _, elem := range expectedElems {
		if !strings.Contains(result, elem) {
			t.Errorf("Result doesn't contain '%v'. Output: %v", elem, result)
		}
	}
}
