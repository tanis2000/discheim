package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tanis2000/discheim/service"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "discheim",
	Short: "discheim is a Discord bot to control Valheim servers",
	Long: `Your Discord bot for Valheim servers management`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := service.NewServer()
		err := srv.Restart()
		if err != nil {
			log.Printf(err.Error())
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}