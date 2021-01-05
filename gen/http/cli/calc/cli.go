// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sm43/calc/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	statusc "github.com/sm43/calc/gen/http/status/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `status status
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` status status` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		statusFlags = flag.NewFlagSet("status", flag.ContinueOnError)

		statusStatusFlags = flag.NewFlagSet("status", flag.ExitOnError)
	)
	statusFlags.Usage = statusUsage
	statusStatusFlags.Usage = statusStatusUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "status":
			svcf = statusFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "status":
			switch epn {
			case "status":
				epf = statusStatusFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "status":
			c := statusc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "status":
				endpoint = c.Status()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// statusUsage displays the usage of the status command and its subcommands.
func statusUsage() {
	fmt.Fprintf(os.Stderr, `The status service provides status of the service.
Usage:
    %s [globalflags] status COMMAND [flags]

COMMAND:
    status: Status implements status.

Additional help:
    %s status COMMAND --help
`, os.Args[0], os.Args[0])
}
func statusStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] status status

Status implements status.

Example:
    `+os.Args[0]+` status status
`, os.Args[0])
}
