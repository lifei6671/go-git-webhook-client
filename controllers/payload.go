package controllers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/go-ini/ini"
	"log"
	"io/ioutil"
	"github.com/lifei6671/go-git-webhook-client/models"
	"os"
	"github.com/lifei6671/go-git-webhook-client/commands"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)

func Payload(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	var err error
	defer func() {

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w,"failure")
		}else{
			fmt.Fprint(w,"success")
		}
	}()

	if key,ok := params["key"];ok {
		f,err := ini.LooseLoad("conf/app.conf");
		if err != nil {
			log.Printf("Configuration file error:%s",err)
			fmt.Fprint(w,"error")
			return
		}

		if section,err := f.GetSection(key);err == nil {
			repo_name,err := section.GetKey("repo_name")
			if err != nil {
				log.Printf("Error:%s",err)
				return
			}
			branch_name,err := section.GetKey("branch_name")
			if err != nil {
				log.Printf("Error:%s",err)
				return
			}
			var command string

			command_key,err := section.GetKey("command")

			if err != nil {
				log.Printf("Error:%s",err)
				return
			}
			command = command_key.String()

			if command == "" {
				log.Println("Command is null.")
				return
			}
			var log_path string
			log_path_key,err := section.GetKey("log_path")
			if err != nil {
				log_path = "logs/log.log"
			}else{
				log_path = log_path_key.String()
			}

			var logger *log.Logger
			if _, err := os.Stat(log_path); os.IsNotExist(err) {
				log_file,err := os.Create(log_path)
				if err != nil {
					log.Printf("log file create failure:%s",err.Error())
				}else{
					logger = log.New(log_file,"",log.LstdFlags|log.Llongfile);
				}

			}else{
				log_file,err := os.OpenFile(log_path,os.O_APPEND, 0666)
				if err != nil {
					log.Printf("log file open failure:%s",err.Error())
				}else{
					logger = log.New(log_file,"",log.LstdFlags|log.Llongfile);
				}
			}

			b,err := ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Printf("Read body error:%s",err.Error())
				return
			}
			defer r.Body.Close()
			body := string(b)

			hook,err := models.ResolveHookRequest(body)
			if err != nil {
				logger.Printf("Resolve webhook data error:%s",err.Error())
				return
			}
			if repo,err := hook.RepositoryName();err != nil || repo != repo_name.String() {
				if err != nil {
					logger.Printf("Repository name error:%s",err.Error())
					return
				}
				logger.Printf("Repository name do not match:%s",repo)
				return
			}
			if bran,err := hook.BranchName();err != nil || (bran != branch_name.String() && bran != "heads/"+ branch_name.String()){
				if err != nil {
					logger.Printf("Branch name error:%s",err.Error())
					return
				}
				logger.Printf("Branch name do not match:%s",bran)
				return
			}

			command = commands.ResolveShellFilePath(command)

			channel := make(chan []byte, 10)
			go commands.Command(command,channel)
			isChannelClosed := false

			for {
				if isChannelClosed {
					logger.Println("chan closed.");
					break
				}
				select {
				case out, ok := <-channel:
					{
						if !ok {
							isChannelClosed = true
							break
						}
						if len(out) > 0 {
							logger.Println(string(out))
						}
					}
				}
			}

		}
		if err != nil {
			log.Printf("Error:%s",err)
			return
		}
	}

}
