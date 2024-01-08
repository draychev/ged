package line

import (
	"testing"

	"github.com/draychev/ged/pkg/common"
)

func TestShowLineNumber(t *testing.T) {
	cfg := &common.Config{
		FileName: "test.txt",
		//	FileSize    int64
		//	TmpFileName string
		//	Debug       bool
		//	CurLine     int
	}
	actual := ShowLineNumber(cfg)
	expected := 0
	if expected != actual {
		t.Errorf("Error testing ShowLineNumber()\n>> expected: %d \n>> actually: %d", expected, actual)
	}

	cfg.CurLine = 99
	actual = ShowLineNumber(cfg)
	expected = 99
	if expected != actual {
		t.Errorf("Error testing ShowLineNumber()\n>> expected: %d \n>> actually: %d", expected, actual)
	}
}
