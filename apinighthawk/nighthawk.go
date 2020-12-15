// Package apinighthawk defines nighthawk runner and config
package apinighthawk

import (
	"fmt"
	"net/url"
	"os/exec"
	"strconv"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// NighthawkConfig describes the configuration structure for loadtest
type NighthawkConfig struct {
	Thread            int
	DurationInSeconds float64
	QPS               float64
	URL               string
}

// NighthawkRun function runs the nighthawk loadtest
func NighthawkRun(config *NighthawkConfig) ([]byte, error) {
	rURL, _ := url.Parse(config.URL)
	if !rURL.IsAbs() {
		err := fmt.Errorf("please give a valid URL %s", config.URL)
		log.Error(err)
		return nil, err
	}

	duration := strconv.FormatFloat(config.DurationInSeconds, 'f', -1, 64)
	qps := strconv.FormatFloat(config.QPS, 'f', -1, 64)
	c := strconv.Itoa(config.Thread)

	args := []string{"--rps " + qps,
		"--connections " + c,
		"--duration " + duration,
		rURL.String(),
		"--output-format experimental_fortio_pedantic"}

	log.Info("Received arguments for run", args)

	out, err := exec.Command("nighthawk_client",
		"--rps "+qps,
		"--concurrency 1",
		"--connections "+c,
		"--duration "+duration,
		rURL.String(),
		"--output-format experimental_fortio_pedantic").Output()

	if err != nil {
		msg := "Unable to run load-test"
		err = errors.Wrapf(err, msg)
		log.Error(err)
		return nil, err
	}

	return out, nil
}
