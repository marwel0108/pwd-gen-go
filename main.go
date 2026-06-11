package main

import (
	"errors"
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
	Run: run,
}

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of generated password")
	rootCmd.Flags().BoolVarP(&noDigits, "no-digits", "d", false, "Defines whether ot not the generated password will contain numeric digits")
	rootCmd.Flags().BoolVarP(&noSpecial, "no-special", "s", false, "Defines whether ot not the generated password will contain special characters")
}

func generatePassword() (generatedPassword string, err error) {


	if length < 8 || length > 64 {
		return "", errors.New("Length must be between 8 and 64")
	}

	genPasswordRuneSlice := make([]rune, 0, length)

	var setsPermitidos [][]rune
	setsPermitidos = append(setsPermitidos, alphabet)

	if !noDigits {
		setsPermitidos = append(setsPermitidos, digits)
	}
	if !noSpecial {
		setsPermitidos = append(setsPermitidos, special)
	}

	for len(genPasswordRuneSlice) < length {
		setElegido := setsPermitidos[rand.Intn(len(setsPermitidos))]
		
		charCandidate := setElegido[rand.Intn(len(setElegido))]
		
		genPasswordRuneSlice = append(genPasswordRuneSlice, charCandidate)
	}

	return string(genPasswordRuneSlice), err
}

func run(cmd *cobra.Command, args []string) {

	generatedPassword, err := generatePassword()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(generatedPassword)

}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}