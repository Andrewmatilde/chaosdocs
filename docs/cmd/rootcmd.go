package cmd

import (
	"strings"

	_ "github.com/alecthomas/template"
	"github.com/pingcap/log"
	"github.com/spf13/cobra"
	_ "github.com/swaggo/swag"
	"go.uber.org/zap"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlzap "sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/chaos-mesh/chaosd/cmd/attack"
	"github.com/chaos-mesh/chaosd/cmd/completion"
	"github.com/chaos-mesh/chaosd/cmd/recover"
	"github.com/chaos-mesh/chaosd/cmd/search"
	"github.com/chaos-mesh/chaosd/cmd/server"
	"github.com/chaos-mesh/chaosd/cmd/version"
	"github.com/chaos-mesh/chaosd/pkg/utils"
)

var logLevel string

func setLog() {
	conf := &log.Config{Level: logLevel}
	lg, r, err := log.InitLogger(conf)
	if err != nil {
		log.Error("fail to init log", zap.Error(err))
		return
	}
	log.ReplaceGlobals(lg, r)

	// set log of controller-runtime, so that can print logs in chaos mesh
	ctrl.SetLogger(ctrlzap.New(ctrlzap.UseDevMode(true)))

	// only in debug mode print log of go.uber.org/fx
	if strings.ToLower(logLevel) == "debug" {
		utils.PrintFxLog = true
	}
}

func GetRootCMD() *cobra.Command {
	// CommandFlags are flags that used in all Commands
	var rootCmd = &cobra.Command{
		Use:   "chaosd",
		Short: "A command line client to run chaos experiment",
	}

	cobra.OnInitialize(setLog)
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "", "", "the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'")

	rootCmd.AddCommand(
		server.NewServerCommand(),
		attack.NewAttackCommand(),
		recover.NewRecoverCommand(),
		search.NewSearchCommand(),
		version.NewVersionCommand(),
		completion.NewCompletionCommand(),
	)

	_ = utils.SetRuntimeEnv()

	return rootCmd
}
