package line

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func Delete(deleteThisLineNumber int, cfg *common.Config) {
	deleteTempFN := filepath.Join("/tmp", uuid.New().String())

	inputFile, err := os.Open(cfg.TmpFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("[appendLine] Error opening input file %s", cfg.TmpFileName)
		}
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(deleteTempFN)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error creating output file %s", deleteTempFN)
		}
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()

		// append the new lines
		if lineNumber == deleteThisLineNumber {
			lineNumber++
			continue
		}

		_, err := fmt.Fprintln(outputFile, line)
		if err != nil {
			if cfg.Debug {
				log.Error().Err(err).Msgf("Error writing to output file %s", deleteTempFN)
			}
			return
		}
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error scanning input file %s", deleteTempFN)
		}
		return
	}

	inputFile.Close()
	outputFile.Close()

	CopyFile(deleteTempFN, cfg.TmpFileName, cfg)

	if err := os.Remove(deleteTempFN); err != nil && cfg.Debug {
		log.Error().Err(err).Msgf("Error deleting file %s: %s\n", deleteTempFN, err)
	}
}
