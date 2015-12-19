package sink_test

import (
	"fmt"
	"os"

	"github.com/koron/go-sink"
)

func ExampleErrors() {
	errs := &sink.Errors{}

	// This is succeeded.
	f, err := os.Open("errors.go")
	errs.Put(err)
	defer f.Close()
	// Others will be all failed.
	_, err = os.Open("never.go")
	errs.Put(err)
	_, err = os.Open("never2.go")
	errs.Put(err)

	// Evaluate errors later.
	if errs.Has() {
		fmt.Println(errs.Error())
		// Output:
		// open never.go: The system cannot find the file specified.; open never2.go: The system cannot find the file specified.
	}
}
