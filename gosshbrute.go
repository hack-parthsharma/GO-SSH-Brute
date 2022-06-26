package sshclient

//structure of GoSshBrute
type GoSshBrute struct {
	hostname string
	port int
	thread int
	username string
	password string
	passwordfile string
	passwords []byte
}