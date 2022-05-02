package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var renphoLoginResponse = new(Renpho)

type Renpho struct {
	StatusCode             string `json:"status_code"`
	StatusMessage          string `json:"status_message"`
	TerminalUserSessionKey string `json:"terminal_user_session_key"`
	DeviceBindsAry         []struct {
		ID                int64  `json:"id"`
		Mac               string `json:"mac"`
		ScaleName         string `json:"scale_name"`
		Demo              string `json:"demo"`
		HwBleVersion      int    `json:"hw_ble_version"`
		DeviceType        int    `json:"device_type"`
		HwSoftwareVersion int    `json:"hw_software_version"`
		CreatedAt         string `json:"created_at"`
		UUID              string `json:"uuid"`
		BUserID           int64  `json:"b_user_id"`
		InternalModel     string `json:"internal_model"`
		WifiName          string `json:"wifi_name"`
		ProductCategory   int    `json:"product_category"`
	} `json:"device_binds_ary"`
	ID                    int64   `json:"id"`
	Email                 string  `json:"email"`
	AccountName           string  `json:"account_name"`
	Gender                int     `json:"gender"`
	Height                float64 `json:"height"`
	HeightUnit            int     `json:"height_unit"`
	Waistline             int     `json:"waistline"`
	Hip                   int     `json:"hip"`
	PersonType            int     `json:"person_type"`
	CategoryType          int     `json:"category_type"`
	WeightUnit            int     `json:"weight_unit"`
	CurrentGoalWeight     float64 `json:"current_goal_weight"`
	WeightGoalUnit        int     `json:"weight_goal_unit"`
	WeightGoal            float64 `json:"weight_goal"`
	Locale                string  `json:"locale"`
	Birthday              string  `json:"birthday"`
	WeightGoalDate        string  `json:"weight_goal_date"`
	AvatarURL             string  `json:"avatar_url"`
	Weight                float64 `json:"weight"`
	FacebookAccount       string  `json:"facebook_account"`
	TwitterAccount        string  `json:"twitter_account"`
	LineAccount           string  `json:"line_account"`
	SportGoal             int     `json:"sport_goal"`
	SleepGoal             int     `json:"sleep_goal"`
	BodyfatGoal           float64 `json:"bodyfat_goal"`
	InitialWeight         float64 `json:"initial_weight"`
	InitialBodyfat        float64 `json:"initial_bodyfat"`
	AreaCode              string  `json:"area_code"`
	Method                int     `json:"method"`
	UserCode              string  `json:"user_code"`
	AgreeFlag             int     `json:"agree_flag"`
	ReachGoalWeightFlag   int     `json:"reach_goal_weight_flag"`
	ReachGoalBodyfatFlag  int     `json:"reach_goal_bodyfat_flag"`
	SetGoalAt             int     `json:"set_goal_at"`
	SellFlag              int     `json:"sell_flag"`
	AllowNotificationFlag int     `json:"allow_notification_flag"`
	Phone                 string  `json:"phone"`
	RegionCode            string  `json:"region_code"`
	DumpFlag              int     `json:"dump_flag"`
	WeighingMode          int     `json:"weighing_mode"`
	PasswordPresentFlag   int     `json:"password_present_flag"`
	Stature               float64 `json:"stature"`
	Custom                string  `json:"custom"`
	IndexExtension        int     `json:"index_extension"`
	PersonBodyShape       int     `json:"person_body_shape"`
	PersonGoal            int     `json:"person_goal"`
}

func login() string {

	// login using creds in env file
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	email := os.Getenv("email")
	hashedpassword := os.Getenv("hashedpassword")

	params := url.Values{}
	params.Add("address", ``)
	params.Add("app_id", `Renpho`)
	params.Add("app_revision", `3.11.10`)
	params.Add("area_code", `BE`)
	params.Add("cellphone_type", `iPhone14,2`)
	params.Add("email", email)
	params.Add("locale", `en`)
	params.Add("password", hashedpassword)
	params.Add("platform", `iphone`)
	params.Add("secure_flag", `1`)
	params.Add("system_type", `15.4.1`)
	params.Add("zone", `Europe/Brussels`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://renpho.qnclouds.com/api/v3/users/sign_in.json", body)
	if err != nil {
		// handle err
	}
	req.Host = "renpho.qnclouds.com"
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("User-Agent", "Renpho/3.11.10 (iPhone; iOS 15.4.1; Scale/3.00)")
	req.Header.Set("Accept-Language", "en-BE;q=1, nl-BE;q=0.9")
	req.Header.Set("Connection", "close")

	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// read API response

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	// store response in struct
	renphoLoginResponse = new(Renpho)

	err = json.Unmarshal(response, &renphoLoginResponse)
	if err != nil {
		panic(err.Error())
	}

	// return session key to function
	return renphoLoginResponse.TerminalUserSessionKey

}
