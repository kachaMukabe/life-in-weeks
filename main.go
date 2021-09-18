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
	Name string
	DOB  time.Time
	C1   string
	C2   string
	C3   string
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
	fmt.Printf("%s", data["record"])
	return data["record"]

}

func main() {
	ui, err := lorca.New("", "", 800, 600)
	ui.Bind("getUserName", getUserName)
	ui.Bind("getSettings", readSettings)

	if err != nil {
		log.Fatal(err)
	}
	createSettings("settings.yaml")
	createSettings("notes.yaml")

	config := Config{Record: Record{Name: "Kacha Mukabe", C1: "#ee333"}}
	writeSettings(&config)

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
