package world

import (
	"fmt"
	"testing"
)

func Test_coordinate_Relative(t *testing.T) {

	c, _ := NewCoordinate(401, 200, 200)

	south := c.Relative(0, 3)

	fmt.Println(south.String())

}
