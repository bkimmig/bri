package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

var timeDilationLong = `A command for calculating the exchange rate of traveling on a ship at some fraction of the speed of light for a given amount of time.

The idea is you sacrifice gamma*X years of your life on a ship to move Y years into the future on Earth. Given a fraction of the speed of light, 
you can travel at and the amount of time you'd like to pass on earth this will give you the amount of time you'd need to spend on the ship. 
`

func timeDialtionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "timedilation",
		Short: "A command for calculating time dilation if you traveled on a ship a given fraction of the speed of light",
		Long:  timeDilationLong,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := cmd.Flags().GetFloat64("percent-c")
			if err != nil {
				fmt.Printf("error getting speed of light fraction: %v\n", err)
				return
			}

			earthYears, err := cmd.Flags().GetFloat64("earth-years")
			if err != nil {
				fmt.Printf("error getting earth years: %v\n", err)
				return
			}

			gamma := 1 / math.Sqrt(1-math.Pow(c, 2))
			shipYears := earthYears / gamma

			fmt.Printf("\nYou'd need to spend %.1f years on ship moving at %.5fc to pass %.1f years on earth", shipYears, c, earthYears)

		},
	}

	cmd.Flags().Float64P("percent-c", "c", 0.997, "Fraction of the speed of light")
	cmd.Flags().Float64P("earth-years", "e", 10.1, "Number of earth year's you'd like to pass")

	return cmd
}
