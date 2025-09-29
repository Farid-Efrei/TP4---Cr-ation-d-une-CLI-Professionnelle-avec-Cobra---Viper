package cmd

import (
	"TP4/config"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(getAllCmd)
}

var getAllCmd = &cobra.Command{
	Use:   "get_all",
	Short: "Lister tous les contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := viper.GetViper()
		st, err := config.NewStoreFromViper(v)
		if err != nil {
			return err
		}
		contacts, err := st.GetAll()
		if err != nil {
			return err
		}
		for _, c := range contacts {
			fmt.Printf("ID: %d, Nom: %s, Email: %s, Créé le: %s, Mis à jour le: %s\n",
				c.ID, c.Name, c.Email, c.CreatedAt.Format("2006-01-02 15:04:05"), c.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
		return nil
	},
}
