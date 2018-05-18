package main

import (
	"github.com/bborbe/version"
	"github.com/bborbe/version/compare"
)

func main() {
	compare.Main(version.GreaterThan)
}
