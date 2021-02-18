// +build go1.16

package dummy

import _ "embed"

//go:embed hello.txt
var X string
