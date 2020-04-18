build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o policebot

cf:
	AWS_REGION=ap-southeast-2 aws cloudformation create-stack --stack-name discord-bot-police --template-body file://./templates/dynamodb.yaml