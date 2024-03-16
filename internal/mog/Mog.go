package mog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

/*
fmt.Fprintf(mog.MW, "\n %s something %d, once = %t", me.Id, 2, true)

fmt.Fprintln(mog.MW, "\n hello")
*/
var MW io.Writer = nil

func Init() error {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	dir := fmt.Sprintf("%s%s%s%s", exPath, string(os.PathSeparator), "mogs/", string(os.PathSeparator))
	filePath := dir + time.Now().Format("2006.01.02.15.04.05") + ".mog.txt"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	MW = io.MultiWriter(os.Stdout, file)
	return nil
}
