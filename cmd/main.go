package main

import (
	"encoding/json"
	"os"

	"github.com/layer5io/nighthawk-go/apinighthawk"
	// "fortio.org/fortio/fhttp"
	"fortio.org/fortio/periodic"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
}
func main() {
	// Duration in seconds nighthawk default format
	testConfig := &apinighthawk.NighthawkConfig{
		Thread:            2,
		DurationInSeconds: 5,
		QPS:               2,
		URL:               "https://www.github.com",
	}

	result, err := apinighthawk.NighthawkRun(testConfig)

	if err != nil {
		msg := "Failed to run load-test"
		err = errors.Wrapf(err, msg)
		log.Fatal(err)
	}

	var result1 periodic.RunnerResults

	err = json.Unmarshal(result, &result1)

	if err != nil {
		err = errors.Wrap(err, "Error while unmarshaling  Nighthawk results to the FortioHTTPRunner")
		log.Fatal(err)
	}

	resultsMap := map[string]interface{}{}
	err = json.Unmarshal(result, &resultsMap)

	if err != nil {
		err = errors.Wrap(err, "Error while unmarshaling Nighthawk results to map")
		log.Fatal(err)
	}

	log.Infof("Mapped version of the test: %+#v", resultsMap)
}
