//go:build !binary_log
// +build !binary_log

package zlog_diode_test

import (
	"fmt"
	"os"

	"github.com/sudoless/zerolog/pkg/zlog"
	zlogDiode "github.com/sudoless/zerolog/pkg/zlog-diode"
)

func ExampleNewWriter() {
	w := zlogDiode.NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zlog.New(w)
	log.Print("test")

	_ = w.Close()

	// Output: {"level":"debug","message":"test"}
}
