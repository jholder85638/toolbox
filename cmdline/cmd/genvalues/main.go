// genvalues creates the values.go content for the args package.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jholder85638/toolbox/atexit"
	"github.com/jholder85638/toolbox/log/jot"
)

func main() {
	out, err := os.Create("values.go")
	jot.FatalIfErr(err)

	writeString(out, `// This file is autogenerated. Do not modify.

package cmdline

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jholder85638/toolbox/errs"
)
`)
	gen(out, "bool", "strconv.ParseBool(str)", false)
	gen(out, "int", "strconv.ParseInt(str, 0, 64)", true)
	gen(out, "int8", "strconv.ParseInt(str, 0, 8)", true)
	gen(out, "int16", "strconv.ParseInt(str, 0, 16)", true)
	gen(out, "int32", "strconv.ParseInt(str, 0, 32)", true)
	gen(out, "int64", "strconv.ParseInt(str, 0, 64)", false)
	gen(out, "uint", "strconv.ParseUint(str, 0, 64)", true)
	gen(out, "uint8", "strconv.ParseUint(str, 0, 8)", true)
	gen(out, "uint16", "strconv.ParseUint(str, 0, 16)", true)
	gen(out, "uint32", "strconv.ParseUint(str, 0, 32)", true)
	gen(out, "uint64", "strconv.ParseUint(str, 0, 64)", false)
	gen(out, "float32", "strconv.ParseFloat(str, 32)", true)
	gen(out, "float64", "strconv.ParseFloat(str, 64)", false)
	gen(out, "string", "str, error(nil)", false)
	gen(out, "time.Duration", "time.ParseDuration(str)", false)

	jot.FatalIfErr(out.Close())
	atexit.Exit(0)
}

func gen(out *os.File, dataType, parser string, needConversion bool) {
	name := dataType
	if i := strings.Index(name, "."); i != -1 {
		name = strings.ToLower(name[i+1:i+2]) + name[i+2:]
	}
	printf(out, `
// ----- %[1]s -----
type %[2]sValue %[1]s

// New%[3]sOption creates a new %[1]s Option and attaches it to this CmdLine.
func (cl *CmdLine) New%[3]sOption(val *%[1]s) *Option {
	return cl.NewOption((*%[2]sValue)(val))
}

// Set implements the Value interface.
func (val *%[2]sValue) Set(str string) error {
	v, err := %[4]s
	*val = %[2]sValue(v)
	return errs.Wrap(err)
}

// String implements the Value interface.
func (val *%[2]sValue) String() string {
`, dataType, name, strings.Title(name), parser)
	switch dataType {
	case "string":
		writeString(out, `	return string(*val)`)
	case "time.Duration":
		writeString(out, `	return time.Duration(*val).String()`)
	default:
		writeString(out, `	return fmt.Sprintf("%v", *val)`)
	}
	printf(out, `
}

// ----- []%[1]s -----
type %[2]sArrayValue []%[1]s

// New%[3]sArrayOption creates a new []%[1]s Option and attaches it to this CmdLine.
func (cl *CmdLine) New%[3]sArrayOption(val *[]%[1]s) *Option {
	return cl.NewOption((*%[2]sArrayValue)(val))
}

// Set implements the Value interface.
func (val *%[2]sArrayValue) Set(str string) error {
	v, err := %[4]s
	*val = append(*val, `, dataType, name, strings.Title(name), parser)
	if needConversion {
		printf(out, "%s(v)", dataType)
	} else {
		writeString(out, "v")
	}
	printf(out, `)
	return errs.Wrap(err)
}

// String implements the Value interface.
func (val *%sArrayValue) String() string {
	var str string
	for _, v := range *val {
		if str == "" {
			str += ", "
		}
`, name)
	switch dataType {
	case "string":
		writeString(out, `		str += v`)
	case "time.Duration":
		writeString(out, `		str += v.String()`)
	default:
		writeString(out, `		str += fmt.Sprintf("%v", v)`)
	}
	writeString(out, `
	}
	return str
}
`)
}

func writeString(out *os.File, text string) {
	_, err := out.WriteString(text)
	jot.FatalIfErr(err)
}

func printf(out io.Writer, format string, params ...interface{}) {
	_, err := fmt.Fprintf(out, format, params...)
	jot.FatalIfErr(err)
}
