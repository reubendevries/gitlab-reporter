package parseArgs

import (
	"gitlab-reporter/checkStatus"
	"gitlab-reporter/getKey"
	"gitlab-reporter/helpMenu"
	"gitlab-reporter/reports"
	"log"
	"os"
	"strings"
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
		err := checkStatus.CheckVaultStatus(v, m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		token, err := getKey.GetVaultSecret(v, m["Token"], m["Namespace"], m["Path"], m["Key"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		m["Token"] = token
		err = checkStatus.CheckGitlabStatus(m["Url"], m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		reports.ExecuteReport(m)
	} else {
		err := checkStatus.CheckGitlabStatus(m["Url"], m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		reports.ExecuteReport(m)
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
			response, err := reports.ValidReportCheck(args[i])
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
