package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tanis2000/discheim/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var discordToken string

var rootCmd = &cobra.Command{
	Use:   "discheim",
	Short: "discheim is a Discord bot to control Valheim servers",
	Long: `Your Discord bot for Valheim servers management`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := service.NewServer()
		bot := service.NewBot(discordToken, srv)
		err := bot.Create()
		if err != nil {
			log.Fatal(err.Error())
		}
		err = bot.Connect()
		if err != nil {
			log.Fatal(err.Error())
		}
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<- sc
		err = bot.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&discordToken, "token", "", "Discord API token")
	err := viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err.Error())
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}