package line

import (
	"os"

	"github.com/draychev/ged/pkg/common"
)

func CleanUp(cfg *common.Config) {
	if cfg.Debug {
		log.Info().Msg("Cleaning up")
	}

	if err := os.Remove(cfg.TmpFileName); err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msgf("Error deleting file %s: %s\n", cfg.TmpFileName, err)
		}
		return
	}
}
