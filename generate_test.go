package main

import (
	"os"
	"path"
	"regexp"
	"testing"
)

func TestGenerate(t *testing.T) {
	tmpDir := t.TempDir()

	testData := `{
  "pluginsdk": {
    "URL": "https://github.com/arcalot/arcaflow-plugin-sdk-go",
    "MainBranch": "main"
  }
}`

	if err := os.WriteFile(path.Join(tmpDir, "packages.json"), []byte(testData), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("GITHUB_REPOSITORY_OWNER", "arcalot"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("DOMAIN_NAME", "go.flow.arcalot.io"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("GITHUB_REPOSITORY", "go.flow.arcalot.io"); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir to %s", tmpDir)
	}
	main()

	cnameContents, err := os.ReadFile(path.Join(tmpDir, "gh-pages", "CNAME"))
	if err != nil {
		t.Fatal(err)
	}
	if string(cnameContents) != "go.flow.arcalot.io" {
		t.Fatalf("Invalid contents for CNAME file: %s", cnameContents)
	}

	indexFile := path.Join(tmpDir, "gh-pages", "pluginsdk", "index.html")
	indexContents, err := os.ReadFile(indexFile)
	if err != nil {
		t.Fatal(err)
	}
	re := regexp.MustCompile(`<meta\s*name="go-import"\s*content="go.flow.arcalot.io/pluginsdk\s*git\s*https://github.com/arcalot/arcaflow-plugin-sdk-go" />`)
	if !re.MatchString(string(indexContents)) {
		t.Fatalf("No go-import tag found in %s", indexFile)
	}
}
