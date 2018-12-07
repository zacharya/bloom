package cmd

import (
	"os"
	"fmt"

	"github.com/zacharya/bloom/pkg/hashfunction"
	"github.com/zacharya/bloom/pkg/bloomfilter"
	"github.com/zacharya/bloom/pkg/loader"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var (
	numHashesArg, arraySizeArg, universeSizeArg, randomSeedArg, hashFunctionArg, loaderArg, portArg int
	dataFileArg, checkFileArg string
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&universeSizeArg, "universe-size", "u", 2147483647, "size of universe")
	rootCmd.PersistentFlags().IntVarP(&arraySizeArg, "array-size", "a", 1917017, "size of bloom filter array")
	rootCmd.PersistentFlags().IntVarP(&numHashesArg, "hashes", "s", 13, "number of hash functions to use")
	rootCmd.PersistentFlags().IntVarP(&randomSeedArg, "random-seed", "r", 877623067, "random seed to use for hash functions")
	rootCmd.PersistentFlags().StringVarP(&dataFileArg, "data-file", "d", "testInput.dat", "datafile to load into bloom filter")
	rootCmd.PersistentFlags().StringVarP(&checkFileArg, "check-file", "c", "checker.dat", "datafile to check for correctness")
	rootCmd.PersistentFlags().IntVarP(&hashFunctionArg, "hash-func", "f", 2, "hash function to use.  currently either 1 - randomized or 2 - mod universe")
	rootCmd.PersistentFlags().IntVarP(&loaderArg, "loader", "l", 1, "loader to use.  currently either 1 - filesystem or 2 - http")
	rootCmd.PersistentFlags().IntVarP(&portArg, "port", "p", 8888, "port to use for http loader")
}

var rootCmd = &cobra.Command{
	Use:   "bloom",
	Short: "Bloom filter implementation",
	Long: `bloom implements Bloom filters with a choice of hash function`,
	Run: func(cmd *cobra.Command, args []string) {
		var hashFunc hashfunction.HashFunction
		switch hashFunctionArg {
		case 1:
			hashFunc = hashfunction.NewRandomizedHash(numHashesArg, arraySizeArg, randomSeedArg)
		default:
			hashFunc = hashfunction.NewModUniverseHash(numHashesArg, universeSizeArg, arraySizeArg, randomSeedArg)
		}

		var ldr loader.Loader
		switch loaderArg {
		case 2:
			ldr = loader.NewHTTPLoader(portArg)
		default:
			ldr = loader.NewFileSystemLoader(dataFileArg)
		}

		bloomFilter := &bloomfilter.BloomFilter{
			HashFunction: hashFunc,
			Array: make([]byte, arraySizeArg),
		}

		ldr.Load(bloomFilter)

		results, err := bloomFilter.Check(checkFileArg)
		if err != nil {
			log.Fatalf("Problem checking results: %v", err)
		}

		if len(results) == 0 {
			fmt.Println("BF working as expected!")
		} else {
			fmt.Println("Unexpected results:")
			for k, v := range results {
				fmt.Printf("%d: got %t, expected %t\n", k, v.Got, v.Expected)
			}
		}
	},
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}