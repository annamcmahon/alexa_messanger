# Alexa Golang Messenger 
Alexa triggered lambda function to send texts written in go 

## How to build
**1. Download the following:**
* [go](https://golang.org/dl/) version > 1.4 
* [apex](https://github.com/apex/apex)
* [Alexa Messenger](https://github.com/annamcmahon/alexa_messanger)

**2. Twilio account setup:**
* 2.1 Make a [twilio](https://www.twilio.com) account (free/trial). 
* 2.2 Fill in your Twilio accountSid and authToken in the main.go file.

**3. Lambda functions setup:**
* 3.1 Make an AWS account
* 3.2 Get your lambda functions' ARN. The easiest way I found to do this was to create an AWS lambda function [here](https://console.aws.amazon.com/lambda). The ARN will then appear in the upper right hand corner of that functions viewing page.
* 3.3 Replace the "role" with your lambda function's arn in the project.json file

**4. Deploy and enable the lambda function:**
* 4.1 Deploy the lambda function by running the following in the root folder of Alexa Messenger:
```
apex deploy
```
* 4.2 Go to your the [aws lambda functions](https://console.aws.amazon.com/lambda) and confirm that a new function was created
* 4.3 Enable lambda function to be an Alexa endpoint. To do this we have to add the trigger "Alexa Skills Kit" our lambda function.

**6. Setup Alexa skill:**
* Next we have to make the alexa skill whose endpoint is the apex-generated lambda function
* 6.1 Make an alexa skill [here](https://developer.amazon.com/edw/home.html). The interaction model and intent schema to be used can be found in alexa_model. In configuration the AWS ARN should include the name of the function, such as: arn:aws:lambda:us-east-1:703378707126:function:go_sender. This points your alexa skill to your lambda function end point.

## Testing with apex
deploy lambda function: 
```
apex deploy
```
run the lambda function: 
```
apex invoke sender
```
run lambda funtion with input: 
```
cat alexa_model/example_lambda_request.json | apex invoke sender
```
after running this, you can see the backend logs w/: 
```
apex logs sender
```


## Tips/ pitfalls/ Other
* make sure your AWS Access Key ID, Secret Access Key, and location are set in your AWS config file
* A note on Apex: Apex uses a Node.js shim for non-native language support. This is a very small program which executes in a child process, and feeds Lambda input through STDIN, and program output through STDOUT. Because of this STDERR must be used for logging, not STDOUT.
* You can use [cloudwatch](https://console.aws.amazon.com/cloudwatch) to look at the backend output logs. To do this, make sure to give your Lambda function's execution role permissions to log to CloudWatch Logs. To do so: lambda>function>function_name>config>create a role with permissions

## The unique part of golang used 
We used goroutines to send multiple messages 



