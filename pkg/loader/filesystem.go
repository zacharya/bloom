package loader

import (	
	"os"

	"github.com/zacharya/bloom/pkg/bloomfilter"

	log "github.com/sirupsen/logrus"
)

type FileSystemLoader struct {
	Path string
	Data map[int]bool
}

func NewFileSystemLoader(path string) *FileSystemLoader {
	return &FileSystemLoader{
		Path: path,
		Data: make(map[int]bool),
	}
}

func (f *FileSystemLoader) Load(bf *bloomfilter.BloomFilter) error {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Errorf("Cannot open file %s: %v", file, err)
		return nil
	}
	if err := readData(file, f.Data, bf); err != nil {
		log.Errorf("Error reading data: %v", err)
		return err
	}
	return nil
}