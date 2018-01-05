# TCI WI Twilio Extension
first draft Version with just a SMS sender using Twilio API, more to come ...

The attached ZIP contains the first release v.1.0 and can just uploaded under TIBCO Cloud Integration Extensions

Please create your own 'free' Access Key on Twilio to enter into the Activity Details.
A TCI WI Connector is planed for later. 

## Activities
available Activities so far
### SMS Sender
Flow Screenshot inside TIBCO Cloud Integration Web Integrator

![Twilio SMS image](screenshots/twilio-SMS-in-TCI-WebIntegrator.png?raw=true "TCI WI Twilio SMS Screenshot")

Input
- accountSID            string
- authToken             string
- FromPhonenumber       string
- ToPhonenumber         string
- SMStext               string

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
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
