package cmd

import (
	"TP4/config"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] --nom new_name --email new_email",
	Short: "Mettre à jour un contact existant",
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
		contact, err := st.GetByID(id)
		if err != nil {
			return err
		}
		if contact == nil {
			return fmt.Errorf("contact avec ID %d non trouvé", id)
		}
		name, _ := cmd.Flags().GetString("nom")
		email, _ := cmd.Flags().GetString("email")
		if name != "" {
			contact.Name = name
		}
		if email != "" {
			contact.Email = email
		}
		if err := st.Update(contact.ID, contact.Name, contact.Email); err != nil {
			return err
		}
		fmt.Println("Contact mis à jour (ID=%d)\n", contact.ID)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("nom", "n", "", "Nouveau nom")
	updateCmd.Flags().StringP("email", "e", "", "Nouvel email")
	rootCmd.AddCommand(updateCmd)
}
