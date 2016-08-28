package compare

import (
	"flag"
	"fmt"
	"io"
	"os"

	"runtime"

	"github.com/bborbe/version"
	"github.com/golang/glog"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

type Compare func(version.Version, version.Version) bool

func Main(compare Compare) {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	writer := os.Stdout
	returnCode, err := do(writer, compare, flag.Args())
	if err != nil {
		glog.Exit(err)
	}
	os.Exit(returnCode)
}

func do(writer io.Writer, compare Compare, versions []string) (int, error) {
	if len(versions) != 2 {
		glog.V(2).Infof("arg count: %d", len(versions))
		fmt.Fprintf(writer, "can only compare to versions\n")
		glog.V(2).Infof("arg count != 2 => failure")
		return FAILURE, nil
	}
	if compare(version.Version(versions[0]), version.Version(versions[1])) {
		glog.V(2).Infof("%v %v => succes", versions[0], versions[1])
		return SUCCESS, nil
	} else {
		glog.V(2).Infof("%v %v => failure", versions[0], versions[1])
		return FAILURE, nil
	}
}
