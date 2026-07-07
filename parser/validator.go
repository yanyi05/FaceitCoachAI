package main

import "fmt"

func ValidateResult(result *Result) {

	fmt.Println("========== Validation ==========")

	fmt.Printf("Players: %d\n", len(result.Players))

	fmt.Printf("PlayerStates: %d\n", len(result.PlayerStates))

	fmt.Printf("Kills: %d\n", len(result.Kills))

	fmt.Printf("Damages: %d\n", len(result.Damages))

	fmt.Printf("Rounds: %d\n", len(result.Rounds))

	fmt.Println("===============================")
}
