package zlog_diode_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/sudoless/zerolog/pkg/zlog"
	zlog_diode "github.com/sudoless/zerolog/pkg/zlog-diode"
)

func Benchmark(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stderr)
	benchs := map[string]time.Duration{
		"Waiter": 0,
		"Pooler": 10 * time.Millisecond,
	}
	for name, interval := range benchs {
		b.Run(name, func(b *testing.B) {
			w := zlog_diode.NewWriter(ioutil.Discard, 100000, interval, nil)
			log := zlog.New(w)
			defer w.Close()

			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					log.Print("test")
				}
			})
		})
	}
}
