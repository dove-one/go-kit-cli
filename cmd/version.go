package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-kit-cli/constant"
	"log"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current binary version info",
	Long:  `All software has versions. This is go-kit's.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func PrintVersion() {
	log.Println(fmt.Sprintf(constant.PROJECTNAME+":%q\n", constant.VERSION))
}
