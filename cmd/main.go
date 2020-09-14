package main

import (
	"fmt"
	"os"
	"encoding/json"

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

	// Output to only for logs above warn level
	log.SetLevel(log.WarnLevel)
}
func main() {
	// Duration in seconds nighthawk default format
	testConfig := &apinighthawk.NighthawkConfig{
		Thread:            1,
		DurationInSeconds: 5,
		QPS:               1,
		URL:               "https://www.github.com",
	}

	result, err := apinighthawk.NighthawkRun(testConfig)

	if err != nil {
		msg := "Failed to run load-test"
		err = errors.Wrapf(err, msg)
		log.Fatal(err)
	}

	fmt.Print(string(result))
	// res1 := string(result)

	var result1 periodic.RunnerResults
	// var bd []byte

	// hres, _ := res1.(*fhttp.HTTPRunnerResults)
	// bd, err = json.Marshal(hres)
	// result1 = hres.Result()

	err = json.Unmarshal([]byte(result), &result1)

	if err != nil {
		err = errors.Wrap(err, "Error while unmarshaling  Nighthawk results to the FortioHTTPRunner")
		// logrus.Error(err)
		log.Fatal(err)
	}

	resultsMap := map[string]interface{}{}
	err = json.Unmarshal(result, &resultsMap)

	if err != nil {
		err = errors.Wrap(err, "Error while unmarshaling Nighthawk results to map")
		// log.Error(err)
		log.Fatal(err)
	}

	log.Debugf("Mapped version of the test: %+#v", resultsMap)


}
