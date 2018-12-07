package loader

import (
	"io"
	"bufio"
	"strconv"

	"github.com/zacharya/bloom/pkg/bloomfilter"

	log "github.com/sirupsen/logrus"
)

type Loader interface {
	Load(bf *bloomfilter.BloomFilter) error
}

type Data map[int]bool

func readData(body io.ReadCloser, data Data, bf *bloomfilter.BloomFilter) error {
	defer body.Close()
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		element, err := strconv.Atoi(line)
		if err != nil {
			log.Errorf("Error converting data to int: %v", err)
			return err
		}
		bf.Add(element)
	}
	return nil
}