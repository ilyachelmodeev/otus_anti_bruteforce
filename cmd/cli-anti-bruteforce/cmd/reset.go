package cmd

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/httpClient"
	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/server"
	"github.com/spf13/cobra"
)

var resetCommand = &cobra.Command{
	Use:   "reset block",
	Short: "Reset block by login and ip",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var b []byte
		hc := httpClient.New(cfg.Host)
		vs := url.Values{}
		vs.Set(server.IPField, network)
		vs.Set(server.LoginField, login)

		b, err = hc.Get(context.Background(), "reset", vs)

		if err != nil {
			return err
		}

		if !bytes.Equal(b, successResponse) {
			return fmt.Errorf("response is not ok. Body: %s", b)
		}

		fmt.Println("Block successfully reset")
		return nil
	},
}

func init() {
	resetCommand.Flags().StringVar(&network, "n", "", "Ip addr")
	resetCommand.Flags().StringVar(&login, "l", "", "Login")
}
