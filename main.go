package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/imaimang/mgo/net"
	"github.com/imaimang/mgo/zip"
)

var config = new(Config)
var timerCheckUpdate *time.Timer

func main() {
	buffer, err := ioutil.ReadFile("./config.toml")
	if err == nil {
		_, err = toml.Decode(string(buffer), config)
		if err == nil {
			if len(config.WaitAddress) > 0 {
				net.WaitAddresses(func(address string, isSuccess bool) {
					log.Println("address available", isSuccess)
				}, config.WaitAddress...)
			}
			for _, server := range config.Servers {
				go run(server)
			}
			timerCheckUpdate = time.AfterFunc(1*time.Second, checkUpdate)

		}
	}
	if err != nil {
		log.Fatalln("daemon run error", err)
	}
	for {
		time.Sleep(24 * time.Hour)
	}
}

func checkUpdate() {
	for _, server := range config.Servers {
		if server.UpdateFileName != "" {
			ext := path.Ext(server.UpdateFileName)
			if server.Directory[len(server.Directory)-1] != '/' {
				server.Directory += "/"
			}
			filePath := server.Directory + server.UpdateFileName
			_, err := os.Stat(filePath)
			if err == nil {
				switch strings.ToLower(ext) {
				case ".zip":
					err := zip.Unzip(filePath, server.Directory)
					if err == nil {
						os.Remove(filePath)
					}
				case ".tar":
				}
			}
		}
	}
	timerCheckUpdate.Reset(1 * time.Second)
}

func run(server *Server) {
	log.Println("run server", server.Directory, server.Exe)
	for {
		if server.Directory != "" && server.Directory[len(server.Directory)-1] != '/' {
			server.Directory += "/"
		}
		cmd := exec.Command(server.Directory+server.Exe, server.Args...)
		cmd.Dir = server.Directory
		cmd.Env = server.Environment
		err := cmd.Run()
		if err != nil {
			log.Println(server.Directory, server.Exe, err)
		}
		if server.FaildSecond == 0 {
			server.FaildSecond = 5
		}
		time.Sleep(time.Duration(server.FaildSecond) * time.Second)
	}
}
