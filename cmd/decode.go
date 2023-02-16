package cmd

import (
	"github.com/levinion/b154/util"

	"github.com/levinion/b154/pkg"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use: "decode",
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		num, err := cmd.Flags().GetInt("num")
		checkErr(err)
		show, err := cmd.Flags().GetBool("show")
		checkErr(err)
		Decode(path, num, show)
	},
	Aliases: []string{"d"},
}

func Decode(path string, num int, show bool) {
	if util.IsDir(path) {
		pkg.DecodeDir(path, num, show)
	} else {
		pkg.DecodeFile(path, num, show)
	}
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}
