package Struct

type User struct {
	Id          int
	Username    string
	Password    string
	ProtectionQ string
	ProtectionA string
}

type Message struct {
	Id             int
	Tousername     string
	Fromusername   string
	Time           []uint8
	Messagecontent string
}
