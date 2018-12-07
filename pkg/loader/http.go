package loader

import (
	"net/http"
	"fmt"
	"context"

	"github.com/zacharya/bloom/pkg/bloomfilter"
	
	log "github.com/sirupsen/logrus"
)

type HTTPLoader struct {
	Server *http.Server
	Port int
	Data map[int]bool
}

func NewHTTPLoader(port int) *HTTPLoader {
	return &HTTPLoader{
		Server: &http.Server{Addr: fmt.Sprintf(":%d", port)},
		Port: port,
		Data: make(map[int]bool),
	}
}

func (h *HTTPLoader) Load(bf *bloomfilter.BloomFilter) error {
	ctx, cancel := context.WithCancel(context.Background())

	http.HandleFunc("/done", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bye")
		cancel()
	})

	http.HandleFunc("/load", func (w http.ResponseWriter, r *http.Request) {
		if err := readData(r.Body, h.Data, bf); err != nil {
			log.Errorf("Error converting data to int: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
	})
	go func(){
		if err := h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	
	<-ctx.Done()

	if err := h.Server.Shutdown(ctx); err != nil && err != context.Canceled {
        log.Println(err)
    }
    log.Println("Done loading from http")
	
	return nil
}
