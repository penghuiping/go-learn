package times

import (
	"testing"
	"time"
	"fmt"
)

func TestTicker(in *testing.T)  {
	c := time.Tick(5 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, "s")
	}
}
