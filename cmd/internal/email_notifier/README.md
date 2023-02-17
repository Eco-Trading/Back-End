# Lambda Notifier Email 
## Teste local

```shell
# caso seja windows
alias sam="sam.cmd"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
sam local invoke "EmailNotifier" -e ../../../test/events/email_notifier.json
```


## Deploy manual

```shell
alias sam="sam.cmd"
sam build
sam deploy \
--region us-east-1 \
--s3-prefix EmailNotifier \
--s3-bucket deploy-sam-aws \
--stack-name EmailNotifier \
--capabilities "CAPABILITY_IAM" \
--no-confirm-changeset --no-fail-on-empty-changeset
```


## Configuration sam cli template.yaml
```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: >
  Function Notifier by email

Resources:
  EmailNotifier:
    Type: AWS::Serverless::Function
    Properties:
      Timeout: 15
      MemorySize: 128
      FunctionName: "email_notifier_dev"
      CodeUri: .
      Handler: main
      Runtime: go1.x
      Architectures:
        - x86_64
```