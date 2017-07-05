package GoDE

import (
	"fmt"
)

var (
	version = SemVer{uint(0), uint(1), uint(0)}
)

type SemVer struct {
	Major uint
	Minor uint
	Patch uint
}

func Version() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}
