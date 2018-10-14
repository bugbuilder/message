package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
)

// Ready represent the readiness from the host
type Ready struct {
	Database       string `json:"database"`
	ExternalAccess string `json:"external_access"`
}

var (
	c         Config
	readiness = false
)

func message(w http.ResponseWriter, r *http.Request, c *Config) {
	m := builder(c)
	w.Header().Set("Content-Type", "application/json")

	glog.V(3).Infoln("Encoding message to Json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if readiness {
		message(w, r, &c)
	} else {
		glog.V(2).Infof("Endpoint is waiting for Readiness")
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func live(w http.ResponseWriter, r *http.Request) {
	if readiness {
		glog.V(2).Infof("Liveness probe is alive")
		w.WriteHeader(http.StatusOK)
	} else {
		glog.V(2).Infof("Liveness probe is waiting for Readiness")
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func ready(w http.ResponseWriter, r *http.Request) {
	if readiness {
		r := Ready{
			"Ready",
			"Ready",
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(r); err != nil {
			panic(err)
		}
		glog.V(2).Infof("everything is going well")
	} else {
		glog.V(2).Infof("Readiness is working")
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\t%s -readinessStart 30 -config config.yml\n\n", os.Args[0])

		flag.PrintDefaults()
	}

	version := flag.Bool("version", false, "Print the version information")
	cFile := flag.String("config", "", "Configuration file")
	rStart := flag.Int("readinessStart", 10, "Readiness start seconds")
	lReload := flag.Int("livenessReload", 10, "Liveness reload seconds")
	lEnds := flag.Int("livenessEnds", 10, "Liveness ends seconds")

	flag.Parse()

	if *version {
		fmt.Println(NewInfo().Print())
		os.Exit(0)
	}

	p := os.Getenv("PORT")
	if p == "" {
		p = "9000"
	}

	e := os.Getenv("EXTERNAL")
	if *cFile == "" && e != "true" {
		flag.Usage()
		os.Exit(1)
	}

	glog.V(3).Infoln("Flags")
	glog.V(3).Infof("\tconfig : %v", *cFile)
	glog.V(3).Infof("\treadinessStart %v seconds", *rStart)
	glog.V(3).Infof("\tlivenessReload %v seconds", *lReload)
	glog.V(3).Infof("\tlivenessEnds %v seconds", *lEnds)

	_, err := c.Read(*cFile)
	if err != nil {
		glog.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/live", live)
	http.HandleFunc("/ready", ready)

	go func() {
		time.Sleep(time.Second * time.Duration(*rStart))
		readiness = true
	}()

	glog.V(1).Infof("Running on http://localhost:%v", p)
	glog.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", p), nil))
}
