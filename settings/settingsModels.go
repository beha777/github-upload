package settings


type Settings struct {
	AppParams 		Params				`json:"app"`
	PostgresParams	PostgresSettings	`json:"postgresParams"`
}

type Params struct {
	ServerName		string	`json:"serverName"`
	PortRun			string		`json:"portRun"`
}

type PostgresSettings struct {
	User		string 	`json:"user"`
	Password	string	`json:"password"`
	Server 		string	`json:"server"`
	Port		string	`json:"port"`
	DataBase 	string	`json:"database"`
}