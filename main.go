package main

import (
	"flag"
	"github.com/inhuman/bst-api/api"
	"github.com/inhuman/bst-api/bst"
	"github.com/inhuman/bst-api/log"
	"github.com/rs/zerolog"
	"io/ioutil"
	"os"
)

func main() {

	l := log.NewLogger()

	err := realMain(l)
	if err != nil {
		l.Err(err).Str("source", "main").Send()
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}

func realMain(l *zerolog.Logger) error {

	fileName := flag.String("file", "bst-data.json", "path to bst json data")
	flag.Parse()

	jsonFile, err := os.Open(*fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			l.Err(err).Str("source", "main").Send()
		}
	}()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	bstContainer, err := bst.NewBstContainer(bytes, l)
	if err != nil {
		return err
	}

	opts := api.Opts{
		BstContainer: bstContainer,
		Logger:       l,
	}

	if err := api.Run(opts); err != nil {
		return err
	}

	return nil
}
