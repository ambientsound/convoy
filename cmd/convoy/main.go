package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ambientsound/convoy/pkg/conversions"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("missing input value")
	}

	val, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		return err
	}

	converters := make([]conversions.Converter, 0)
	converters = append(converters, conversions.Temperatures...)
	converters = append(converters, conversions.Masses...)

	in := conversions.Amount{
		Value: val,
	}
	for _, c := range converters {
		conversion := c(in)
		fmt.Printf("%15s = %s\n", conversion.Input, conversion.Output)
	}

	return nil
}
