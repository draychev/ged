package line

import (
	"github.com/draychev/ged/pkg/common"
)

func ShowLineNumber(cfg *common.Config) int {
	return cfg.CurLine
}
