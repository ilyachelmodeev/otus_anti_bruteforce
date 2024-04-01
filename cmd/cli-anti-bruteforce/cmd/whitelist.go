package cmd

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/httpClient"
	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/server"
	"github.com/spf13/cobra"
)

var whiteListCommand = &cobra.Command{
	Use:   "whitelist",
	Short: "Add/remove from whitelist",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		method := args[0]
		if _, ok := allowedMethods[method]; !ok {
			return errNotFoundMethod
		}

		var err error
		var b []byte
		hc := httpClient.New(cfg.Host)
		vs := url.Values{}
		vs.Set(server.IPField, network)

		if method == "add" {
			b, err = hc.Post(context.Background(), "whiteList", vs)
		} else {
			b, err = hc.Delete(context.Background(), "whiteList", vs)
		}

		if err := checkResponse(b, err); err != nil {
			return err
		}

		fmt.Println("Addr successfully whitelisted")
		return nil
	},
}
