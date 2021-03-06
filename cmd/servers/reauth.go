// Copyright 1999-2020. Plesk International GmbH.

package cmd

import (
	"github.com/plesk/pleskapp/plesk/actions"
	"github.com/plesk/pleskapp/plesk/config"
	"github.com/plesk/pleskapp/plesk/locales"
	"github.com/plesk/pleskapp/plesk/utils"
	"github.com/spf13/cobra"
)

var reauthCmd = &cobra.Command{
	Use:   locales.L.Get("server.reauth.cmd"),
	Short: locales.L.Get("server.reauth.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := config.GetServer(args[0])
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true
		err = actions.ServerReauth(*server)

		if err == nil {
			utils.Log.PrintL("server.reauth.success")
		}

		return err
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	ServersCmd.AddCommand(reauthCmd)
}
