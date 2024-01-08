package line

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func TestDelete4(t *testing.T) {
	from := 4

	cfg := &common.Config{
		FileName: "test.txt",
		//	FileSize    int64
		//	TmpFileName string
		//	Debug       bool
		//	CurLine     int
	}

	// initialize cfg
	cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())
	CopyFile(cfg.FileName, cfg.TmpFileName, cfg) // copy the contents of the main file to the new temp file

	Delete(from, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `one
two
three
five
six
`
	if expected != actual {
		t.Errorf("Error testing Delete(%d)\n>> expected: %s \n>> actually: %s", from, expected, actual)
	}
}

func TestDelete1(t *testing.T) {
	from := 1

	cfg := &common.Config{
		FileName: "test.txt",
		Debug:    true,
	}

	// initialize cfg
	cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())
	CopyFile(cfg.FileName, cfg.TmpFileName, cfg) // copy the contents of the main file to the new temp file

	Delete(from, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `two
three
four
five
six
`

	if expected != actual {
		t.Errorf("Error testing Delete(%d)\n>> expected: %s \n>> actually: %s", from, expected, actual)
	}
}
