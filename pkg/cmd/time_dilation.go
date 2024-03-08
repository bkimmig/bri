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

			years, err := cmd.Flags().GetFloat64("years")
			if err != nil {
				fmt.Printf("error getting years: %v\n", err)
				return
			}

			shipReferenceFrame, err := cmd.Flags().GetBool("ship")
			if err != nil {
				fmt.Printf("error getting reference frame: %v\n", err)
				return
			}

			gamma := 1 / math.Sqrt(1-math.Pow(c, 2))
			var dilatedYears float64
			output := "\nYou'd need to spend %.1f years on ship moving at %.5fc to pass %.1f years on earth.\n"
			if shipReferenceFrame {
				dilatedYears := years * gamma
				fmt.Printf(output, years, c, dilatedYears)
				return
			}

			dilatedYears = years / gamma
			fmt.Printf(output, dilatedYears, c, years)

		},
	}

	cmd.Flags().Float64P("percent-c", "c", 0.997, "Fraction of the speed of light")
	cmd.Flags().Float64P("years", "y", 10.0, "Number of years you'd like to pass")
	cmd.Flags().BoolP("ship", "s", false, "Use reference frame of the ship for years")

	return cmd
}
