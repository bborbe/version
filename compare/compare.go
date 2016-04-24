package compare

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bborbe/log"
	"github.com/bborbe/version"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
	SUCCESS            = 0
	FAILURE            = 1
)

type Compare func(version.Version, version.Version) bool

func Main(compare Compare) {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")

	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	writer := os.Stdout
	returnCode, err := do(writer, compare, flag.Args())
	if err != nil {
		logger.Fatal(err)
		returnCode = 1
	}
	logger.Close()
	os.Exit(returnCode)
}

func do(writer io.Writer, compare Compare, versions []string) (int, error) {
	logger.Debug("start")
	if len(versions) != 2 {
		fmt.Fprintf(writer, "can only compare to versions\n")
		return FAILURE, nil
	}
	if compare(version.Version(versions[0]), version.Version(versions[1])) {
		return SUCCESS, nil
	} else {
		return FAILURE, nil
	}
}
