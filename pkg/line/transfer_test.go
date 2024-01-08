package line

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func TestTransfer4to0(t *testing.T) {
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

	Transfer(from, from, dest, cfg)

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
four
five
six
`
	if expected != actual {
		t.Errorf("Error testing Transfer(%d, %d, %d)\n>> expected: %s \n>> actually: %s", from, from, dest, expected, actual)
	}
}

func TestTransfer1to3(t *testing.T) {
	from := 1
	dest := 3 // make the first line the third; third becomes fourth

	cfg := &common.Config{
		FileName: "test.txt",
		Debug:    true,
	}

	// initialize cfg
	cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())
	CopyFile(cfg.FileName, cfg.TmpFileName, cfg) // copy the contents of the main file to the new temp file

	Transfer(from, from, dest, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `one
two
three
one
four
five
six
`

	if expected != actual {
		t.Errorf("Error testing Transfer(%d, %d, %d)\n>> expected: %s \n>> actually: %s", from, from, dest, expected, actual)
	}

}

func TestTransfer4comma5to0(t *testing.T) {
	from := 4
	to := 5
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

	Transfer(from, to, dest, cfg)

	fileContent, err := os.ReadFile(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading %s", cfg.TmpFileName)
		t.Errorf("err")
	}

	actual := string(fileContent)
	expected := `four
five
one
two
three
four
five
six
`
	if expected != actual {
		t.Errorf("Error testing Transfer(%d, %d, %d)\n>> expected: %s \n>> actually: %s", from, to, dest, expected, actual)
	}
}
