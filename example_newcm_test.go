package newcm_test

import (
	"fmt"
	"log"

	"github.com/rickykimani/go-newcm/newcm10regular"
	"golang.org/x/image/font/sfnt"
)

func Example() {
	otf, err := sfnt.Parse(newcm10regular.Font)
	if err != nil {
		log.Fatalf("could not parse New Computer Modern font: %+v", err)
	}
	fmt.Printf("num glyphs: %d\n", otf.NumGlyphs())

	// Output:
	// num glyphs: 3506
}
