/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Rascal0814/boot/tools/boot/crud"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// crudCmd represents the crud command
var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "gen database crud file",
	Long:  `The dsn and model package name should be specified when using it,and the output path can also be specified.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dsn, err := cmd.Flags().GetString("dsn")
		if err != nil {
			return errors.New("dsn cannot be empty")
		}
		out, _ := cmd.Flags().GetString("output")
		m, err := cmd.Flags().GetString("pkg")
		if err != nil {
			return errors.New("model package cannot be empty")
		}
		curd, err := crud.NewGenCurd(dsn, m, out)
		if err != nil {
			return err
		}

		err = curd.Gen()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(crudCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//crudCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	crudCmd.Flags().StringP("dsn", "", "", "")
	crudCmd.Flags().StringP("output", "o", "", "generate file output path")
	crudCmd.Flags().StringP("pkg", "", "", "model package path")
}
