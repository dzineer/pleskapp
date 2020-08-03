// Copyright 1999-2020. Plesk International GmbH.

package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/plesk/pleskapp/plesk/actions"
	"github.com/plesk/pleskapp/plesk/api/factory"
	"github.com/plesk/pleskapp/plesk/config"
	"github.com/plesk/pleskapp/plesk/locales"
	"github.com/plesk/pleskapp/plesk/utils"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   locales.L.Get("database.deploy.cmd"),
	Short: locales.L.Get("database.deploy.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := config.GetServer(args[0])
		if err != nil {
			return err
		}

		domain, err := config.GetDomain(*server, args[1])
		if err != nil {
			return err
		}

		// TODO: Restrict to one domain
		db, err := actions.DatabaseFindNonLocal(factory.GetDatabaseManagement(server.GetServerAuth()), *server, args[2])
		if err != nil {
			return err
		}

		s, err := os.Stat(args[3])
		if err != nil {
			return err
		}

		if s.IsDir() {
			return errors.New(locales.L.Get("errors.path.is.directory", args[3]))
		}

		cmd.SilenceUsage = true
		path, err := actions.UploadFileToRoot(*server, *domain, true, args[3])
		if err != nil {
			utils.Log.Error(locales.L.Get("errors.execution.failed.generic", err.Error()))
			return err
		}

		fp := strings.Split(args[3], "/")
		return utils.Log.PrintSuccessOrError("database.deploy.success", nil, actions.DatabaseDeploy(*server, *domain, *db, path+"/"+fp[len(fp)-1]))
	},
	Args: cobra.ExactArgs(4),
}

func init() {
	DatabasesCmd.AddCommand(uploadCmd)
}
