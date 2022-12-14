package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	AgentCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [command]", os.Args[0]),
		Short: "opsAgent at your service.",
		Long: `
The Agent faithfully collects events and metrics and brings them
to Server on your behalf so that you can do something useful with your
monitoring and performance data.`,
		SilenceUsage: true,
	}
	confFilePath string
	flagNoColor  bool
)

func init() {
	AgentCmd.PersistentFlags().StringVarP(&confFilePath, "cfgpath", "c", "", "path to directory containing opsAgent.conf")
	AgentCmd.PersistentFlags().BoolVarP(&flagNoColor, "no-color", "n", false, "disable color output")
}
