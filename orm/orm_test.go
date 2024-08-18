package orm

import (
	"fmt"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	fmt.Println(hashPassword("test"))
}

func TestCreateConConfig(t *testing.T) {
	testCases := []ConnectionData{
		{
			Host:     "localhost",
			User:     "root",
			Password: "password",
			DBName:   "DB",
			Port:     "22",
		},
		/* Will need to add more*/
	}
	for _, v := range testCases {
		createConConfig(v)
	}
}
