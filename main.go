package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/jinycoo/jiny/cmd"
)

func main() {
	command := &cobra.Command{
		Use:   "jiny",
		Short: "快速创建基于Jinygo框架的Golang项目，及部署配置",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommands(command)

	if err := command.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
