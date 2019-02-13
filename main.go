package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

// set at compile-time using -ldflags "-X main.version=$VERSION"
var version string

type verifyPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type verifyPasswordResponse struct {
	Kind         string  `json:"king"`
	LocalID      string  `json:"localId"`
	Email        string  `json:"email"`
	DisplayName  string  `json:"displayName"`
	IDToken      IDToken `json:"idToken"`
	Registered   bool    `json:"registered"`
	RefreshToken string  `json:"refreshToken"`
	expiresIn    string  `json:""expiresIn`
}

type badRequestResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		}
	} `json:"error"`
}

type JWTPayload struct {
	CUUID string `json:"cuuid,omitempty"`
	Role  string `json:"role,omitempty"`
}

func (p JWTPayload) String() string {
	return fmt.Sprintf("CUUID: %s Role: %s", p.CUUID, p.Role)
}

type IDToken string

func (t IDToken) Claims() (string, error) {
	parts := strings.Split(string(t), ".")
	decoded, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	var payload JWTPayload
	err = json.Unmarshal(decoded, &payload)
	if err != nil {
		fmt.Errorf("%v", err)
		return "", err
	}

	return payload.String(), nil
}

func verifyPassword(key, email, password string) (*verifyPasswordResponse, error) {
	vpReq := verifyPasswordRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}
	j, err := json.Marshal(vpReq)
	if err != nil {
		return nil, err
	}

	// build the URL including Query params
	v := url.Values{}
	v.Set("key", key)
	u := url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "identitytoolkit/v3/relyingparty/verifyPassword",
		ForceQuery: false,
		RawQuery:   v.Encode(),
	}

	// build and execute the request
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(j))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode >= 400 {
		var badReqRes badRequestResponse
		err = json.Unmarshal(body, &badReqRes)
		if err != nil {
			return nil, err
		}

		fmt.Fprintf(os.Stderr, "%d %s\n", badReqRes.Error.Code, badReqRes.Error.Message)
		os.Exit(1)
	}

	var vpRes verifyPasswordResponse
	err = json.Unmarshal(body, &vpRes)
	if err != nil {
		return nil, err
	}
	return &vpRes, nil
}

func main() {
	var key, email, password string
	showVersion := flag.Bool("v", false, "Shows the version of the command line tool.")
	flag.StringVar(&key, "w", "", "Web API Key (Project Overview -> Users and Permissions -> General).")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if key == "" {
		fmt.Printf("You must specify a Web API Key using the -w flag.\nSee Firebase (Project Overview -> Users and Permissions -> General).\n")
		os.Exit(1)
	}
	fmt.Printf("Web API Key: %s\n", key)

	promptE := &survey.Input{
		Message: "Email:",
	}
	survey.AskOne(promptE, &email, nil)

	promptP := &survey.Password{
		Message: "Password:",
	}
	survey.AskOne(promptP, &password, nil)

	vpr, err := verifyPassword(key, email, password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error verifying credentials: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Display Name: %s\n", vpr.DisplayName)
	fmt.Printf("Email: %s\n", vpr.Email)
	fmt.Printf("\nIDToken:\n%s\n", vpr.IDToken)

	claims, err := vpr.IDToken.Claims()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error extracting claims: %v", err)
		os.Exit(1)
	}

	fmt.Printf("\nClaims:\n%s\n", claims)
	os.Exit(0)
}
