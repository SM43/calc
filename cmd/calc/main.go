package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	calc "github.com/sm43/calc"
	status "github.com/sm43/calc/gen/status"
	v1 "github.com/sm43/calc/v1"
	v1svc "github.com/sm43/calc/v1/gen/calc"
	v2 "github.com/sm43/calc/v2"
	v2svc "github.com/sm43/calc/v2/gen/calc"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[calc] ", log.Ltime)
	}

	// Initialize the services.
	var (
		statusSvc status.Service
		v1Svc     v1svc.Service
		v2Svc     v2svc.Service
	)
	{
		statusSvc = calc.NewStatus(logger)
		v1Svc = v1.NewCalc(logger)
		v2Svc = v2.NewCalc(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		statusEndpoints *status.Endpoints
		v1SvcEndpoints  *v1svc.Endpoints
		v2SvcEndpoints  *v2svc.Endpoints
	)
	{
		statusEndpoints = status.NewEndpoints(statusSvc)
		v1SvcEndpoints = v1svc.NewEndpoints(v1Svc)
		v2SvcEndpoints = v2svc.NewEndpoints(v2Svc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, ":80")
			}
			handleHTTPServer(ctx, u, statusEndpoints, v1SvcEndpoints, v2SvcEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
