package main

  import (

 	"os/exec"
 	"fmt"

 	log "github.com/sirupsen/logrus"
  "github.com/pkg/errors"
 )

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  log.SetOutput(os.Stdout)

}

func wrkRun() []byte {
  out, err := exec.Command("docker", "run", "envoyproxy/nighthawk-dev:latest","nighthawk_client", "--help")
 	if err != nil {
 		log.Fatal(err)
  }
  result, err := out.Output
  if err != nil {
    log.Fatal(err)
  }
  return result
}
func main() {
 	result := wrkRun()
 	fmt.Printf(string(result))
}