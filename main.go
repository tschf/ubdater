package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/urfave/cli/v2"
)

// ExtensionMoh is a nested object in the Ubity extension API dataset. It doesn't
// get returned in the GET request, but when PUTting new data, it has it there.
// This is in the JSON attribute "moh".
type ExtensionMoh struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Extension represents the json structure returned by the Ubity web app from
// the URL: https://studio.ubity.com/extensions_api/n where the /n represents
// the extension ID
type Extension struct {
	DefaultLang          string       `json:"default_language"`
	VMFullName           string       `json:"vm_fullname"`
	FmFmMusicOnHold      bool         `json:"fmfm_musiconhold"`
	Assigned             bool         `json:"assigned"`
	FmFmRingDuration     int          `json:"fmfm_ring_duration"`
	FindmeFollowme       bool         `json:"findmefollowme"`
	VoicemailDelete      bool         `json:"voicemail_delete"`
	LongDistanceCode     string       `json:"longdistance_code"`
	VoicemailMaxMessages int          `json:"voicemail_max_messages"`
	CallerIDNum          string       `json:"callerid_num"`
	HotdeskingTarget     bool         `json:"hotdesking_target"`
	RemotePickup         bool         `json:"remote_pickup"`
	RecordCalls          bool         `json:"record_calls"`
	Username             string       `json:"username"`
	FmFmRequireID        bool         `json:"fmfm_requireid"`
	HideCallerID         bool         `json:"hidecallerid"`
	RingDuration         int          `json:"ring_duration"`
	CallerIDNumAdmin     string       `json:"callerid_num_admin"`
	ForwardTo            string       `json:"forward_to"`
	UcMobileEnabled      bool         `json:"uc_mobile_enabled"`
	VideoEnabled         bool         `json:"video_enabled"`
	MusicClass           string       `json:"musicclass"`
	ListOnDir            bool         `json:"list_on_dir"`
	BusyMsgBehaviour     string       `json:"busy_message_behaviour"`
	VoicemailPassword    string       `json:"voicemail_password"`
	Record               string       `json:"record"`
	VoicemailEnabled     bool         `json:"voicemail_enabled"`
	RecordCallsIncoming  bool         `json:"record_calls_incoming"`
	FmFmMaxRetries       int          `json:"fmfm_maxretries"`
	CallerIDName         string       `json:"callerid_name"`
	FullName             string       `json:"fullname"`
	UcDesktopEnabled     bool         `json:"uc_desktop_enabled"`
	FmFmSkipInto         bool         `json:"fmfm_skipintro"`
	HasVoicmail          string       `json:"has_voicemail"`
	Moh                  ExtensionMoh `json:"moh"`
	NumLines             int          `json:"nb_lines"`
}

func main() {

	app := &cli.App{
		Name:  "ubdater",
		Usage: "Ubity update tool",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "extension",
				Aliases:  []string{"e"},
				Usage:    "the extension to modify",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "forward-to",
				Aliases:  []string{"f"},
				Usage:    "the forward to number",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			updateExtension(c.Int("extension"), c.String("forward-to"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func updateExtension(extension int, forwardNum string) {
	fmt.Println("Updating extension", extension)

	var extensionURL string = fmt.Sprintf("https://studio.ubity.com/extensions_api/%v", extension)
	fmt.Printf("API URL: %s\n", extensionURL)

	cookieJar, _ := cookiejar.New(nil)

	requestBody := url.Values{}
	requestBody.Set("login", os.Getenv("UBITY_LOGIN"))
	requestBody.Set("password", os.Getenv("UBITY_PASSWORD"))

	client := &http.Client{
		Jar: cookieJar,
	}

	loginReq, _ := http.NewRequest(http.MethodPost, "https://studio.ubity.com/login_handler", bytes.NewBufferString(requestBody.Encode()))
	loginReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	loginResp, _ := client.Do(loginReq)

	defer loginResp.Body.Close()

	if loginResp.StatusCode == http.StatusOK {
		fmt.Println("Login OK")
		validatedResp, _ := client.Get(extensionURL)
		defer validatedResp.Body.Close()

		bodyBytes, _ := ioutil.ReadAll(validatedResp.Body)

		extAttributes := Extension{}
		json.Unmarshal(bodyBytes, &extAttributes)

		fmt.Printf("Number currently set to %s\n", extAttributes.ForwardTo)

		extAttributes.ForwardTo = forwardNum
		extAttributes.Moh.Name = "default"
		extAttributes.Moh.Description = "default"
		extAttributes.NumLines = 10

		modifiedPayload, _ := json.Marshal(extAttributes)

		updateReq, _ := http.NewRequest(http.MethodPut, extensionURL, bytes.NewBuffer(modifiedPayload))

		updateResp, _ := client.Do(updateReq)
		defer updateResp.Body.Close()

		bodyBytes, _ = ioutil.ReadAll(updateResp.Body)
		fmt.Printf("Server says %s\n", string(bodyBytes))
	}
}
