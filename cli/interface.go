package cli

import (
	"fmt"
	"fudgemasterultra/go-authy/orm"

	"github.com/alexflint/go-arg"
)


func Entry () {
	arg.MustParse(&args)
	if args.CreateUser != nil {
		fmt.Println("Create User")
	}
	if args.SetDBEnvPath != nil {
		var db = args.SetDBEnvPath
		orm.IntialSetup(db.Host, db.User, db.Password, db.DBName, db.Port)
	}
	
}