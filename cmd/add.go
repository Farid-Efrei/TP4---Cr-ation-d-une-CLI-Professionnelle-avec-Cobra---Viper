package cmd

import (
	"TP4/config"
	"TP4/internal/models"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addName, addEmail string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := viper.GetViper()
		st, err := config.NewStoreFromViper(v)
		if err != nil {
			return err
		}
		contact := &models.Contact{
			Name:  addName,
			Email: addEmail,
		}
		if err := st.Add(contact); err != nil {
			return err
		}
		fmt.Printf("Contact ajouté (ID=%d)\n", contact.ID)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&addName, "nom", "n", "", "Nom (required)")
	addCmd.Flags().StringVarP(&addEmail, "email", "e", "", "Email (required)")
	addCmd.MarkFlagRequired("nom")
	addCmd.MarkFlagRequired("email")
	rootCmd.AddCommand(addCmd)
}
