package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/EconomistDigitalSolutions/watchman/journal"
	_ "github.com/EconomistDigitalSolutions/watchman/meter"

	"github.com/joho/godotenv"
)

var (
	port        string
	buildstamp  string
	githash     string
	version     string
	serviceName string
)

func init() {
	flag.StringVar(&port, "port", ":9494", "port to listen on")
	flag.StringVar(&version, "version", "", "output build date and commit data")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = filepath.Base(os.Args[0])
	}
	journal.Service = serviceName
}

func main() {

	flag.Parse()

	if version != "" {
		journal.LogChannel("build", fmt.Sprintf("build date: %s commit: %s", buildstamp, githash))
	}

	journal.LogChannel("information", fmt.Sprintf("%s up on port %s", serviceName, port))
	log.Fatal(http.ListenAndServe(port, nil))
}
