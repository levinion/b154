package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "b154",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `b154 help` to see how to use")
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolP("show","s",false,"")
	rootCmd.PersistentFlags().IntP("num","n",1,"")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}