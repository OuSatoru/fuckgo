package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/OuSatoru/fuckgo/common"
)

var (
	users = []string{"687", "1298", "1299"}
)

type jsonTime time.Time

func (jt *jsonTime) UnmarshalJSON(data []byte) error {
	stamp, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*jt = jsonTime(time.Unix(stamp/1000, 0))
	return nil
}

func (jt jsonTime) String() string {
	return time.Time(jt).Format("2006-01-02 15:04:05")
}

type infoJSON struct {
	Name       string   `json:"name"`
	ExamModel  int      `json:"examModel"`
	CreateTime jsonTime `json:"createTime"`
	BeginDate  jsonTime `json:"beginDate"`
	EndDate    jsonTime `json:"endDate"`
}

type one struct {
	UserID     int   `json:"userId"`
	TopicIndex int   `json:"topicIndex"`
	Topic      topic `json:"topic"`
}

type topic struct {
	Content     string `json:"content"`
	TopicOption string `json:"topicOption,omitempty"`
	Answer      string `json:"answer"`
}

func main() {
	var bg bool
	flag.BoolVar(&bg, "bg", true, "background, -bg=false/true")
	flag.Parse()
	if bg {
		args := []string{}
		for _, arg := range os.Args[1:] {
			if arg != "-bg=false" {
				args = append(args, arg)
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		log.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	log.Println("modified file:", event.Name)
				// }
				if isDir(event.Name) && event.Op == fsnotify.Create {
					log.Println("create directory:", event.Name)
					go watchDir(event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add("F:\\Work\\afaafeaha")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func isDir(p string) bool {
	finfo, err := os.Stat(p)
	if err != nil {
		return false
	}
	return finfo.IsDir()
}

func watchDir(dir string) {
	timer := time.NewTicker(2 * time.Second)
	for {
		<-timer.C

		if common.Exists(path.Join(dir, "info.json")) {
			info, err := ioutil.ReadFile(path.Join(dir, "info.json"))
			if err != nil {
				log.Print(err)
				return
			}
			var infoj infoJSON
			err = json.Unmarshal(info, &infoj)
			if err != nil {
				log.Print(err)
				return
			}
			if infoj.ExamModel != 1 {
				log.Println(infoj.Name, "非选择题")
				return
			}
			dinfo, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Print(err)
				return
			}
			if len(dinfo) != 2 {
				continue
			}
			var topicFile string
			for _, d := range dinfo {
				if d.IsDir() {
					topicFile = path.Join(dir, d.Name(), "paper", "u687", "i1.json")
				}
			}
			if !common.Exists(topicFile) {
				continue
			}
			t, err := ioutil.ReadFile(topicFile)
			if err != nil {
				log.Print(err)
				return
			}
			var topics []one
			err = json.Unmarshal(t, &topics)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Println(topics)
		}
	}
}
