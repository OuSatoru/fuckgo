package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/OuSatoru/fuckgo/common"
)

var wg sync.WaitGroup

func main() {
	// wg.Add(1)
	dt := "20180530"
	task(dt)
	// fmt.Println(table(`/root/data/20180530/CMIS/P_094_BUSINESS_CONTRACT_20180530.del`))
}

func task(dt string) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		dir := fmt.Sprintf("/root/data/%s", dt)
		if common.Exists(dir+"/ok.end") || common.Exists(dir+"/ok.lock") {
			continue
		}
		lock, _ := common.CreateFile(dir + "/ok.lock")
		lock.Close()
		var way string
		filepath.Walk(dir, func(p string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() || f.Name() == "ok.OK" {
				return nil
			}
			if path.Ext(p) == ".OK" {
				r, _ := ioutil.ReadFile(p)
				way = strings.TrimSuffix(string(r), "\n")
				fmt.Println(p, way)
			}
			if path.Ext(p) == ".del" {
				switch way {
				case "0":
					wg.Add(1)
					go func() {
						tab := table(p)
						log.Println(p, " -> ", tab)
						out, err := exec.Command("./replace.sh", p, tab).Output()
						if err != nil {
							fmt.Printf("exec error:%v", err)
						}
						log.Printf("%s\n", out)
						wg.Done()
					}()

				case "1":
					wg.Add(1)
					go func() {
						tab := table(p)
						log.Println(p, " -> ", tab)
						out, err := exec.Command("./insert.sh", p, tab).Output()
						if err != nil {
							fmt.Printf("exec error:%v", err)
						}
						log.Printf("%s\n", out)
						wg.Done()
					}()
				case "2":
					wg.Add(1)
					go func() {
						tab := table(p)
						log.Println(p, " -> ", tab)
						out, err := exec.Command("./insert.sh", p, tab).Output()
						if err != nil {
							fmt.Printf("exec error:%v", err)
						}
						log.Printf("%s\n", out)
						wg.Done()
					}()
				}
			}

			return nil
		})

		wg.Wait()
		os.Rename(fmt.Sprintf("/root/data/%s/ok.lock", dt), fmt.Sprintf("/root/data/%s/ok.end", dt))
	}
}

func table(p string) string {
	reg := regexp.MustCompile(`/root/data/\d{8}/(.+)/P_094_(.+)_\d{8}.del`)
	match := reg.FindAllStringSubmatch(p, -1)
	return match[0][1] + "_" + match[0][2]
}
