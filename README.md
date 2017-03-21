alexa triggered lambda function written in go 

required: 
go version > 1.4 
Twilio account /keys
lambda function arn
alexa skill (the alexa skill endpoint is the apex-generated lambda function)
set AWS Access Key ID, Secret Access Key, and location

to deploy function to lambda:
apex deploy 

run the lambda function:
apex invoke sender

run lambda funtion with input:
cat alexa_model/example_lambda_request.json | apex invoke sender
 
after running this, you can see the backend logs w/:
apex logs sender
