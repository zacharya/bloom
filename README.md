## Bloom

```
bloom implements a Bloom filter with your choice of hash function

Usage:
  bloom [flags]

Flags:
  -a, --array-size int      size of bloom filter array (default 1917017)
  -c, --check-file string   datafile to check for correctness (default "checker.dat")
  -d, --data-file string    datafile to load into bloom filter (default "testInput.dat")
  -f, --hash-func int       hash function to use.  currently either 1 - randomized or 2 - mod universe (default 2)
  -s, --hashes int          number of hash functions to use (default 13)
  -h, --help                help for bloom
  -l, --loader int          loader to use.  currently either 1 - filesystem or 2 - http (default 1)
  -p, --port int            port to use for http loader (default 8888)
  -r, --random-seed int     random seed to use for hash functions (default 877623067)
  -u, --universe-size int   size of universe (default 2147483647)
  ```

The two loaders currently implemented are `FileSystemLoader` and `HTTPLoader` which read from from files you specify and from HTTP POSTs to a temporary HTTP server that spins up, respectively. You can also implement your own using the `Loader` interface.

The two hash functions currently implemented are a pseudorandom hash function and a function based on modular arithmetic.  You can also implement your own using the `HashFunction` interface.
