package cmd

import (
	"log"
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"github.com/seplak/i9g/internal/numeronym"
)

var rootCmd = &cobra.Command{
	Use: "i9g",
	Short: "Turn strings into numeronyms",
	Long: `Turn strings into numeronyms, but really,
		please don't do this.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}
		toConvert := strings.Join(args, " ")
		num := numeronym.Numeronym{Text: toConvert}
		fmt.Println(num.Convert())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}