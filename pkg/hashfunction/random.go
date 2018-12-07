package hashfunction

import (
	"math/rand"
)

type RandomizedHash struct {
	N int
	Seeds []int
}

func NewRandomizedHash(numHashes, arraySize, randSeed int) *RandomizedHash {
	if randSeed > 0 {
		rand.Seed(int64(randSeed))
	}
	var seeds []int
	for i := 0; i < numHashes; i++ {
		seeds = append(seeds, rand.Intn(arraySize))
	}
	return &RandomizedHash{
		N: arraySize,
		Seeds: seeds,
	}
}

func (r *RandomizedHash) GetHashList(x int) []int {
	var ret []int
	for i := 0; i < len(r.Seeds); i++ {
		rand.Seed(int64(r.Seeds[i] + x))
		ret = append(ret, rand.Intn(r.N))
	}
	return ret
}