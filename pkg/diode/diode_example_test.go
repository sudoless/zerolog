//go:build !binary_log
// +build !binary_log

package diode_test

import (
	"fmt"
	"os"

	"github.com/sudoless/zerolog/pkg/diode"
	"github.com/sudoless/zerolog/pkg/zerolog"
)

func ExampleNewWriter() {
	w := diode.NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zerolog.New(w)
	log.Print("test")

	_ = w.Close()

	// Output: {"level":"debug","message":"test"}
}
