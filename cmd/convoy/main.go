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

	conv := []conversions.Converter{
		conversions.CelsiusFahrenheit,
		conversions.CelsiusKelvin,
		conversions.FahrenheitCelsius,
		conversions.FahrenheitKelvin,
		conversions.KelvinCelsius,
		conversions.KelvinFahrenheit,
	}

	in := conversions.Amount{
		Value: val,
	}
	for _, c := range conv {
		conversion := c(in)
		fmt.Printf("%15s = %s\n", conversion.Input, conversion.Output)
	}

	return nil
}
