package cmd 

import (

	"fmt"
	"github.com/spf13/cobra"
	"crypto/rand"
    "math/big"
)


var length int 
var symbols bool 
var uppercase bool 



var generateCmd = &cobra.Command {
	Use: "generate", 
	Short: "Generate random password, from given constraints. ",
	Long: "",
	RunE: func (cmd *cobra.Command, args []string) error {
        // 👇 Read flags *inside* Run/RunE
		var err error

        length, err = cmd.Flags().GetInt("length")
        if err != nil { return err }

        symbols, err = cmd.Flags().GetBool("symbols")
        if err != nil { return err }

        uppercase, err = cmd.Flags().GetBool("uppercase")
        if err != nil { return err }

		var pwd string
		pwd, err = generatePassword(length, symbols, uppercase) 
		fmt.Println(pwd)
        return nil
	},
}


func init (){ 
	rootCmd.AddCommand(generateCmd)
    generateCmd.Flags().IntVarP(&length, "length", "l", 12, "Password length")
    generateCmd.Flags().BoolVarP(&symbols, "symbols", "s",false, "Include symbols")
    generateCmd.Flags().BoolVarP(&uppercase, "uppercase","u", false, "Include uppercase letters")
}


func generatePassword(length int, useSymbols bool, useUppercase bool) (string, error) {
	var numbers = "123456789"
	var lowercase = "abcdefghijklmnopqrstuvwxyz"
	var uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var symbols = "!@#$%^&*()-_=+[]{};:,.?/\\|"

	var pwdsource = numbers+lowercase
	if useSymbols {
		pwdsource += symbols
	}
	if useUppercase {
		pwdsource += uppercase
	}

	var lasaddressofpasswordsource = len(pwdsource)

	var finalpassword string

	for i :=0; i< length; i++{

		randomaddr, e := rand.Int(rand.Reader, big.NewInt(int64(lasaddressofpasswordsource)))
		if e != nil {return "", e}
		finalpassword += string(pwdsource[randomaddr.Int64()])
	}

	return finalpassword, nil
}