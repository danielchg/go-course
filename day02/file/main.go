package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if (len(os.Args)) < 2 {
		log.Fatal("Need to add a filene as argument to read")
	}

	var (
		w   io.Writer
		r   io.ReadCloser
		err error
		dir string
	)

	fn := os.Args[1]
	fnWrite := os.Args[2]
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal("cannot get work dir")
	}

	if !filepath.IsAbs(fn) {
		fn = filepath.Join(dir, fn)
	} else {
		fn = filepath.Clean((fn))
	}

	if !filepath.IsAbs(fnWrite) {
		fnWrite = filepath.Join(dir, fnWrite)
	} else {
		fnWrite = filepath.Clean((fnWrite))
	}

	if !filepath.HasPrefix(fn, dir) {
		log.Fatal("r path not allowed")
	}

	if !filepath.HasPrefix(fnWrite, dir) {
		log.Fatal("w path not allowed")
	}

	r, err = os.OpenFile(fn, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	defer r.Close()

	w, err = os.OpenFile(fnWrite, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalf("error opening file to write %s", fnWrite)
	}
	buffer := make([]byte, 64)

	for {
		var n, m int
		n, err = r.Read(buffer)
		if err != nil {
			break
		}
		m, err = w.Write(buffer[:n])
		if err != nil {
			break
		}
		if n != m {
			err = fmt.Errorf("Source value %d differ from destination content %d", n, w)
			break
		}
		fmt.Printf("%s\n", buffer[:n])
	}
	fmt.Println("")
	if err != nil && err != io.EOF {
		log.Fatalf("error reading file: %v", err)
	}
}
