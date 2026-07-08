package main

import (
	"fmt"
	"os"
)

func ValidateResult(result *Result) {

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "========== Data Validation ==========")

	fmt.Fprintln(os.Stderr, "Players :", len(result.Players))

	fmt.Fprintln(os.Stderr, "Kills   :", len(result.Kills))

	fmt.Fprintln(os.Stderr, "Damage  :", len(result.Damages))

	fmt.Fprintln(os.Stderr, "States  :", len(result.PlayerStates))

	fmt.Fprintln(os.Stderr, "Rounds  :", len(result.Rounds))
}
