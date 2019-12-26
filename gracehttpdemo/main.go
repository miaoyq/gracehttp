package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/tabalt/gracehttp"
)

func main() {
	if len(os.Args) <=1 {
		log.Fatal("please set the commend: server or client !")
	}

	if os.Args[1] == "server" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("@Request user header@ user-def-key:", r.Header.Get("user-def-key"))
			fmt.Fprintf(
				w,
				"started at %s from pid %d. user-def-key: %s \n",
				time.Now(),
				os.Getpid(),
				r.Header.Get("user-def-key"),
			)
		})

		pid := os.Getpid()
		address := ":8080"

		log.Printf("process with pid %d serving %s.\n", pid, address)
		err := gracehttp.ListenAndServe(address, nil)
		log.Printf("process with pid %d stoped, error: %s.\n", pid, err)
	} else if os.Args[1] == "client" {
		var addr string
		var mdValues []string

		if len(os.Args) >= 2 {
			addr = os.Args[2]
		} else {
			log.Fatalf("Please setting the address of server")
		}

		l := 1
		if len(os.Args) > 3 {
			mdValues = os.Args[3:]
			l = len(mdValues)
		}

		client := &http.Client{}
		for {
			randData := rand.Int()
			index := randData % l

			req, err := http.NewRequest("GET", "http://" + addr, nil)
			if err != nil {
				log.Printf("Failed to create a client %v", err)
				return
			}
			req.Header.Add("user-def-key", mdValues[index])
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Request header:", req.Header)

			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			log.Println("Response Host:", resp.Request.Host)
			log.Println( "Response body:", string(body))

			time.Sleep(1 * time.Second)
		}
	}





}
