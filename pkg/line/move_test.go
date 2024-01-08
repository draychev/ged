package line

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func TestMove4to0(t *testing.T) {
	from := 4
	dest := 0 // make the fourth line the first

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

	Move(from, from, dest, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `four
one
two
three
five
six
`
	if expected != actual {
		t.Errorf("Error testing Move(%d, %d, %d)\n>> expected: %s \n>> actually: %s", from, from, dest, expected, actual)
	}
}

func TestMove1to3(t *testing.T) {
	from := 1
	dest := 3 // make the first line the third; third becomes fourth

	cfg := &common.Config{
		FileName: "test.txt",
		Debug:    true,
	}

	// initialize cfg
	cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())
	CopyFile(cfg.FileName, cfg.TmpFileName, cfg) // copy the contents of the main file to the new temp file

	Move(from, from, dest, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `two
three
one
four
five
six
`

	if expected != actual {
		t.Errorf("Error testing Move(%d, %d, %d)\n>> expected: %s \n>> actually: %s", from, from, dest, expected, actual)
	}

}

func TestMove1comma2to3(t *testing.T) {
	from := 1
	to := 2
	dest := 3 // make the first line the third; third becomes fourth

	cfg := &common.Config{
		FileName: "test.txt",
		Debug:    true,
	}

	// initialize cfg
	cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())
	CopyFile(cfg.FileName, cfg.TmpFileName, cfg) // copy the contents of the main file to the new temp file

	Move(from, to, dest, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `three
one
two
four
five
six
`

	if expected != actual {
		t.Errorf("Error testing Move(%d, %d, %d)\n>> expected: %s \n>> actually: %s", from, to, dest, expected, actual)
	}

}
