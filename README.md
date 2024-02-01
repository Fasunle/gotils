# A simple utility package for golang

This package should be used only in development. If neccessary, the code can be used for production after neccessary security evaluations have been done.

## Installation

```bash
go get github.com/Fasunle/gotils
```

## Usage

```go

import (
    "fmt"
    "github.com/Fasunle/gotils"
)

func main(){

    var tools gotils.Tools

    random := tools.RandomString(12)
    fmt.Println(random)
}

```