package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println("viper: ", cmd.PersistentFlags().Lookup("viper").Value)
		fmt.Println("author: ", cmd.PersistentFlags().Lookup("author").Value)
		fmt.Println("license: ", cmd.PersistentFlags().Lookup("license").Value)
		fmt.Println("config: ", cmd.PersistentFlags().Lookup("config").Value)
		fmt.Println("source: ", cmd.Flags().Lookup("source").Value)
		fmt.Println("-------------------------------------")
		fmt.Println("viper author: ", viper.GetString("author"))
		fmt.Println("viper license: ", viper.GetString("license"))
		fmt.Println("root cmd run end")
	},
	TraverseChildren: true, // 向子命令传递所有父命令解析的标志。
}

func Execute() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(cusArgsCheckCmd)
	rootCmd.AddCommand(argsNoCheckCmd)
	rootCmd.AddCommand(argsLimitCheckCmd)
	rootCmd.AddCommand(argsValidCheckCmd)
	rootCmd.Execute()
}

var cfgFile string
var userLicense string

func init() {
	cobra.OnInitialize(initConfig)
	// 按名称接受命令行参数
	rootCmd.PersistentFlags().Bool("viper", false, "")
	// 指定flag缩写
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "作者")
	// 通过 指针 赋值到字段中
	rootCmd.PersistentFlags().StringVar(&userLicense, "license", "NO LICENSE", "证书")
	// 通过 指针 赋值到字段中 并指定缩写
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "NO CONFIG", "配置文件")
	// 本地标识
	rootCmd.Flags().StringP("source", "s", "NO SOURCE", "")

	if err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author")); err != nil {
		log.Println(err)
		return
	}

	if err := viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license")); err != nil {
		log.Println(err)
		return
	}

	viper.SetDefault("author", "范伟锋")
	viper.SetDefault("license", "github")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	// 检查环境变量，将配置键加载到viper中
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	fmt.Println("Using config file: ", viper.ConfigFileUsed())
}
