package line

import (
	"io"
	"os"

	"github.com/draychev/ged/pkg/common"
)

func WriteFile(cfg *common.Config) {
	fromFileName := cfg.TmpFileName
	toFileName := cfg.FileName
	CopyFile(fromFileName, toFileName, cfg)
}

func CopyFile(fromFileName, toFileName string, cfg *common.Config) {
	fromFile, err := os.Open(fromFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error opening original file (%s): %s\n", fromFileName, err)
		}
		return
	}
	defer fromFile.Close()

	toFile, err := os.Create(toFileName)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error creating destination file %s: %s\n", toFileName, err)
		}
		return
	}
	defer toFile.Close()

	if _, err = io.Copy(toFile, fromFile); err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error copying file from (%s) to (%s):%s\n", fromFileName, toFileName, err)
		}
		return
	}
	return
}
