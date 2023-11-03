package qualysapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) QualysLogin(baseUrl, user, pass string) (string, error) {
	qUrl := baseUrl + "/api/2.0/fo/session/"
	formData := url.Values{
		"action":   {"login"},
		"username": {user},
		"password": {pass},
	}

	reqBody := strings.NewReader(formData.Encode())

	req, err := http.NewRequest("POST", qUrl, reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Add("X-Requested-With", "QualysToolBox")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed with status code:", resp.StatusCode)
		return "", errors.New("Failed Login!")
	}

	cookies := resp.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "QualysSession" {
			c.Session = cookie.Value
		}
	}

	return "Login Successful!", nil
}
