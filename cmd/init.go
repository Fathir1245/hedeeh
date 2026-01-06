package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"hedeeh/internal/generator"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inisialisasi projek baru",
	Run: func(cmd *cobra.Command, args []string) {
		config := generator.ProjectConfig{}

		// 1. Input Nama Projek
		nameQuestion := &survey.Input{
			Message: "Apa nama projek kamu?",
			Default: "my-app",
		}
		err := survey.AskOne(nameQuestion, &config.ProjectName)
		if err != nil {
			fmt.Println("Batal:", err)
			return
		}

		// 2. Pilih Bahasa
		langQuestion := &survey.Select{
			Message: "Pilih Bahasa Pemrograman:",
			Options: []string{"Go"},
		}
		err = survey.AskOne(langQuestion, &config.Language)
		if err != nil {
			return
		}

		// 3. Pilih Database (PENGGANTI Framework Lama)
		var dbOptions []string
		if config.Language == "Go" {
			dbOptions = []string{"MySQL", "PostgreSQL"}
		}

		dbQuestion := &survey.Select{
			Message: "Pilih Database Utama:",
			Options: dbOptions,
		}
		err = survey.AskOne(dbQuestion, &config.Database)
		if err != nil {
			return
		}

		// 4. Pilih Router / Framework HTTP
		var routerOptions []string
		if config.Language == "Go" {
			routerOptions = []string{"Standard", "Gin", "Chi"}
		}

		routerQuestion := &survey.Select{
			Message: "Pilih Router/Framework:",
			Options: routerOptions,
		}
		err = survey.AskOne(routerQuestion, &config.Router)
		if err != nil {
			return
		}

		if config.ModuleName == "" {
			config.ModuleName = config.ProjectName
		}
		config.Port = "8080"

		fmt.Printf("\n🚀 Oke, 'hedeeh' bakal buatin projek %s...\n", config.ProjectName)
		fmt.Printf("   DB: %s | 🌐 Router: %s\n", config.Database, config.Router)

		// 5. Panggil Logic Generator (Membuat Folder & File)
		err = generator.ScaffoldProject(config)
		if err != nil {
			fmt.Printf("❌ Waduh, gagal scaffold: %v\n", err)
			return
		}

		if err := saveConfig(config); err != nil {
			fmt.Printf("⚠️  Gagal menyimpan config json: %v\n", err)
		}

		fmt.Println("✅ Selesai! Sekarang tinggal ngoding santai.")
		fmt.Printf("👉 cd %s\n", config.ProjectName)
		fmt.Println("👉 go mod tidy (untuk download driver DB)")
	},
}

func saveConfig(cfg generator.ProjectConfig) error {
	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(cfg.ProjectName, "hedeeh.config.json")
	return os.WriteFile(path, file, 0644)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
