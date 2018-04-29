# Basic AWS SAM template with GO Lambda function

1. Install aws-cli, go and aws-sam-local (use npm or download binary)
2. Create a S3 bucket for the code package
3. Update the s3 bucket name and stack name in package.json under scripts
4. Get the go dependencies (`go get ./...` or see Findings --> 8)
5. `yarn run build` --> cross compiles the go code for Linux to the ./dist folder
6. `yarn run package` --> creates a cloudformation template (serverless-output.yaml), zips the go executable and uploads it to the s3 bucket
7. `yarn run deploy` --> executes the cloudformation template using the code on S3

To test locally:

1. Build the go executable: `yarn run build` --> todo: create a watcher
2. `sam local start-api`
3. Api is accessable on http://127.0.0.1:3000 (probably, check console)

To use the env variables file: `sam local start-api --env env-vars.json` 

Todo: how to define parameters when using sam local start-api?
Doesn't work: sam local start-api --env-vars env-vars.json --parameter-values 'GlobalVarTest=hallo' (still gives globaldefault as defined  in template.yaml)

The go executable is automatically reloaded after building. The local test api accepts having a request body on a GET request, when deployed this generates a 403 bad request.

Findings:

1. Aws-sam-local doesn't install with yarn and doesn't install on WSL (windows only problems). Binaries are available.
2. Go packages can only have one handler. Currently using a switch statement to "route" the requests. Maybe there is some middleware available to handle this better.
3. The go function must return a APIGatewayProxyResponse (see: https://github.com/aws/aws-lambda-go/blob/master/events/README_ApiGatewayEvent.md)
4. The go function must be executable when zipped (windows only problem again)
5. You can use global variables to pool database connections (https://docs.aws.amazon.com/lambda/latest/dg/best-practices.html).
6. Response times are around ~70ms
7. You must use AWS::Serverless::Function or AWS::Serverless::Api in the template file for the API. Other cloudformation resources can be used but the Serverless syntax is used to test the API locally.
8. Go code has to be in the GOPATH which means your whole project is forced inside the GOPATH. If you place the code outside the gopath you have to get the dependencies manually:
    - go get github.com/aws/aws-lambda-go/events
    - go get github.com/aws/aws-lambda-go/lambda
    - etc.
9. To speed up the local api it is neccesary to set AWS Api keys (https://github.com/awslabs/aws-sam-local/pull/159, from docs: https://docs.aws.amazon.com/lambda/latest/dg/test-sam-local.html):   
   As with the AWS CLI and SDKs, SAM Local looks for credentials in the following order:

   - Environment variables (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)

   - The AWS credentials file, located at ~/.aws/credentials on Linux, MacOS, or Unix, or at C:\Users\USERNAME \.aws\credentials on Windows)

   - Instance profile credentials, if running on an Amazon EC2 instance with an assigned instance role
