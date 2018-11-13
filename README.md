# TCI WI Twilio Extension
[![Go Report Card](https://goreportcard.com/badge/github.com/JGrotex/tci-wi-twilio-extension)](https://goreportcard.com/report/github.com/JGrotex/tci-wi-twilio-extension) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first draft Version with just a SMS sender using Twilio API, more to come ...

The attached ZIP contains the first release v.1.0 and can just uploaded under TIBCO Cloud Integration Extensions

Please create your own 'free' Access Key on Twilio (http://www.twilio.com) to enter into the Activity Details.

## Sample Integration
TCI Sample APP Exports can be found in [TCIAppExport](TCIAppExport) Folder, they are all ready to use.
here a sample of how to [connect this Service](connectAMXBPM.md) from TIBCO ActiveMatrix BPM. 

## Activities
available Activities so far
### SMS Sender
Flow Screenshot inside TIBCO Cloud Integration Flogo / Web Integrator

![Twilio SMS image](screenshots/twilio-SMS-in-TCI-WebIntegrator.png?raw=true "TCI Flogo Twilio SMS Screenshot")

Input
- accountSID            string
- authToken             string
- FromPhonenumber       string
- ToPhonenumber         string
- SMStext               string

<i>Hint:</i> Sender is working as well if you specify only a String as FromPhonenumber e.g. 'TIBCO Casemanagement'.
So you don't need to order a SMS Sender Number from Twilio.

Output
- send               bool   `json:"send"`
  
Sample Input Data
your Twilio Account Data for accountSID and authToken
+49171.... 
"some text ..."

Sample Output Data

``json:
{"send":true}
``

<hr>
<sub><b>Note:</b> more TCI Flogo Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
