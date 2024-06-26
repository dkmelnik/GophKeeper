package configs

type ClientConf struct {
	ADDR  string
	Token string
}

var Client = ClientConf{
	ADDR:  "",
	Token: "",
}
