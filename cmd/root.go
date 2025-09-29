package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Un outil de gestion de contacts",
	Long:  "Un outil de gestion de contacts pour ajouter, supprimer et lister des contacts.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Fichier de configuration (par défaut dans le répertoire courant)")
}
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Recherche dans le répertoire courant
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}
	// Lecture du fichier de configuration
	viper.AutomaticEnv() // Lire les variables d'environnement qui correspondent

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("WARNING: Impossible de lire le fichier de configuration: %v\n", err)
		fmt.Println("INFO: utilisation des valeurs par défaut")
		// Valeurs par défaut
		viper.SetDefault("storage.type", "json")
		viper.SetDefault("storage.json_path", "data/contacts.json")
		viper.SetDefault("storage.gorm_path", "data/crm.db")
	} else {
		fmt.Println("INFO: Utilisation du fichier de configuration:", viper.ConfigFileUsed())
	}
	// DEBUG : afficher la valeur lue
	fmt.Printf("DEBUG: storage.type: %s\n", viper.GetString("storage.type"))
}
