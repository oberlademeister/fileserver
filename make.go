package indexmaker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Make creates the index
func Make(root string, opts ...Option) error {
	config := DefaultConfig()

	// apply options
	for _, o := range opts {
		o(&config)
	}

	if err := doDir(root, config.recursionDepth, &config); err != nil {
		return errors.Wrap(err, "failed to do dir")
	}
	return nil
}

func doDir(dir string, recLevelsToGo int, config *Config) error {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read dir %s", dir))
	}
	indexPath := filepath.Join(dir, config.indexName)
	if _, errStat := os.Stat(indexPath); errStat != nil && os.IsNotExist(errStat) || config.overwriteExisting {
		// indexpath is either not existing or we shall overwrite
		w, err := os.Create(indexPath)
		if err != nil {
			if !config.skipOnError {
				return errors.Wrap(err, fmt.Sprintf("failed to create %s", indexPath))
			}
			goto NOCREATE
		}
		defer w.Close()
		// write body of index.html here
		config.indexRenderer(w, fis)
	}
NOCREATE:
	if recLevelsToGo == 0 {
		return nil
	}
	for _, fi := range fis {
		if fi.IsDir() {
			doDir(filepath.Join(dir, fi.Name()), recLevelsToGo - 1, config)
		}
	}

	return nil
}
