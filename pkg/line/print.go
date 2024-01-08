package line

import (
	"bufio"
	"fmt"
	"os"

	"github.com/draychev/ged/pkg/common"
)

func PrintLines(fromLine, toLine int, withNumbers bool, cfg *common.Config) {
	tmpFile, err := os.Open(cfg.TmpFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Error opening input file (%s): %s\n", cfg.TmpFileName, err)
		return
	}
	defer tmpFile.Close()

	scanner := bufio.NewScanner(tmpFile)
	buff := makeBuff(fromLine, toLine)

	lineNum := 1
	for scanner.Scan() {
		if lineNum < fromLine {
			lineNum += 1
			continue
		}
		if lineNum > toLine {
			break
		}

		// TODO: move to a *buffer
		if withNumbers {
			fmt.Printf("%d%s%s\n", lineNum, buff, scanner.Text())
		} else {
			fmt.Printf("%s\n", scanner.Text())
		}
		lineNum += 1
	}

	if err := scanner.Err(); err != nil && cfg.Debug {
		log.Info().Msgf("%s: read: %s\n", cfg.FileName, err)
	}
}

func makeBuff(fromLine, toLine int) string {
	// TODO: memoize
	return "        "
}
