# Lambda Create Account 
## Teste local

```shell
# caso seja windows
alias sam="sam.cmd"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
sam local invoke "CreateAccount" -e ../../../test/events/account_manager.json
```


## Deploy manual

```shell
alias sam="sam.cmd"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
sam build
sam deploy --no-confirm-changeset --no-fail-on-empty-changeset
```


## Configuration sam cli template.yaml
```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: >
  Function Create Account

Resources:
  CreateAccount:
    Type: AWS::Serverless::Function
    Properties:
      Environment:
        Variables:
          DATABASE_MONGODB_URL: "mongodb+srv://admin:qpal10zm@carbon-project-cluster.cpxiruo.mongodb.net"
          EMAIL_ADMIN: "dev.diego.morais@gmail.com"
          EMAIL_SUBJECT_CREATE_ACCOUNT: "Your account was created"
      Timeout: 15
      MemorySize: 128
      FunctionName: "account_manager"
      CodeUri: .
      Handler: main
      Runtime: go1.x
      Architectures:
        - x86_64

```