package cli

import (
	"fudgemasterultra/go-authy/orm"

	"github.com/alexflint/go-arg"
)

func Entry() {
	arg.MustParse(&args)
	if args.CreateUser != nil {
		orm.CreateUser(args.CreateUser.Username, args.CreateUser.Password)
	}
	if args.SetDBEnvPath != nil {
		var db = args.SetDBEnvPath
		orm.IntialSetup(db.Host, db.User, db.Password, db.DBName, db.Port)
	}

}
