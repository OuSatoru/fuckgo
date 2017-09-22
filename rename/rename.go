package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// sep := string(os.PathSeparator)
	var dir string
	flag.StringVar(&dir, "d", "s", "directory")
	flag.Parse()
	fmt.Println(dir)
	if dir != "" {
		// files, err := ioutil.ReadDir(dir)
		// if err != nil {
		// 	panic(err)
		// }
		// for _, file := range files {
		// 	fmt.Println(file.Name())
		// 	os.Rename(dir+sep+file.Name(), dir+sep+gbk.GBKtoUTF8(file.Name()))
		// }
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			println(path, f.Name())
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		panic("空")
	}
}
