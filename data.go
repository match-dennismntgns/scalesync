package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var data = new(Data)

type Data struct {
	LastAt       int `json:"last_at"`
	PreviousFlag int `json:"previous_flag"`
	PreviousAt   int `json:"previous_at"`
	PreviousAry  []struct {
		ID                    int64   `json:"id"`
		BUserID               int64   `json:"b_user_id"`
		TimeStamp             int     `json:"time_stamp"`
		LocalCreatedAt        string  `json:"local_created_at"`
		TimeZone              string  `json:"time_zone"`
		CreatedAt             string  `json:"created_at"`
		CreatedStamp          int     `json:"created_stamp"`
		ScaleType             int     `json:"scale_type"`
		ScaleName             string  `json:"scale_name"`
		Mac                   string  `json:"mac"`
		Gender                int     `json:"gender"`
		Height                float64 `json:"height"`
		HeightUnit            int     `json:"height_unit"`
		Birthday              string  `json:"birthday"`
		Waistline             int     `json:"waistline"`
		Hip                   int     `json:"hip"`
		CategoryType          int     `json:"category_type"`
		PersonType            int     `json:"person_type"`
		Weight                float64 `json:"weight"`
		Bodyfat               float64 `json:"bodyfat"`
		Water                 float64 `json:"water"`
		Bmr                   int     `json:"bmr"`
		WeightUnit            int     `json:"weight_unit"`
		Bodyage               int     `json:"bodyage"`
		Muscle                float64 `json:"muscle"`
		Bone                  float64 `json:"bone"`
		Subfat                float64 `json:"subfat"`
		Visfat                int     `json:"visfat"`
		Bmi                   float64 `json:"bmi"`
		Sinew                 float64 `json:"sinew"`
		Protein               float64 `json:"protein"`
		BodyShape             int     `json:"body_shape"`
		FatFreeWeight         float64 `json:"fat_free_weight"`
		Resistance            int     `json:"resistance"`
		SecResistance         int     `json:"sec_resistance"`
		InternalModel         string  `json:"internal_model"`
		ActualResistance      int     `json:"actual_resistance"`
		ActualSecResistance   int     `json:"actual_sec_resistance"`
		HeartRate             int     `json:"heart_rate"`
		CardiacIndex          float64 `json:"cardiac_index"`
		Method                int     `json:"method"`
		SportFlag             int     `json:"sport_flag"`
		LeftWeight            float64 `json:"left_weight"`
		RightWeight           float64 `json:"right_weight"`
		BodyfatLeftArm        float64 `json:"bodyfat_left_arm"`
		BodyfatLeftLeg        float64 `json:"bodyfat_left_leg"`
		BodyfatRightLeg       float64 `json:"bodyfat_right_leg"`
		BodyfatRightArm       float64 `json:"bodyfat_right_arm"`
		BodyfatTrunk          float64 `json:"bodyfat_trunk"`
		SinewLeftArm          float64 `json:"sinew_left_arm"`
		SinewLeftLeg          float64 `json:"sinew_left_leg"`
		SinewRightArm         float64 `json:"sinew_right_arm"`
		SinewRightLeg         float64 `json:"sinew_right_leg"`
		SinewTrunk            float64 `json:"sinew_trunk"`
		Resistance20LeftArm   float64 `json:"resistance20_left_arm"`
		Resistance20LeftLeg   float64 `json:"resistance20_left_leg"`
		Resistance20RightLeg  float64 `json:"resistance20_right_leg"`
		Resistance20RightArm  float64 `json:"resistance20_right_arm"`
		Resistance20Trunk     float64 `json:"resistance20_trunk"`
		Resistance100LeftArm  float64 `json:"resistance100_left_arm"`
		Resistance100LeftLeg  float64 `json:"resistance100_left_leg"`
		Resistance100RightArm float64 `json:"resistance100_right_arm"`
		Resistance100RightLeg float64 `json:"resistance100_right_leg"`
		Resistance100Trunk    float64 `json:"resistance100_trunk"`
		Remark                string  `json:"remark"`
		Score                 float64 `json:"score"`
		PregnantFlag          int     `json:"pregnant_flag"`
		Stature               float64 `json:"stature"`
	} `json:"previous_ary"`
	LastAry       []interface{} `json:"last_ary"`
	DeleteAry     []interface{} `json:"delete_ary"`
	StatusCode    string        `json:"status_code"`
	StatusMessage string        `json:"status_message"`
}

func getUserID() int64 {

	return renphoLoginResponse.ID
}

func getData() {

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	terminal_user_session_key := login()
	user_id := fmt.Sprint(getUserID())

	req, err := http.NewRequest("GET", "https://renpho.qnclouds.com/api/v2/measurements/list.json?address=&app_id=Renpho&app_revision=3.11.10&area_code=BE&cellphone_type=iPhone14%2C2&locale=en&platform=iphone&system_type=15.4.1&terminal_user_session_key="+terminal_user_session_key+"&user_id="+user_id+"&zone=Europe/Brussels", nil)

	if err != nil {
		// handle err
	}
	req.Host = "renpho.qnclouds.com"
	req.Header.Set("User-Agent", "Renpho/3.11.10 (iPhone; iOS 15.4.1; Scale/3.00)")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-BE;q=1, nl-BE;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	// read API response

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	// store response in struct
	data = new(Data)

	err = json.Unmarshal(response, &data)
	if err != nil {
		panic(err.Error())
	}

}

func getLastDate() string {
	return fmt.Sprint(data.PreviousAt)
}

func getLastWeight() string {
	return fmt.Sprint(renphoLoginResponse.Weight)
}
