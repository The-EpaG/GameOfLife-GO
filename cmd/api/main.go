package main

import (
	"flag"
	"log"

	"github.com/The-EpaG/GameOfLife-GO/cmd/initCommand"
	"github.com/The-EpaG/GameOfLife-GO/cmd/startCommand"
	"github.com/The-EpaG/GameOfLife-GO/internal/enum/mode"
	"github.com/The-EpaG/GameOfLife-GO/internal/errors"
)

func getMode() (mode.Mode, error) {
	var modeChosen mode.Mode

	initParam := flag.Bool("i", false, "init")
	startParam := flag.Bool("s", false, "start")

	initCommand.Flags()
	startCommand.Flags()

	flag.Parse()

	if (*initParam && *startParam) || (!*initParam && !*startParam) {
		return 0, &errors.ParamError{}
	}

	if *initParam {
		modeChosen = mode.Init
	} else if *startParam {
		modeChosen = mode.Start
	}

	return modeChosen, nil
}

func main() {
	modeFlag, err := getMode()
	if err != nil {
		log.Fatal(err)
	}

	switch modeFlag {
	case mode.Init:
		err := initCommand.InitCommand()
		if err != nil {
			log.Fatal(err)
		}
	case mode.Start:
		err := startCommand.StartCommand()
		if err != nil {
			log.Fatal(err)
		}
	}
}
