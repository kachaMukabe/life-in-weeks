package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/zserge/lorca"
	"gopkg.in/yaml.v2"
)

type Record struct {
	Name   string
	DOB    time.Time
	Colors string
	C2     string
	C3     string
	IsSet  bool
}

type Config struct {
	Record Record `yaml: "Settings"`
}

type Notes struct {
	Week int
	Note string
}

func getUserName() string {
	var envKey string
	if runtime.GOOS == "windows" {
		envKey = "USERNAME"
	} else {
		envKey = "USER"
	}
	return os.Getenv(envKey)
}

func isUserSet() bool {
	record := readSettings()
	isZero := record.DOB.IsZero()
	return !(isZero)

}

func createSettings(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr)
			}
			file.Close()
		}
	}
}

func writeSettings(config *Config) {
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("settings.yaml", data, 0)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func readSettings() Record {
	settings, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]Record)

	err2 := yaml.Unmarshal(settings, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
		fmt.Printf(v.Name)
	}
	return data["record"]

}

func updateDOB(name string, dob string, colors string) {
	pdob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		log.Fatal(err)
	}
	record := readSettings()
	record.Name = name
	record.DOB = pdob
	record.Colors = colors
	fmt.Printf("%s", record)
	config := Config{Record: record}
	writeSettings(&config)
}

func getWeeksDone() int {
	record := readSettings()
	t1 := record.DOB
	t2 := time.Now()
	days := t2.Sub(t1).Hours() / 24
	return int(days / 7)
}

func main() {
	ui, err := lorca.New("", "", 800, 600)
	ui.Bind("getUserName", getUserName)
	ui.Bind("getSettings", readSettings)
	ui.Bind("getWeeksDone", getWeeksDone)
	ui.Bind("isUserSet", isUserSet)
	ui.Bind("updateDOB", updateDOB)

	if err != nil {
		log.Fatal(err)
	}
	createSettings("settings.yaml")
	createSettings("notes.yaml")

	//config := Config{Record: Record{Name: "Kacha Mukabe", Colors: "#ee333"}}
	//writeSettings(&config)

	//readSettings()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.Dir("./www")))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	defer ui.Close()
	<-ui.Done()
}
