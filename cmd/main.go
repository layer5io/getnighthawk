package main

import (
	"fmt"
	"os"

	"github.com/layer5io/nighthawk-go/api"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Output to only for logs above warn level
	log.SetLevel(log.WarnLevel)

}

func main() {

	//Duration in seconds nighthawk default format
	testConfig := &api.NighthawkConfig{
		Thread:            1,
		DurationInSeconds: 5,
		QPS:               1,
		URL:               "https://www.github.com",
	}

	result, err := api.NighthawkRun(testConfig)

	if err != nil {
		msg := "Failed to run load-test"
		err = errors.Wrapf(err, msg)
		log.Fatal(err)
	}

	fmt.Printf(string(result))

}
