package cmd

import (
	"TP4/config"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Supprimer un contact par son ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		v := viper.GetViper()
		st, err := config.NewStoreFromViper(v)
		if err != nil {
			return err
		}
		if err := st.Delete(id); err != nil {
			return err
		}
		fmt.Printf("Contact supprimé (ID=%d)\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
