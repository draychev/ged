package line

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func AppendLine(cureLine int, lines []string, cfg *common.Config) {
	appendTempFN := filepath.Join("/tmp", uuid.New().String())

	inputFile, err := os.Open(cfg.TmpFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("[appendLine] Error opening input file %s", cfg.TmpFileName)
		}
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(appendTempFN)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error creating output file %s", appendTempFN)
		}
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()

		// append the new lines
		if lineNumber == cfg.CurLine {
			for _, newLine := range lines {
				if _, err := fmt.Fprintln(outputFile, newLine); err != nil && cfg.Debug {
					log.Error().Err(err).Msgf("Error writing to %s: %s", appendTempFN, err)
				}
			}
		}

		_, err := fmt.Fprintln(outputFile, line)
		if err != nil {
			if cfg.Debug {
				log.Error().Err(err).Msgf("Error writing to output file %s", appendTempFN)
			}
			return
		}
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error scanning input file %s", appendTempFN)
		}
		return
	}

	inputFile.Close()
	outputFile.Close()

	CopyFile(appendTempFN, cfg.TmpFileName, cfg)

	if err := os.Remove(appendTempFN); err != nil && cfg.Debug {
		log.Error().Err(err).Msgf("Error deleting file %s: %s\n", appendTempFN, err)
	}

	if cfg.Debug {
		log.Info().Msgf("File copied line by line from %s to %s\n", cfg.TmpFileName, appendTempFN)
	}
}
