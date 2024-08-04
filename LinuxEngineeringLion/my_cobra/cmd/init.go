package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "add",
	Short: "short init",
	Long:  "long init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init cmd run start")
		fmt.Println("viper: ", cmd.Flags().Lookup("viper").Value)
		fmt.Println("author: ", cmd.Flags().Lookup("author").Value)
		fmt.Println("license: ", cmd.Flags().Lookup("license").Value)
		fmt.Println("config: ", cmd.Flags().Lookup("config").Value)
		fmt.Println("source: ", cmd.Parent().Flags().Lookup("source").Value)
		fmt.Println("-------------------------------------")
		fmt.Println("viper author: ", viper.GetString("author"))
		fmt.Println("viper license: ", viper.GetString("license"))
		fmt.Println("init cmd run end")
	},
}
