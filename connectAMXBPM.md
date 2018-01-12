# how to connect TIBCO ActiveMatix BPM
here a little Tutorial about how to connect the TIBCO Cloud Integration (TCI) Webintegrator Service with TIBCO AciveMatrix BPM.

## Business Studio Configuration (Designtime)
Flow Activity somewhere within a BPMN Flow 
![pic](screenshots/bpm/BStask.png?raw=true "Business Studio BPMN Flow Activity")

Flow Activity configuration, plus graphical Request/Response mapping 
![pic](screenshots/bpm/BSproperties.png?raw=true "Business Studio Activity Configuration")

REST .RSD File definition and .JSD Schema Files in REST Service Folder of Business Studio
![pic](screenshots/bpm/RESTcall.png?raw=true "REST Folder Files")

REST Configuration Overview in Business Studio Properties .RSD
![pic](screenshots/bpm/RESTconfig.png?raw=true "REST Configuration")

REST Request Schema .JSD
![pic](screenshots/bpm/requestSchema.png?raw=true "REST Request Schema")

REST Response Schema .JSD
![pic](screenshots/bpm/responseSchema.png?raw=true "REST Response Schema")

## ActiveMatrix (AMX) Service Grid Configuration (Runtime)
as AMX shared Resource defined httpClient and installed on the BPMNode
![pic](screenshots/bpm/AMXhttpClient.png?raw=true "AMX http Client")

as TCI is running on secured https, the SSL Configuration is needed for the AMX httpClient 
![pic](screenshots/bpm/AMXhttpClientSSL.png?raw=true "AMX http Client for SSL")

for SSL a SSL Client Provider Definition is needed
![pic](screenshots/bpm/AMXsslclient.png?raw=true "AMX SSL Client")

the SSL Client Provider need a Keystore Provider, you can use any password
![pic](screenshots/bpm/AMXkeystore.png?raw=true "AMX Keystore")

## Cloud Intergration Web Integrator (Microservice)
to have the complete picture here the Swagger 2.0 JSON File of the Service running on TIBCO Cloud Integration Webintegrator

[SMSsender.json](screenshots/bpm/SMSsender.json)

<i>Note: "basePath": replaced with "/<<your basePath>>" as you can download it from your own account
This JSON File is just attached to show BPM is configured this way.</i>