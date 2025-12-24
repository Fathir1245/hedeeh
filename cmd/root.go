package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hedeeh",
	Short: "hedeeh - AI Scaffolding Tool biar ga pusing setup projek (buat yang malas aje wkwkwk)",
	Long:  `Alat bantu CLI untuk generate boilerplate project Go, JS, dan Python dengan best practice.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
