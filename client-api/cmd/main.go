package main

import (
	"clientapi/config"
	"clientapi/loader"
	"log"
	"os"
)

func importFile(fileName string) {
	if fileName == "" {
		log.Printf("[info] no file to load")
		return
	}

	log.Printf("[info] load file: %s", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("[error] file open %s: %v", fileName, err)
		return
	}
	defer file.Close()

	l, err := loader.New(file)
	if err != nil {
		log.Printf("[error] file read: %v", err)
	}

	count := 0
	for row, err := l.NextRow(); err != loader.ErrEOF; row, err = l.NextRow() {
		if err != nil {
			log.Printf("[error] file read: %v", err)
			return
		}

		count++
		log.Printf("+%v", row)
	}
	log.Printf("[info] file load complete, %d rows read and sent", count)
}

func main() {
	conf := config.Load()
	importFile(conf.FileName)
}
