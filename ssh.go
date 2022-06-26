package sshclient

import (
	"fmt"
	"log"
	"golang.org/x/crypto/ssh"
	"strings"
	"os"
)

//connect to ssh server
func sshConnect(port int, hostname,username,password string)  {
	//ssh config
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
	socket := fmt.Sprintf("%s:%d",hostname,port)

    _,  err:= ssh.Dial("tcp", socket, config) //connects
    if err != nil {
        if strings.Contains(err.Error(),"unable to authenticate") == true{ //check invalid creds
			fmt.Println("[+] Failed -->", username,":",password)
			return
		}
		log.Fatal(err) //other errors
    }
	fmt.Println("[+] Success -->",username,":",password) //success
	os.Exit(1)
}

