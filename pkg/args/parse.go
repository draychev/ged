package args

import (
	"fmt"
	"os"
	"strings"

	"github.com/draychev/ged/pkg/common"
)

func ParseArgs(cfg *common.Config) {
	skipNextArg := false
	for idx, arg := range os.Args {
		if idx == 0 { // skip the first arument - it will be the binary itself
			continue
		}
		if skipNextArg {
			skipNextArg = false
			continue
		}
		if strings.HasPrefix(arg, "-p") {
			if len(arg) > 2 { // the prompt is immediately after -p
				cfg.Prompt = strings.TrimPrefix(arg, "-p")
			} else if len(os.Args) > idx+1 { // there is a space after -p
				cfg.Prompt = string(os.Args[idx+1])
				skipNextArg = true // Make sure we don't treat the prompt as a file
			}
		} else if strings.HasPrefix(arg, "--prompt") {
			cfg.Prompt = strings.TrimPrefix(arg, "--prompt=")
		} else if strings.HasPrefix(arg, "-d") {
			cfg.Debug = true
		} else if !strings.HasPrefix(arg, "-") { // must be a filename
			fileName := arg
			file, err := os.Stat(fileName)
			if err != nil {
				if cfg.Debug {
					fmt.Printf("%s: No such file or directory\n", fileName)
				}
				continue
			}
			cfg.FileName = arg
			cfg.FileSize = file.Size()
		}
	}
}
