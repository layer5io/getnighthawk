package api

import (
	"fmt"
	"os/exec"
	"strconv"
	"net/url"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type NighthawkConfig struct {
	Thread                  int
	DurationInSeconds       int
	QPS                     int
	URL                     string
}

func NighthawkRun (config *NighthawkConfig) ([]byte, error) {

	// nighthawkDocker := "docker run envoyproxy/nighthawk-dev:latest nighthawk_client"

	rURL, _ := url.Parse(config.URL)
	if !rURL.IsAbs() {
		err := fmt.Errorf("Please give a valid URL", config.URL)
		log.Error(err)
		return nil, err
	}

	// if rURL.Port() == "" {
	// 	if rURL.Scheme == "https" {
	// 		rURL.Host += ":443"
	// 	} else {
	// 		rURL.Host += ":80"
	// 	}
	// }

	duration := strconv.Itoa(config.DurationInSeconds)
	qps := strconv.Itoa(config.QPS)
	c := strconv.Itoa(config.Thread)
	fmt.Printf(duration)
	fmt.Printf(qps)
	fmt.Printf(c)
	fmt.Printf(rURL.String())

	// args := []string{"--rps "+strconv.Itoa(qps), "--concurrency "+strconv.Itoa(c),"--duration "+strconv.Itoa(duration),rURL,"--output-format json"}
	
	// log.Debugf("Received arguments for run %v", args)
	
	out, err := exec.Command("docker", "run", "envoyproxy/nighthawk-dev:latest", "nighthawk_client","--rps "+qps,"--concurrency "+c,"--duration "+duration,rURL.String(),"--output-format json").Output()

	if err != nil {
		log.Error(err)
		err = errors.Wrapf(err, "Unable to run load-test")
		log.Error(err)
		return nil, err
	}

	return out, nil

}