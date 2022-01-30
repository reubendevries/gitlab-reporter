package reports

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//func getPages(w http.ResponseWriter, r *http.Request) int {
//	p := r.Header.Get(s)
//}

func ListUsers(u string, t string) {
	var result string
	var bearer = "Bearer " + t
	lu := "/api/v4/users"
	req, err := http.NewRequest("GET", u+lu, nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	h := req.Header

	h.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()
	totalPages := (resp.Header.Get("x-total-pages"))
	pages, err := strconv.Atoi(totalPages)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i := 1; i >= pages; i++ {
		q := url.Values{}
		q.Add("pages=", strconv.Itoa(i))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		result = string(body)
		fmt.Println(result)
	}
	//return result, nil
}

func ListActiveUsers() {

}

func ListBlockedUsers() {

}

func ListExternalUsers() {

}

func ListUsersUsing2FA() {

}

func ListGroups() {

}

func ListGroupProjects() {

}
