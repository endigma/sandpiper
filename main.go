package main

// import "fmt"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/endigma/sandpiper/modes/minecraft"

	log "github.com/sirupsen/logrus"
)

var config Config
var basePath string
var docker bool

// Config contains the entire config file
type Config struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Port     string    `json:"port"`
	Interval string    `json:"interval"`
	Monitors []Monitor `json:"monitors"`
	Counter  map[string]bool
}

// Monitor contains the descriptors for a service
type Monitor struct {
	Name      string               `json:"name"`
	Address   string               `json:"address"`
	Mode      string               `json:"mode"`
	Port      int                  `json:"port"`
	Minecraft minecraft.ServerInfo `json:"minecraft"`
	Status    bool
}

func (mon *Monitor) check() {
	switch mon.Mode {
	case "http":
		resp, err := http.Get(mon.Address)
		checkErr(err)

		if resp.StatusCode == 200 {
			mon.Status = true
		}
	case "minecraft":
		conn, err := minecraft.EstablishConnection(mon.Address, mon.Port)
		if err != nil {
			break
		}

		result, err := minecraft.QueryServer(conn, mon.Address, uint16(mon.Port))
		checkErr(err)

		mon.Minecraft = *result
		mon.Status = true
	default:
		log.Warnf("Unsupported mode: %s", mon.Mode)
	}

	log.Infof("    %s: %s", mon.Name, mon.Status)
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}

func unpackConfig() Config {
	configFile, err := os.Open(os.Args[1])

	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	checkErr(err)

	var config Config

	err = json.Unmarshal(byteValue, &config)
	checkErr(err)

	log.Info("Successfully Unpacked Config")

	return config
}

func init() {
	if len(os.Args) == 1 {
		log.Fatal("Please provide a valid config file.")
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal("Please provide a valid config file.")
	}

	config = unpackConfig()

	config.Counter = make(map[string]bool)
	for i := 0; i < len(config.Monitors); i++ {
		config.Counter[config.Monitors[i].Mode] = true
	}

	updateMonitors()
}

func updateMonitors() {
	log.Info("Updating Monitors")
	for i := 0; i < len(config.Monitors); i++ {
		go config.Monitors[i].check()
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	go updateMonitors()
	tmpl := template.Must(template.ParseFiles("assets/statuspage.html"))
	tmpl.Execute(w, config)
}

func main() {
	fs := http.FileServer(http.Dir("assets/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/favicon.ico", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handle)

	log.Infof("Starting HTTP server on :%s!", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
