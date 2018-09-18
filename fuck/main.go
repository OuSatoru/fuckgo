package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/OuSatoru/fuckgo/common"
)

var (
	wg sync.WaitGroup
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
	// Content     string `json:"content"`
	// TopicOption string `json:"topicOption,omitempty"`
	Answer string `json:"answer"`
}

func main() {
	al, err := ioutil.ReadFile("C:\\Users\\w\\Desktop\\help\\全市人口信息.txt")
	if err != nil {
		log.Fatal(err)
	}
	spaces := regexp.MustCompile(`\x20+`)

	alr := spaces.ReplaceAll(al, []byte(","))
	err = ioutil.WriteFile("C:\\Users\\w\\Desktop\\help\\全市人口信息2.txt", alr, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func task() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		dir := "F:\\Work\\Go\\src"
		if common.Exists(dir+"/ok.end") || common.Exists(dir+"/ok.lock") {
			continue
		}
		lock, _ := common.CreateFile(dir + "/ok.lock")
		lock.Close()
		filepath.Walk(dir, func(p string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			if path.Ext(p) == ".go" {
				wg.Add(1)
				go func() {
					fmt.Println(p)
					wg.Done()
				}()

			}

			return nil
		})

		wg.Wait()
		os.Rename(dir+"/ok.lock", dir+"/ok.end")
	}
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
			break
		}
	}
}
