package test

import (
	"testing"
	"basic_type"
)

func TestFunc(t testing.T) {
	basic_type.DoActionTime();
	t.Fail()
}

func TestBenchmark(b testing.B) {

}
