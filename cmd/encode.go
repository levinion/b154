package cmd

import (
	"github.com/levinion/b154/util"

	"github.com/levinion/b154/pkg"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use: "encode",
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		num, err := cmd.Flags().GetInt("num")
		checkErr(err)
		Encode(path, num)
	},
	Aliases: []string{"e"},
}

func Encode(path string, num int) {
	if util.IsDir(path) {
		pkg.EncodeDir(path, num)
	} else {
		pkg.EncodeFile(path, num)
	}
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
