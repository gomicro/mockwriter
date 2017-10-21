# Pen Name
[![Build Status](https://travis-ci.org/gomicro/penname.svg)](https://travis-ci.org/gomicro/penname)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/penname)](https://goreportcard.com/report/github.com/gomicro/penname)
[![GoDoc](https://godoc.org/github.com/gomicro/penname?status.svg)](https://godoc.org/github.com/gomicro/penname)
[![License](https://img.shields.io/github/license/gomicro/penname.svg)](https://github.com/gomicro/penname/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/gomicro/penname.svg)](https://github.com/gomicro/penname/releases/latest)

A mock that implements the Closer & Writer interfaces for testing.

# Example

```go
import(
	"fmt"
	"io"
	"os"

	"github.com/gomicro/penname"
)

func main(){
	mockWrite := penname.New()
	mw := io.MultiWriter(os.Stdout, mockWrite)

	mw.Write("A random line to write")

	if strings.Contains( string(mockWrite.Written), "random" ){
		fmt.Println("Found a random")
	}
}
```
