package main

import (
	"os"

	"github.com/refgd/csocks-core"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		switch args[1] {
		case "-version", "-v", "-V":
			println(csocks.Version)
			os.Exit(0)
		}
	}

	if len(args) < 3 {
		printHelp()
	}

	quiet := false
	listenConfig := csocks.NewListenConfig()

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--quiet", "-q":
			quiet = true
		case "--http":
			listenConfig.WithHttp = true
		case "--listen", "-l":
			if i+1 < len(args) {
				listenConfig.ListenPort = args[i+1]
				i++
			} else {
				println(args[i] + " requires a port number")
				os.Exit(0)
			}
		case "--server", "-s":
			if i+1 < len(args) {
				listenConfig.ServerAddress = args[i+1]
				i++ // Move past the argument value
			} else {
				println(args[i] + " requires an IP address or hostname")
				os.Exit(0)
			}
		case "--secret":
			if i+1 < len(args) {
				listenConfig.Secret = args[i+1]
				i++ // Move past the argument value
			} else {
				println(args[i] + " requires a string")
				os.Exit(0)
			}
		case "--key", "-k":
			if i+2 < len(args) {
				listenConfig.ServerCertFile = args[i+1]
				listenConfig.ServerKeyFile = args[i+2]
				i += 2 // Move past the two arguments
			} else {
				println(args[i] + " requires paths to both a certificate file and a key file")
				os.Exit(0)
			}
		default:
			printHelp()
		}
	}

	err := csocks.StartServer(listenConfig, quiet)
	if err != nil {
		printHelp()
	}
}

func printHelp() {
	println("usage:")
	println(`    "--listen, -l Port to listen" example: "--listen 1080"`)
	println(`    "--server, -s Server address" example: "--server 127.0.0.1:1080"`)
	println(`    "--key, -k Server key" example: "--key server.crt server.key"`)
	println(`    "--secret secret key" example: "--secret 123456"`)
	println(`    "--http handle http proxy"`)
	println(`    "--quiet, -q"`)
	println(`    "--version"`)
	os.Exit(0)
}
