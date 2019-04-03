package main

import (
	"os"

	"github.com/oberlademeister/indexmaker"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	err := indexmaker.Make(os.Args[1], indexmaker.WithLogger(l), indexmaker.OverwriteExisting(true))
	if err != nil {
		panic(err)
	}
}
