package main

import (
	"strings"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"flag"
)

var (
	parallel  = flag.Int("parallel", 10, "parallel params")

)
func main() {
	
	flag.Parse()
	
	links := flag.Args()
	
	c := make(chan struct{}, *parallel)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(links))

	for _, link := range links {
		c <- struct{}{}
		go func(link string) {
			defer func() {
				waitGroup.Done()
				<-c
			}()
			checkLink(link)
		}(link)
	}
	waitGroup.Wait()
}

func checkLink(link string) {
	if strings.HasPrefix(link, "http://") == false{
		link = "http://"+link
	}
	resp, err := http.Get(link)
	if err != nil {
		log.Fatalln(err)
		
		return
	}
	defer resp.Body.Close()
	
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyString := string(bodyBytes)
	data := []byte(bodyString)
	fmt.Printf("%v %x \n",link, md5.Sum(data))
	
}