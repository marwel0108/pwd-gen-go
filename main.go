package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

var (
	length int
	noDigits bool
	noSpecial bool
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("0123456789")
var special = []rune("!@#$%^&*()-_=+[]{}|;:',.<>/?`~\"\\")

var rootCmd = &cobra.Command{
	Use: "pwd-gen",
	Run: generatePassword,
}

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of generated password")
	rootCmd.Flags().BoolVarP(&noDigits, "no-digits", "d", false, "Defines whether ot not the generated password will contain numeric digits")
	rootCmd.Flags().BoolVarP(&noSpecial, "no-special", "s", false, "Defines whether ot not the generated password will contain special characters")
}

func generatePassword(cmd *cobra.Command, args []string) {

	if length < 8 || length > 64 {
		log.Fatalf("Length must be between 8 and 64")
	}

	genPassword := make([]rune, 0, length)

	var setsPermitidos [][]rune
	setsPermitidos = append(setsPermitidos, alphabet)

	if !noDigits {
		setsPermitidos = append(setsPermitidos, digits)
	}
	if !noSpecial {
		setsPermitidos = append(setsPermitidos, special)
	}

	for len(genPassword) < length {
		setElegido := setsPermitidos[rand.Intn(len(setsPermitidos))]
		
		charCandidate := setElegido[rand.Intn(len(setElegido))]
		
		genPassword = append(genPassword, charCandidate)
	}

	fmt.Println(string(genPassword))

}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}