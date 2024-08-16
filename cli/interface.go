package cli

import (
	"fmt"

	"github.com/alexflint/go-arg"
)


func Entry () {
	arg.MustParse(&args)
	if args.CreateUser != nil {
		fmt.Println("Create User")
	}
	if args.SetDBEnvPath != nil {
		fmt.Println("Set DB Env Path")
	}
	
}