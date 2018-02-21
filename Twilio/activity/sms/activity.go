package sms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivAccountSID = "accountSID"
	ivauthToken  = "authToken"
	ivFromNumber = "FromNumber"
	ivToNumber   = "ToNumber"
	ivSMStext    = "SMStext"
	ovsend       = "send"
)

var activityLog = logger.GetLogger("twilio-activity-sms")

type smsActivity struct {
	metadata *activity.Metadata
}

//NewActivity TCI Wi Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &smsActivity{metadata: metadata}
}

func (a *smsActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *smsActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing Twilio SMS Sender activity")
	//Read Inputs
	if context.GetInput(ivAccountSID) == nil {
		// AccountSID string is not configured
		// return error to the engine
		return false, activity.NewError("Twilio AccountSID string is not configured", "Twilio-SMS-4001", nil)
	}
	AccountSID := context.GetInput(ivAccountSID).(string)

	if context.GetInput(ivauthToken) == nil {
		// authToken string is not configured
		// return error to the engine
		return false, activity.NewError("Twilio authToken string is not configured", "Twilio-SMS-4002", nil)
	}
	authToken := context.GetInput(ivauthToken).(string)

	if context.GetInput(ivFromNumber) == nil {
		// FromNumber string is not configured
		// return error to the engine
		return false, activity.NewError("Twilio FromNumber string is not configured", "Twilio-SMS-4003", nil)
	}
	FromNumber := context.GetInput(ivFromNumber).(string)

	if context.GetInput(ivToNumber) == nil {
		// ToNumber string is not configured
		// return error to the engine
		return false, activity.NewError("Twilio ToNumber string is not configured", "Twilio-SMS-4004", nil)
	}
	ToNumber := context.GetInput(ivToNumber).(string)
	activityLog.Info("ToNumber: " + ToNumber)

	if context.GetInput(ivSMStext) == nil {
		// SMStext string is not configured
		// return error to the engine
		return false, activity.NewError("Twilio SMStext string is not configured", "Twilio-SMS-4005", nil)
	}
	SMStext := context.GetInput(ivSMStext).(string)
	activityLog.Info("SMStext: " + SMStext)

	// execute validation - Start
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + AccountSID + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", ToNumber)
	v.Set("From", FromNumber)
	v.Set("Body", SMStext)
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(AccountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			context.SetOutput(ovsend, true)
			var sid string
			sid = data["sid"].(string)
			activityLog.Info("Twilio SMS SID: " + sid)
		}
	} else {
		context.SetOutput(ovsend, false)
		activityLog.Error("Twilio SMS Status: " + resp.Status)
	}

	return true, nil
}
