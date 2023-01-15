package tmplfuncs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func PrefixLines(prefix, text string) string {
	return prefix + strings.Join(strings.Split(text, "\n"), "\n"+prefix)
}

func CodeFile(format, file string) (string, error) {
	// paths are relative to the rendering process work dir, which
	// may be undesirable, probably need to think about it
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(wd, file)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("unable to read content from %q: %w", file, err)
	}

	sContent := strings.TrimSpace(string(content))
	if sContent == "" {
		return "", fmt.Errorf("no file content in %q", file)
	}

	md := &strings.Builder{}
	_, err = md.WriteString("```" + format + "\n")
	if err != nil {
		return "", err
	}

	_, err = md.WriteString(sContent)
	if err != nil {
		return "", err
	}

	_, err = md.WriteString("\n```")
	if err != nil {
		return "", err
	}

	return md.String(), nil
}

func SubCategory(name string) string {
	switch name {
	case "azuredevops_environment",
		"azuredevops_environment_kubernetes",
		"azuredevops_environment_permissions":
		return "Pipelines"
	case "azuredevops_process",
		"azuredevops_project":
		return "Projects"
	case "azuredevops_serviceendpoint_azurerm",
		"azuredevops_serviceendpoint_bitbucket",
		"azuredevops_serviceendpoint_github",
		"azuredevops_serviceendpoint_kubernetes",
		"azuredevops_serviceendpoint_share",
		"azuredevops_serviceendpoint_vsappcenter":
		return "Service Endpoints"
	case "azuredevops_group",
		"azuredevops_group_membership",
		"azuredevops_groups",
		"azuredevops_team",
		"azuredevops_teams",
		"azuredevops_user",
		"azuredevops_users":
		return "Users & Groups"
	case "azuredevops_area",
		"azuredevops_iteration":
		return "Work Items"
	default:
		return ""
	}
}
