package patterns_test

import(
	"patterns/singleton"
	"testing"
	"gotest"
)

func TestNew(t *testing.T){
	singleton := patterns.New()
	singleton1 := patterns.New()
}