package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/uuid"

	"github.com/draychev/ged/pkg/args"
	"github.com/draychev/ged/pkg/common"
	"github.com/draychev/ged/pkg/inparser"
	"github.com/draychev/ged/pkg/line"
	"github.com/draychev/go-toolbox/pkg/logger"
)

var log = logger.NewPretty("main")

const (
	version                string = "0.1"
	defaultPrompt                 = ""
	columnsAfterLineNumber        = 8
)

var stop chan struct{}

var cfg = &common.Config{
	Prompt:   defaultPrompt,
	FileName: "",
	FileSize: -1,
	Debug:    false,
	CurLine:  1,
}

var exitSignals = []os.Signal{os.Interrupt, syscall.SIGTERM} // SIGTERM is POSIX specific

func getLine(line string) int {
	if line == "." {
		return cfg.CurLine
	}
	if line == "$" {
		// TODO
		return 9
	}
	lineNum, err := strconv.Atoi(line)
	if err != nil {
		// TODO
		return cfg.CurLine
	}
	return lineNum
}

func main() {
	stop = make(chan struct{})
	s := make(chan os.Signal, len(exitSignals))
	signal.Notify(s, exitSignals...)
	go func() {
		// Wait for a signal from the OS before dispatching
		// a stop signal to all other goroutines observing this channel.
		<-s
		close(stop)
		line.CleanUp(cfg)
		os.Exit(0)
	}()

	args.ParseArgs(cfg)

	var userInput string

	// TODO: move to loadFile
	if cfg.FileSize != -1 {
		fmt.Println(cfg.FileSize)
	}

	for {
		if cfg.TmpFileName == "" {
			cfg.TmpFileName = filepath.Join("/tmp", uuid.New().String())
			line.CopyFile(cfg.FileName, cfg.TmpFileName, cfg)
		}
		fmt.Print(cfg.Prompt)                    // show the configuration
		fmt.Scanln(&userInput)                   // read the user command
		userInput = strings.TrimSpace(userInput) // remove any trailing spaces
		cmd, err := inparser.ReadCommand(userInput)
		if err != nil {
			fmt.Println("?")
			continue
			// log.Error().Err(err).Msgf("Error parsing command: %s\n", err)
			// os.Exit(1)
		}

		switch cmd.Command {

		case "a":
			lines := line.UserWrites()
			line.AppendLine(cfg.CurLine, lines, cfg)
		case "c":
			panic("not implemnted")
		case "d":
			line.Delete(getLine(cmd.From), cfg)
		case "p":
			line.PrintLines(getLine(cmd.From), getLine(cmd.To), false, cfg)
			cfg.CurLine = getLine(cmd.To)
		case "n":
			line.PrintLines(getLine(cmd.From), getLine(cmd.To), true, cfg)
			cfg.CurLine = getLine(cmd.To)
		case "w":
			line.WriteFile(cfg)
		case "q":
			// TODO: print ? if there are unsaved changes as ed does
			//       on the second q - actually exit
			line.CleanUp(cfg)
			os.Exit(0)
		case "i":
			line.Insert(getLine(cmd.From), cfg)
		case "t":
			line.Transfer(getLine(cmd.From), getLine(cmd.To), getLine(cmd.Parameters[0]), cfg)
		case "m":
			line.Move(getLine(cmd.From), getLine(cmd.To), getLine(cmd.Parameters[0]), cfg)
		case "x":
			panic("not implemnted")
		case "=":
			line.ShowLineNumber(cfg)
		default:
			fmt.Println("?")
		}

		/*
			if userInput == "q" {
				cleanUp()
				os.Exit(0)
			} else if userInput == "p" {
				fromLine := 1
				toLine := 3
				printLines(fromLine, toLine, false)
			} else if userInput == "n" {
				fromLine := 1
				toLine := 3
				printLines(fromLine, toLine, true)
			} else if userInput == "a" {
				var lines []string
				for {
					fmt.Scanln(&userInput)
					if userInput == "." {
						break
					}
					lines = append(lines, userInput)
				}
				appendLine(cfg.CurLine, lines)
			} else if userInput == "w" {
				writeFile()
			} else {
				fmt.Println("?")
			}
		*/
	}
}
