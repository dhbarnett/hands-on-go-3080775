// challenges/packages/begin/proverbs.go
package main


// import the proverbs package
import (
  "github.com/jboursiquot/go-proverbs"
	"fmt"
)

// getRandomProverb returns a random proverb from the proverbs package

func getRandomProverb() string {
    return proverbs.Random().Saying
}

func main() {
	fmt.Println(getRandomProverb())
}
