package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfg            config.Config
	allowedMethods = map[string]struct{}{
		"add":    {},
		"remove": {},
	}
	network           string
	login             string
	errNotFoundMethod = errors.New("method not found")
	successResponse   = []byte("{\"ok\":true}")
)

var rootCmd = &cobra.Command{
	Use:   "cli-anti-bruteforce",
	Short: "Anti bruteforce client",
}

func init() {
	var err error
	cfg, err = config.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(blackListCommand)
	rootCmd.AddCommand(whiteListCommand)
	rootCmd.AddCommand(resetCommand)

	whiteListCommand.Flags().StringVar(&network, "n", "", "Ip addr with mask")
	blackListCommand.Flags().StringVar(&network, "n", "", "Ip addr with mask")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkResponse(b []byte, err error) error {
	if err != nil {
		return err
	}

	if !bytes.Equal(b, successResponse) {
		return fmt.Errorf("response is not ok. Body: %s", b)
	}

	return nil
}
