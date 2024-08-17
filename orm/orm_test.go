package orm

import (
	"fmt"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	fmt.Println(hashPassword("test"))
}
