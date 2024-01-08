package line

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/draychev/ged/pkg/common"
	"github.com/google/uuid"
)

func Move(from, to, dest int, cfg *common.Config) {
	if from == dest && to <= from { // nothing to do
		return
	}

	moveTempFileName := filepath.Join("/tmp", uuid.New().String())

	origFile, err := os.Open(cfg.TmpFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("[appendLine] Error opening input file %s", cfg.TmpFileName)
		}
		return
	}
	defer origFile.Close()

	outputFile, err := os.Create(moveTempFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error creating output file %s", moveTempFileName)
		}
		return
	}
	defer outputFile.Close()

	sourceLines, _ := GetLines(from, to, cfg) // TODO: handle error

	if dest == 0 { /* special case to insert the line at the very first position */
		for _, line := range sourceLines {
			if _, err := fmt.Fprintln(outputFile, line); err != nil {
				if cfg.Debug {
					log.Error().Err(err).Msgf("Error writing to output file %s", moveTempFileName)
				}
				return // TODO return error
			}
		}
	}

	lineNumber := 1

	scanner := bufio.NewScanner(origFile)
	for scanner.Scan() {
		line := scanner.Text()

		if lineNumber >= from && lineNumber <= to {
			lineNumber++
			continue
		}

		// keep the line in the file
		if _, err := fmt.Fprintln(outputFile, line); err != nil {
			if cfg.Debug {
				log.Error().Err(err).Msgf("Error writing to output file %s", moveTempFileName)
			}
			return // TODO return error
		}

		// append the moved line
		if lineNumber == dest {
			for _, line := range sourceLines {
				if _, err := fmt.Fprintln(outputFile, line); err != nil {
					log.Error().Err(err).Msgf("Error writing to output file %s", moveTempFileName)
					return // TODO return error
				}
				lineNumber += 1
			}
		}

		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error scanning input file %s", moveTempFileName)
		}
		return
	}

	origFile.Close()
	outputFile.Close()

	CopyFile(moveTempFileName, cfg.TmpFileName, cfg)

	if err := os.Remove(moveTempFileName); err != nil && cfg.Debug {
		log.Error().Err(err).Msgf("Error deleting file %s: %s\n", moveTempFileName, err)
	}
}
