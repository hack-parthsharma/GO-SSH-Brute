package main

import (
	"ssh/gosshbrute"
	"sync"
	"flag"
)
func main()  {
	//cmd flags
	hostname := flag.String("host","","Hostname or ip of the target.")
	port := flag.Int("port", 22, "Port number of target.")
	user := flag.String("user","","username of target.")
	//password := flag.String("p","","password of the target")
	thread := flag.Int("t",5,"number of thread to use")
	passwordfile := flag.String("P","","password file")
	flag.Parse()

	var goSshBrute sshclient.GoSshBrute 
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	
	goSshBrute.Init(*hostname,*port,*user,"",*passwordfile, *thread)
	goSshBrute.Bruteforce(&wg,cond)
	wg.Wait()
}