package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var cusArgsCheckCmd = &cobra.Command{
	Use: "curArgs",
	// 参数校验
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cmd curArgs start")
		fmt.Println(args)
		fmt.Println("run cmd curArgs end")
	},
}

var argsNoCheckCmd = &cobra.Command{
	Use:  "argsNoCheck",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cmd argsNoCheck start")
		fmt.Println(args)
		fmt.Println("run cmd argsNoCheck end")
	},
}

var argsLimitCheckCmd = &cobra.Command{
	Use:  "argsLimitCheck",
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cmd argsLimitCheck start")
		fmt.Println(args)
		fmt.Println("run cmd argsLimitCheck end")
	},
}

var argsValidCheckCmd = &cobra.Command{
	Use:       "argsValidCheck",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"a", "b", "c"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cmd argsValidCheck start")
		fmt.Println(args)
		fmt.Println("run cmd argsValidCheck end")
	},
}
