package bloomfilter

import (
	"os"
	"strconv"
	"bufio"
	"strings"

	"github.com/zacharya/bloom/pkg/hashfunction"

	log "github.com/sirupsen/logrus"
)

type BloomFilter struct {
	HashFunction hashfunction.HashFunction
	Array []byte
}

type CheckerResult struct {
	Expected bool
	Got bool
}

type CheckerResults map[int]CheckerResult


func (b *BloomFilter) Add(x int) {
	hashes := b.HashFunction.GetHashList(x)
	for i:=0; i<len(hashes); i++ {
		b.Array[hashes[i]] = 1
	}
}

func (b *BloomFilter) Contains(x int) bool {
	hashes := b.HashFunction.GetHashList(x)
	if len(hashes) == 0 {
		return false
	}
	for i:=0; i<len(hashes); i++ {
		if b.Array[hashes[i]] == 0 {
			return false
		}
	}
	return true
}

func (b *BloomFilter) Check(path string) (CheckerResults, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Cannot open file %s: %v", file, err)
		return nil, err
	}
	results := make(CheckerResults)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		element, err := strconv.Atoi(line[0])
		if err != nil {
			log.Errorf("Error converting data to int: %v", err)
			return nil, err
		}
		expected, err := strconv.ParseBool(line[1])
		if err != nil {
			log.Errorf("Error converting element to int: %v", err)
			return nil, err
		}
		got := b.Contains(element)
		if got != expected {
			results[element] = CheckerResult{
				Expected: expected,
				Got: got,
			}
		}
	}
	return results, nil
}