package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const ServerURL = "https://blops.me"

var Token string

type ApiTest struct {
	Name           string
	Method         string
	Body           string
	URL            string
	ExpectedStatus int
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	RealName    string `json:"real_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

func RunTest(test ApiTest) (int, string) {
	return SendRequest(test.Method, test.URL, test.Body)
}

var (
	FailedTests = 0
	PassedTests = 0
)

func RunTests(tests []ApiTest) {
	for _, test := range tests {
		status, body := RunTest(test)
		if status != test.ExpectedStatus {
			log.Println("Failed:", test.Name, "status:", status, "body:", body)
			FailedTests++
		} else {
			log.Println("Passed:", test.Name, "status:", status, "body:", body)
			PassedTests++
		}
	}
}

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func SendRequest(method, url, body string) (int, string) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	if Token != "" {
		req.Header.Set("Cookie", "token="+Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyStr := ""
	if resp.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyStr = string(bodyBytes)
	}

	return resp.StatusCode, bodyStr
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	CreateTokenForDebug()

	TestUserRouter()

	log.Println("Passed tests:", PassedTests)
	log.Println("Failed tests:", FailedTests)

	test := ApiTest{
		Name:           "Drop all data",
		Method:         "GET",
		URL:            ServerURL + "/api/v1/debug/dropalldata",
		ExpectedStatus: 200,
	}
	status, body := RunTest(test)
	if status != 200 {
		panic("Failed to drop all data: " + body)
	}

	log.Println("All tests passed")
}

func CreateTokenForDebug() {
	newUser := User{
		Username:    RandomString(10),
		RealName:    "Test User",
		Email:       RandomString(10) + "@test.com",
		PhoneNumber: "01012345678",
	}
	bodyData, _ := json.Marshal(newUser)

	test := ApiTest{
		Name:   "Create token for debug",
		Method: "POST",
		Body:   string(bodyData),
		URL:    ServerURL + "/api/v1/debug/token",
	}
	status, body := RunTest(test)
	if status != 200 {
		panic("Failed to create token for debug: " + body)
	}

	var response struct {
		Token string `json:"token"`
	}
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}

	Token = response.Token

	log.Println("Token:", Token)
}

func TestUserRouter() {
	newUser := User{
		Username:    RandomString(10),
		RealName:    "Test User",
		Email:       RandomString(10) + "@test.com",
		PhoneNumber: "01012345678",
	}

	updatedUser := User{
		ID:          1,
		Username:    RandomString(10),
		RealName:    "Test User",
		Email:       RandomString(10) + "@test.com",
		PhoneNumber: "01012345678",
	}

	bodyData, _ := json.Marshal(newUser)
	bodyData2, _ := json.Marshal(updatedUser)

	tests := []ApiTest{
		{
			Name:           "Get logged in user",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/user",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get user by ID",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/user/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Create user",
			Method:         "POST",
			Body:           string(bodyData),
			URL:            ServerURL + "/api/v1/user",
			ExpectedStatus: 200,
		},
		{
			Name:           "Update user",
			Method:         "PATCH",
			Body:           string(bodyData2),
			URL:            ServerURL + "/api/v1/user",
			ExpectedStatus: 200,
		},
	}

	RunTests(tests)
}
