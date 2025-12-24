package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"hedeeh/internal/ai"
	"hedeeh/internal/ai/prompts"
	"hedeeh/internal/generator"

	"github.com/spf13/cobra"
)

var fieldsFlag string

var addCmd = &cobra.Command{
	Use:   "add [feature-name]",
	Short: "Menambahkan fitur baru (Handler, Service, Repo, Router) pakai AI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		config, err := loadConfig()
		if err != nil {
			fmt.Println("❌ Gagal membaca 'hedeeh.config.json'.")
			fmt.Println("👉 Pastikan kamu menjalankan perintah ini di dalam folder root project.")
			return
		}

		featureName := args[0]

		fmt.Printf(" Mode: %s | Router: %s | DB: %s\n", config.Language, config.Router, config.Database)
		fmt.Printf("🚀 Meminta AI membuat fitur '%s'...\n", featureName)

		generateLayer("Model", featureName, fieldsFlag, "internal/model", config)

		generateLayer("Repository", featureName, fieldsFlag, "internal/repository", config)

		generateLayer("Service", featureName, fieldsFlag, "internal/service", config)

		generateLayer("Handler", featureName, fieldsFlag, "internal/handler", config)

		generateLayer("Router", featureName, fieldsFlag, "internal/router", config)

		fmt.Println("\n✅ Selesai! Jangan lupa cek dan rapihin import-nya.")
	},
}

func loadConfig() (generator.ProjectConfig, error) {
	var cfg generator.ProjectConfig

	fileData, err := os.ReadFile("hedeeh.config.json")
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(fileData, &cfg)
	return cfg, err
}

func generateLayer(layer, name, fields, folder string, cfg generator.ProjectConfig) {
	fmt.Printf("   -> Generating %s...", layer)

	var prompt string

	if cfg.Language == "Go" {

		prompt = prompts.GenerateGoPrompt(layer, name, fields, cfg.Router)
	} else {

		fmt.Printf(" (Bahasa %s belum didukung AI) ", cfg.Language)
		return
	}

	if prompt == "" {
		fmt.Printf(" Error: Prompt kosong untuk layer '%s'. Cek logic prompt.\n", layer)
		return
	}

	code, err := ai.GenerateCode(prompt)
	if err != nil {
		fmt.Printf(" Error AI: %v\n", err)
		return
	}

	filename := fmt.Sprintf("%s_%s.go", strings.ToLower(name), strings.ToLower(layer))

	if layer == "Model" {
		filename = fmt.Sprintf("%s.go", strings.ToLower(name))
	}

	path := filepath.Join(folder, filename)

	os.MkdirAll(folder, 0755)

	err = os.WriteFile(path, []byte(code), 0644)
	if err != nil {
		fmt.Printf(" Gagal nulis file: %v\n", err)
	} else {
		fmt.Printf(" ✅ Created: %s\n", path)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&fieldsFlag, "fields", "f", "", "Daftar field entitas (cth: 'name:string, price:int')")
}
