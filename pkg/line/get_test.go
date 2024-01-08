package line

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func TestGetLines(t *testing.T) {
	cfg := &common.Config{
		FileName: "test.txt",
		Debug:    true,
	}

	// initialize cfg
	cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())

	CopyFile(cfg.FileName, cfg.TmpFileName, cfg) // copy the contents of the main file to the new temp file

	// --- TEST

	actualLines, err := GetLines(3, 4, cfg)
	if err != nil {
		log.Error().Err(err).Msg("Error running GetLines(3,4)")
		t.Error("GetLines(3,4) returned an error")
	}

	expected := []string{"three", "four"}

	if !reflect.DeepEqual(expected, actualLines) {
		t.Errorf("Error testing GetLines(3,4)\n>> expected: %s \n>> actually: %s", expected, actualLines)
	}
}
