package cmd

import (
	"context"
	"path/filepath"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/satrap-illustrations/zs/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func Execute() error {
	cobra.OnInitialize(initConfig)

	rootCmd := &cobra.Command{
		Version: "v0.0.1",
		Use:     "zs",
		Short:   "Zendesk Search",
		Long: `Zendesk Search (zs)

It searches Zendesk.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(tui.InitialModel())
			if _, err := p.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config/zs/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	return rootCmd.ExecuteContext(context.Background())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "zs"))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("zs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config", "file", viper.ConfigFileUsed(), "error", err)
		return
	}
}