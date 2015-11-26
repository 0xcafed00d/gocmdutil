package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tarm/serial"
)

type Config struct {
	Help         bool
	SerialDevice string
	SerialSpeed  int
}

var config Config

func init() {
	flag.BoolVar(&config.Help, "h", false, "display help")
	flag.StringVar(&config.SerialDevice, "d", "", "serial device name")
	flag.IntVar(&config.SerialSpeed, "b", 9600, "serial baudrate")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: serialmon [options]")
		flag.PrintDefaults()
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func openComms(config Config) io.ReadWriteCloser {

	if len(config.SerialDevice) > 0 && config.SerialSpeed != 0 {
		serialcfg := serial.Config{Name: config.SerialDevice, Baud: config.SerialSpeed}
		port, err := serial.OpenPort(&serialcfg)
		exitOnError(err)
		return port
	}

	fmt.Fprintln(os.Stderr, "comms port incorrectly specified")
	flag.Usage()
	os.Exit(1)

	return nil
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 0 || config.Help {
		flag.Usage()
		os.Exit(1)
	}

	comms := openComms(config)
	defer comms.Close()
}
