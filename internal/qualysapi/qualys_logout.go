package qualysapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) QualysLogout(baseUrl string) error {
	qUrl := baseUrl + "/api/2.0/fo/session/"
	sessionCookie := fmt.Sprintf("QualysSession=%s", c.Session)
	formData := url.Values{
		"action": {"logout"},
	}

	reqBody := strings.NewReader(formData.Encode())

	req, err := http.NewRequest("POST", qUrl, reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Add("X-Requested-With", "QualysToolBox")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", sessionCookie)

	fmt.Println("Logging out...")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Logout failed with status code:", resp.StatusCode)
		return errors.New("Failed Logout!")
	}

	return nil
}
