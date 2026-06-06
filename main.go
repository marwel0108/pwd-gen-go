package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	length int
	noDigits bool
	noSpecial bool
)

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
	fmt.Printf("Length: %v\n", length)
	fmt.Printf("No digits: %t\n", noDigits)
	fmt.Printf("No special chars: %t\n", noSpecial)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}