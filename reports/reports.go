package reports

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//func getPages(w http.ResponseWriter, r *http.Request) int {
//	p := r.Header.Get(s)
//}

func ListUsers(m map[string]string) (string, error) {
	var result string
	var bearer = "Bearer " + m["Token"]
	lu := "/api/v4/users"
	req, err := http.NewRequest("GET", m["Url"]+lu, nil)
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
	totalPages := (response.Header.Get("x-total-pages"))
	pages, err := strconv.Atoi(totalPages)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i := 1; i >= pages; i++ {
		q := url.Values{}
		q.Add("pages=", strconv.Itoa(i))
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		result = string(body)
	}
	return result, nil
}

func ListActiveUsers(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func ListBlockedUsers(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func ListExternalUsers(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func ListUsersUsing2FA(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func ListGroups(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func ListGroupProjects(m map[string]string) (string, error) {
	var result string
	return result, nil
}
