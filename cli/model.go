package cli

type CreateUser struct {
	Username string `arg:"required, -u"`
	Password string `arg:"required, -p"`
	Email string `arg:"required, -e"`
}

type SetDBEnvPath struct {
	Host string `arg:"required, -h"`
	User string `arg:"required, -u"`
	Password string `arg:"required, -p"`
	DBName string `arg:"required, -d"`
	Port string `arg:"required, --port"`
}


var args struct {
	CreateUser *CreateUser `arg:"subcommand:create-user"`
	SetDBEnvPath *SetDBEnvPath `arg:"subcommand:db-setup"`
}
