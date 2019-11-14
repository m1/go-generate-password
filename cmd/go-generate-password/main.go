package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/m1/go-generate-password/generator"
)

var (
	rootCmd                    *cobra.Command
	length                     int
	characterSet               string
	includeSymbols             bool
	includeNumbers             bool
	includeLowercaseLetters    bool
	includeUppercaseLetters    bool
	excludeSimilarCharacters   bool
	excludeAmbiguousCharacters bool
	times                      int
)

func main() {
	rootCmd = &cobra.Command{
		Run:   generate,
		Use:   "go-generate-password",
		Short: "go-generate-password is a password generating engine.",
		Long:  "go-generate-password is a password generating engine written in Go.",
	}

	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", generator.DefaultConfig.Length, "Length of the password")
	rootCmd.PersistentFlags().StringVar(&characterSet, "characters", generator.DefaultConfig.CharacterSet, "Character set for the config")
	rootCmd.PersistentFlags().BoolVar(&includeSymbols, "symbols", generator.DefaultConfig.IncludeSymbols, "Include symbols")
	rootCmd.PersistentFlags().BoolVar(&includeNumbers, "numbers", generator.DefaultConfig.IncludeNumbers, "Include numbers")
	rootCmd.PersistentFlags().BoolVar(&includeLowercaseLetters, "lowercase", generator.DefaultConfig.IncludeLowercaseLetters, "Include lowercase letters")
	rootCmd.PersistentFlags().BoolVar(&includeUppercaseLetters, "uppercase", generator.DefaultConfig.IncludeSymbols, "Include uppercase letters")
	rootCmd.PersistentFlags().BoolVar(&excludeSimilarCharacters, "exclude-similar", generator.DefaultConfig.ExcludeSimilarCharacters, "Exclude similar characters")
	rootCmd.PersistentFlags().BoolVar(&excludeAmbiguousCharacters, "exclude-ambiguous", generator.DefaultConfig.ExcludeAmbiguousCharacters, "Exclude ambiguous characters")
	rootCmd.PersistentFlags().IntVarP(&times, "times", "n", 1, "How many passwords to generate")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func generate(_ *cobra.Command, args []string) {
	config := generator.Config{
		Length:                     length,
		CharacterSet:               characterSet,
		IncludeSymbols:             includeSymbols,
		IncludeNumbers:             includeNumbers,
		IncludeLowercaseLetters:    includeLowercaseLetters,
		IncludeUppercaseLetters:    includeUppercaseLetters,
		ExcludeSimilarCharacters:   excludeSimilarCharacters,
		ExcludeAmbiguousCharacters: excludeAmbiguousCharacters,
	}
	g, err := generator.New(&config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	pwds, err := g.GenerateMany(times)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, pwd := range pwds {
		fmt.Println(pwd)
	}
}
