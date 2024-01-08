package line

import (
	"bufio"
	"errors"
	"os"

	"github.com/draychev/ged/pkg/common"
)

// TODO: move these to types.go
var errReadingTempFile = errors.New("error reading temp file")

func GetLines(from, to int, cfg *common.Config) ([]string, error) {
	var lines []string
	origFile, err := os.Open(cfg.TmpFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("[appendLine] Error opening input file %s", cfg.TmpFileName)
		}
		return nil, errReadingTempFile
	}
	defer origFile.Close()

	scanner := bufio.NewScanner(origFile)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum >= from && lineNum <= to {
			lines = append(lines, line)
		}
		lineNum++
	}
	return lines, nil
}
