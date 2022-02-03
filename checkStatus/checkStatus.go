package checkStatus

import (
	"errors"
	"log"
	"net/http"
)

func CheckGitlabStatus(u string) error {
	url := u + "/users/sign_in"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
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

func CheckVaultStatus(a string) error {
	endpoint := a + "/ui/vault"
	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
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

func ValidReportCheck(s string) error {
	reports := []string{"list_users", "list_active_users", "list_blocked_users", "list_external_users", "list_users_using_2FA", "list_groups", "list_group_projects"}
	for _, v := range reports {
		if v == s {
			return nil
		}
		continue
	}
	return errors.New("report name passed is not on list of valid reports")
}
