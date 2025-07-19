# go-newcm

[![GitHub release](https://img.shields.io/github/release/rickykimani/go-newcm.svg)](https://github.com/rickykimani/go-newcm/releases)
[![GoDoc](https://godoc.org/github.com/rickykimani/go-newcm?status.svg)](https://godoc.org/github.com/rickykimani/go-newcm)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/rickykimani/go-newcm/raw/main/LICENSE)

`go-newcm` provides the [New Computer Modern](https://ctan.org/pkg/newcomputermodern) fonts as importable Go packages.

The fonts are released under the [GUST font](https://github.com/rickykimani/go-newcm/raw/main/GUST-FONT-LICENSE.txt) license.
The Go packages under the [MIT](https://github.com/rickykimani/go-newcm/raw/main/LICENSE) license.

## Example

```go
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
```

## Available Fonts

The package includes all variants of New Computer Modern fonts:

- Regular, Bold, Book, Italic styles
- Sans Serif variants
- Monospace variants  
- Mathematical symbols
- Uncial variants
- Devanagari script support
