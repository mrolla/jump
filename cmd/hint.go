package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func hintCmd(args cli.Args, conf config.Config) error {
	term := termFromArgs(args, conf)

	entry, err := cdEntry(term, conf)
	if err == errNoEntries {
		return nil
	}
	if err != nil {
		return err
	}

	cli.Outf("%s\n", entry.Path)

	return nil
}

func init() {
	cli.RegisterCommand("hint", "Hints relevant paths for jumping.", hintCmd)
}
