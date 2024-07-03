package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOne(t *testing.T) {
	//var (
	//	input  = 1
	//	output = 3
	//)
	//
	//actual := AddOne(1)
	//if actual != output {
	//	t.Errorf("AddOne %d,  input %d, actual %d", input, output, actual)
	//}

	assert.Equal(t, AddOne(2), 4, "AddOne (2) should be equal 3")

}
