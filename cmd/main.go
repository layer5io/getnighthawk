package main

import (
	"fmt"
	"os"
	// "os/exec"
	// "strconv"

	// "github.com/pkg/errors"
	"github.com/layer5io/nighthawk-go/api"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)

}

// func runNighthawk(duration int, qps int, c int, url string) []byte {

// 	out, err := exec.Command("docker", "run", "envoyproxy/nighthawk-dev:latest", "nighthawk_client", "--rps "+strconv.Itoa(qps), "--concurrency "+strconv.Itoa(c), "--duration "+strconv.Itoa(duration), url, "--output-format json").Output()
// 	if err != nil {
// 		err = errors.Wrapf(err, "unable to run nighthawk")
// 		log.Fatal(err)
// 	}

// 	return out

// }
func main() {

	//Duration in seconds nighthawk default format
	config := &api.NighthawkConfig{
		Thread:            1,
		DurationInSeconds: 5,
		QPS:               10,
		URL:               "https://www.github.com",
	}

	result, _ := api.NighthawkRun(config)
	fmt.Printf(string(result))

}
