# Serverless with Go

Testing out lambda functions in Go with Terraform

## DynamoDB
Tablename: tasks

|id|title|description|completed|
|---|---|---|---|

## Lambda
### create_task
Creates a task and puts it into DynamoDB

## Terraform resources
- Policies:
    - sts:AssumeRole
    - logs:CreateLogGroup
    - logs:CreateLogStream
    - logs:PutLogEvents
    - dynamodb:PutItem
- Zip source code
- Lambda function
- DynamoDB table

