package parseArgs

import (
	"errors"
	"fmt"
	"gitlab-reporter/helpMenu"
	"gitlab-reporter/reports"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

func getArgs(s []string) ([]string, error) {
	var args []string
	for i := 0; i < len(os.Args); i++ {
		if i != 0 {
			args = append(args, os.Args[i])
		}
	}
	return args, nil
}

func checkGitlabStatus(u, t string) error {
	apiEndpoint := u + "/api/v4/version"
	var bearer = "Bearer " + t
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}
	h := req.Header
	h.Add("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case 400:
		return errors.New(response.Status)
	case 401:
		return errors.New(response.Status)
	case 403:
		return errors.New(response.Status)
	case 404:
		return errors.New(response.Status)
	case 405:
		return errors.New(response.Status)
	case 409:
		return errors.New(response.Status)
	case 412:
		return errors.New(response.Status)
	case 422:
		return errors.New(response.Status)
	case 429:
		return errors.New(response.Status)
	case 500:
		return errors.New(response.Status)
	}
	return nil
}

func checkVaultStatus(a, t string) error {
	apiEndpoint := a + "/v1/secret?help=1"
	fmt.Println(apiEndpoint)
	var bearer = "Bearer " + t
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	h := req.Header
	h.Add("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case 400:
		return errors.New(response.Status)
	case 403:
		return errors.New(response.Status)
	case 404:
		return errors.New(response.Status)
	case 405:
		return errors.New(response.Status)
	case 412:
		return errors.New(response.Status)
	case 429:
		return errors.New(response.Status)
	case 473:
		return errors.New(response.Status)
	case 500:
		return errors.New(response.Status)
	case 502:
		return errors.New(response.Status)
	case 503:
		return errors.New(response.Status)
	}
	return nil
}

func getVaultSecret(a, t, n, p, k string) (string, error) {
	config := &api.Config{
		Address: a,
	}
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	client.SetToken(t)
	secret, err := client.Logical().Read(p)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	m, ok := secret.Data["data"].(map[string]string)
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
	}
	return m[k], nil
}

func validReportCheck(s string) (string, error) {
	reports := []string{"list_users", "list_active_users", "list_blocked_users", "list_external_users", "list_users_using_2FA", "list_groups", "list_group_projects"}
	for _, v := range reports {
		if v == s {
			return s, nil
		}
		continue
	}
	return s, errors.New("report name passed is not on list of valid reports")
}

// TODO add a function that will allow the user to change the directory where the report gets saved.
//func setExportDirectory(s string) (string, error) {
//	return s, nil
//}

// TODO add a function that will allow the user to change the output of the report.
//func setOutputFormat(s string) (string, error) {
//	return s, nil
//}

func parseToken(m map[string]string) {
	if v, found := m["Address"]; found {
		err := checkVaultStatus(v, m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		token, err := getVaultSecret(v, m["Token"], m["Namespace"], m["Path"], m["Key"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		m["Token"] = token
		err = checkGitlabStatus(m["Url"], m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
	} else {
		err := checkGitlabStatus(m["Url"], m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		executeReport(m)
	}
}

func executeReport(m map[string]string) (string, error) {
	var result string
	switch m["Report"] {
	case "list_users":
		result, err := reports.ListUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_active_users":
		result, err := reports.ListActiveUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_blocked_users":
		result, err := reports.ListBlockedUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_external_users":
		result, err := reports.ListExternalUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_users_using_2FA":
		result, err := reports.ListUsersUsing2FA(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_groups":
		result, err := reports.ListGroups(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_group_projects":
		result, err := reports.ListGroupProjects(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	default:
		return result, errors.New("report not found, exiting application")
	}
}

func ParseArgs() {
	data := make(map[string]string)
	args, err := getArgs(os.Args)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--gitlab_url", "-gitlab_url", "gitlab_url", "-gu":
			if strings.Contains(args[i+1], "help") {
				helpMenu.GitlabUrlHelp()
			} else {
				data["Url"] = args[i+1]
			}
		case "--vault_address", "-vault_address", "vault_address", "-va":
			if strings.Contains(args[i+1], "help") {
				helpMenu.VaultAddressHelp()
			} else {
				data["Address"] = args[i+1]
			}
		case "--namespace", "-namespace", "namespace", "-n":
			if strings.Contains(args[i+1], "help") {
				helpMenu.NamespaceHelp()
			} else {
				data["Namespace"] = args[i+1]
			}
		case "--path", "-path", "path", "-p":
			if strings.Contains(args[i+1], "help") {
				helpMenu.PathHelp()
			} else {
				data["Path"] = args[i+1]
			}
		case "--secret_key", "-secret_key", "secret_key", "-sk":
			if strings.Contains(args[i+1], "help") {
				helpMenu.SecretKeyHelp()
			} else {
				data["Key"] = args[i+1]
			}
		case "--api_token", "-api_token", "api_token", "-at":
			if strings.Contains(args[i+1], "help") {
				helpMenu.ApiTokenHelp()
			} else {
				data["Token"] = args[i+1]
			}
		case "--report", "-report", "report", "-r":
			response, err := validReportCheck(args[i])
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			if strings.Contains(args[i+1], "help") {
				helpMenu.ReportHelp()
			} else {
				data["Report"] = response
			}
		case "--export_dir", "-export_dir", "export_dir", "-ed":
			if strings.Contains(args[i+1], "help") {
				helpMenu.ExportDirectoryHelp()
			} else {
				data["Directory"] = args[i+1]
			}
		case "--ouput_format", "-output_format", "output_format", "-of":
			if strings.Contains(args[i+1], "help") {
				helpMenu.OutputFormatHelp()
			} else {
				data["Format"] = args[i+1]
			}
		case "--version", "-version", "version", "-v":
			helpMenu.DisplayVersion()
		case "--help", "-help", "help", "-h":
			helpMenu.DefaultHelp()
		}
	}
	parseToken(data)
}
