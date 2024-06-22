package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 🔑 Get the auth token from the environment variable
var mobitagAPIKey = os.Getenv("OPTNC_MOBITAGNC_API_KEY")

// sendSMS sends an SMS to the specified receiver mobile number
// receiverMobile: the mobile number of the receiver, like 654321
// message: the message to send
func sendSMS(receiverMobile string, message string) {
	apiURL := "https://api.opt.nc/mobitag/sendSms"

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		fmt.Printf("An error occurred while creating the request: %v\n", err)
		return
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	// Authenticate the request
	req.Header.Set("x-apikey", mobitagAPIKey)

	// set request paylod with receiver mobile number and message
	reqBody := fmt.Sprintf(`{"to":"%s","message":"%s"}`, receiverMobile, message)
	req.Body = ioutil.NopCloser(strings.NewReader(reqBody))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❗An error occurred while sending the request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("ℹ️  Accusé reception: %v\n", resp.Status)
	fmt.Printf("📜  Code retour: %v\n", resp.StatusCode)
}

func main() {
	// Parse command line arguments
	to := ""
	message := ""
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		// Parse recipient mobile number
		case "-t", "-to":
			i++
			if i < len(os.Args) {
				to = os.Args[i]
			}
		// Parse message
		case "-m", "-message":
			i++
			if i < len(os.Args) {
				message = os.Args[i]
			}
		case "-h", "--help":
			fmt.Println("mobitag -t <recipient_mobile> -m <message> [-d] [-v] [-h]")
			return
		case "-d", "--dry-run":
			// Perform dry run test
			if mobitagAPIKey != "" {
				fmt.Println("✅  Dry run test passed")
			} else {
				fmt.Println("❌  Dry run test failed: OPTNC_MOBITAGNC_API_KEY environment variable is not set")
			}
			return
		case "-v":
			fmt.Println("v0.0")
			return
		}
	}

	// Check if required arguments are provided
	if to == "" || message == "" {
		fmt.Println("❌ Missing required arguments : run <mobitag -h> for help")
		return
	}

	// Send SMS
	sendSMS(to, message)
}
