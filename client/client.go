package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/term"
)

type Mail struct {
	ID   int    `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
	Read bool   `json:"read"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func handleErr(err error, reason string) {
	if err != nil {
		fmt.Println("Error:", reason, err)
	}
}

func CheckUserExists(to string) bool {
	_ = godotenv.Load("client.env")
	var server = fmt.Sprintf("http://%s:80/", os.Getenv("pinguServerIP"))
	req, err := http.NewRequest("GET", server+"user", nil)
	handleErr(err, "Error creating request")

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err, "Error making request")
	defer resp.Body.Close()

	// parse the response body and loop thought it to check if user exist
	responseBody, err := ioutil.ReadAll(resp.Body)
	handleErr(err, "Error reading response body")

	var users []User
	json.Unmarshal(responseBody, &users)

	for _, user := range users {
		if user.Name == to {
			return true
		}
	}

	return false
}

func SendMail(to string, body string) {
	_ = godotenv.Load("client.env")
	var server = fmt.Sprintf("http://%s:80/", os.Getenv("pinguServerIP"))

	if !CheckUserExists(to) {
		fmt.Println("User does not exist")
		return
	}

	var mail = Mail{
		From: os.Getenv("pinguUserName"),
		To:   to,
		Body: body,
	}

	bodyRequest, err := json.Marshal(mail) // Replace with your custom body
	handleErr(err, "Error marshalling body")

	req, err := http.NewRequest("POST", server+"mail", bytes.NewBuffer(bodyRequest))
	handleErr(err, "Error creating request")

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err, "Error making request")
	defer resp.Body.Close()
}

func Reload() []Mail {
	err := godotenv.Load("client.env")
	handleErr(err, "Error loading .env file")
	var server = fmt.Sprintf("http://%s:80/mail", os.Getenv("pinguServerIP"))

	req, err := http.NewRequest("GET", server, nil)
	handleErr(err, "Error creating request")

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err, "Error making request")
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	handleErr(err, "Error reading response body")

	var mails []Mail
	json.Unmarshal(responseBody, &mails)

	return mails
}

func Login(userName string) string {

	var user User
	user.Name = userName

	fmt.Println("Enter password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		println("Error reading password")
		return ""
	}

	user.Password = string(password)

	// Make an HTTP request to the server to login
	_ = godotenv.Load("client.env")
	var server = fmt.Sprintf("http://%s:80/", "92.222.216.38")// os.Getenv("pinguServerIP"))

	bodyRequest, err := json.Marshal(user)
	handleErr(err, "Error marshalling body")

	req, err := http.NewRequest("POST", server+"login", bytes.NewBuffer(bodyRequest))
	handleErr(err, "Error creating request")

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err, "Error making request")
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	handleErr(err, "Error reading response body")

	return string(responseBody)
}
