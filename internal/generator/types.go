package generator

type ProjectConfig struct {
	ProjectName string `json:"project_name"`
	ModuleName  string `json:"module_name"`
	Language    string `json:"language"`
	Database    string `json:"database"`
	Router      string `json:"router"`
	Port        string `json:"port"`
}
