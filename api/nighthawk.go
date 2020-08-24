package api

import (
	"fmt"
	"os/exec"
	"strconv"
	"net/url"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// NighthawkConfig describes the configuration structure for loadtest
type NighthawkConfig struct {
	Thread                  int
	DurationInSeconds       int
	QPS                     int
	URL                     string
}

// NighthawkRun function runs the nighthawk loadtest
func NighthawkRun (config *NighthawkConfig) ([]byte, error) {

	imageName := "envoyproxy/nighthawk-dev"
	_, err := exec.Command("docker", "inspect", imageName).Output()
	if err != nil {
		msg := "Setup nighthawk image before executing load-test"
		err = errors.Wrapf(err, msg)
		log.Error(err)
		return nil, err
	}


	rURL, _ := url.Parse(config.URL)
	if !rURL.IsAbs() {
		err := fmt.Errorf("Please give a valid URL %s", config.URL)
		log.Error(err)
		return nil, err
	}

	duration := strconv.Itoa(config.DurationInSeconds)
	qps := strconv.Itoa(config.QPS)
	c := strconv.Itoa(config.Thread)

	args := []string{"--rps "+qps, "--concurrency "+c,"--duration "+duration,rURL.String(),"--output-format json"}
	
	log.Info("Received arguments for run", args)

	out, err := exec.Command("docker", "run",
	 						"envoyproxy/nighthawk-dev:latest",
	  						"nighthawk_client",
	 						"--rps "+qps,
	  						"--concurrency "+c,
	  						"--duration "+duration,
	 						rURL.String(),
	 						"--output-format json").Output()

	if err != nil {
		msg := "Unable to run load-test"
		err = errors.Wrapf(err, msg)
		log.Error(err)
		return nil, err
	}

	return out, nil

}