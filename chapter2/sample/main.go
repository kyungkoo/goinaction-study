package main

import (
	"log"
	"os"

	_ "github.com/kyungkoo/goinaction-study/chapter2/sample/matchers"
	"github.com/kyungkoo/goinaction-study/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Homes")
}
