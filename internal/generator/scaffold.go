package generator

import (
	"bytes"
	"fmt"
	"hedeeh/internal/templates"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func templateBasePath(cfg ProjectConfig) string {
	return fmt.Sprintf(
		"go/%s/%s/scaffold",
		normalize(cfg.Database),
		normalize(cfg.Router),
	)
}

func ScaffoldProject(cfg ProjectConfig) error {
	fmt.Printf("🛠  Sedang membangun project: %s ...\n", cfg.ProjectName)

	baseTplPath := templateBasePath(cfg)

	if _, err := templates.FS.ReadDir(baseTplPath); err != nil {
		return fmt.Errorf("template tidak ditemukan: %s", baseTplPath)
	}

	templateMap := map[string]string{
		"cmd/api/main.go": baseTplPath + "/main_http.go.tpl",

		"internal/config/config.go":         baseTplPath + "/config.go.tpl",
		"internal/database/db.go":           baseTplPath + "/db.go.tpl",
		"internal/model/models.go":          baseTplPath + "/models.go.tpl",
		"internal/repository/repository.go": baseTplPath + "/repository.go.tpl",
		"internal/service/service.go":       baseTplPath + "/service.go.tpl",

		"internal/handler/handler.go": baseTplPath + "/handler.go.tpl",
		"internal/router/router.go":   baseTplPath + "/router_http.go.tpl",

		".gitignore":         baseTplPath + "/gitignore.tpl",
		".env":               baseTplPath + "/env.tpl",
		"Dockerfile":         baseTplPath + "/Dockerfile.tpl",
		"docker-compose.yml": baseTplPath + "/docker-compose.yml.tpl",
	}

	for targetPath, tplPath := range templateMap {
		fullTargetPath := fmt.Sprintf("%s/%s", cfg.ProjectName, targetPath)

		if err := generateFromTemplate(tplPath, fullTargetPath, cfg); err != nil {
			return fmt.Errorf("gagal generate %s: %w", targetPath, err)
		}

		fmt.Printf("   ✅ Created: %s\n", targetPath)
	}

	return initGoModule(cfg)
}

func generateFromTemplate(tplPath, targetPath string, data ProjectConfig) error {
	tplBytes, err := templates.FS.ReadFile(tplPath)
	if err != nil {
		return fmt.Errorf("template tidak ditemukan: %s", tplPath)
	}

	tpl, err := template.New(filepath.Base(tplPath)).Parse(string(tplBytes))
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		return err
	}

	return writeFile(targetPath, buffer.Bytes())
}

func initGoModule(cfg ProjectConfig) error {
	fmt.Println("📦 Menginisialisasi Go Module...")

	cmdInit := exec.Command("go", "mod", "init", cfg.ModuleName)
	cmdInit.Dir = cfg.ProjectName
	if out, err := cmdInit.CombinedOutput(); err != nil {
		return fmt.Errorf("go mod init error: %s", string(out))
	}

	cmdTidy := exec.Command("go", "mod", "tidy")
	cmdTidy.Dir = cfg.ProjectName
	if out, err := cmdTidy.CombinedOutput(); err != nil {
		fmt.Printf("⚠️  Warning go mod tidy: %s\n", string(out))
	}

	return nil
}
