package main

  import (

 	"os/exec"
 	"fmt"
  "os"
  "strconv"

 	log "github.com/sirupsen/logrus"
  "github.com/pkg/errors"
 )

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  log.SetOutput(os.Stdout)

}

func runNighthawk(duration int, qps int, c int, url string) []byte {

  out, err := exec.Command("docker", "run", "envoyproxy/nighthawk-dev:latest", "nighthawk_client", "--rps " + strconv.Itoa(qps), "--concurrency " + strconv.Itoa(c), "--duration " + strconv.Itoa(duration), url, "--output-format json").Output()
 	if err != nil {
    log.Fatal(err)
    err = errors.Wrapf(err, "unable to run nighthawk")
 		log.Fatal(err)
  }

  return out

}
func main() {

  //Duration in seconds nighthawk default format
  var duration int = 15
  var qps int = 50
  var c int = 10
  var url string = "https://www.github.com"

 	result := runNighthawk(duration, qps, c, url)

 	fmt.Printf(string(result))

}