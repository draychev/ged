package line

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func Transfer(from, to, dest int, cfg *common.Config) {
	if from == dest && to <= from { // nothing to do
		return
	}

	transferTempFileName := filepath.Join("/tmp", uuid.New().String())

	origFile, err := os.Open(cfg.TmpFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("[appendLine] Error opening input file %s", cfg.TmpFileName)
		}
		return
	}
	defer origFile.Close()

	outputFile, err := os.Create(transferTempFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error creating output file %s", transferTempFileName)
		}
		return
	}
	defer outputFile.Close()

	sourceLines, _ := GetLines(from, to, cfg) // TODO: handle error

	if dest == 0 { /* special case to insert the line at the very first position */
		for _, line := range sourceLines {
			if _, err := fmt.Fprintln(outputFile, line); err != nil {
				if cfg.Debug {
					log.Error().Err(err).Msgf("Error writing to output file %s", transferTempFileName)
				}
				return // TODO return error
			}
		}
	}

	lineNumber := 1

	scanner := bufio.NewScanner(origFile)
	for scanner.Scan() {
		line := scanner.Text()

		// keep the line in the file
		if _, err := fmt.Fprintln(outputFile, line); err != nil {
			if cfg.Debug {
				log.Error().Err(err).Msgf("Error writing to output file %s", transferTempFileName)
			}
			return // TODO return error
		}

		// append the transferd line
		if lineNumber == dest {
			for _, line := range sourceLines {
				if _, err := fmt.Fprintln(outputFile, line); err != nil {
					if cfg.Debug {
						log.Error().Err(err).Msgf("Error writing to output file %s", transferTempFileName)
					}
					return // TODO return error
				}
				lineNumber += 1
			}
		}

		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error scanning input file %s", transferTempFileName)
		}
		return
	}

	origFile.Close()
	outputFile.Close()

	CopyFile(transferTempFileName, cfg.TmpFileName, cfg)

	if err := os.Remove(transferTempFileName); err != nil && cfg.Debug {
		log.Error().Err(err).Msgf("Error deleting file %s: %s\n", transferTempFileName, err)
	}
}
