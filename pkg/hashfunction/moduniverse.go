package hashfunction

import (
	"math"
	"math/rand"
)

type ModUniverseHash struct {
	NumHashes int
	ArraySize int
	P int
	ARandom []int
	BRandom []int
}

func NewModUniverseHash(numHashes, universeSize, arraySize, randSeed int) *ModUniverseHash {

	if randSeed != 0 {
		rand.Seed(int64(randSeed))
	}
	var a, b []int
	for i := 0; i<numHashes; i++ {
		a = append(a, rand.Intn(arraySize) + 1)
		b = append(b, rand.Intn(arraySize))
	}

	return &ModUniverseHash{
		NumHashes: numHashes,
		ArraySize: arraySize,
		P: getNextPrime(universeSize),
		ARandom: a,
		BRandom: b,
	}

}

func (m *ModUniverseHash) GetHashList(x int) []int{
	var ret []int
	for i := 0; i<m.NumHashes; i++ {
		ret = append(ret, ((m.ARandom[i] * x + m.BRandom[i]) % m.P) % m.ArraySize)
	}
	return ret
}

func checkIfPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	sqrtN := math.Sqrt(float64(n))
	i := 5
	w := 2
	for {
		if float64(i) > sqrtN {
			break
		}
		if n % i == 0 {
			return false
		}
		i += w
		w = 6-w
	}
	return true
}

func getNextPrime(n int) int {
	if n % 2 == 0 {
		n++
	}
	for i := n; i<n+336; n+=2 {
		if checkIfPrime(i) {
			return i
		}
	}
	return -1
}