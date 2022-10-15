# golang-image-api

This repo contains the source code for a serverless API I have made with AWS.

The API has 3 parts: AWS Lambda, a S3 Bucket, and an API Gateway trigger for the Lambda function.

On the S3 bucket I've uploaded an image. The Lambda function lists that object as an url and performs a HTTP get request to obtain the image, and returns it as bytes.
The final response displays the image, and an API Gateway is added as a trigger to call the function and display the response.

Final API Gateway trigger endpoint: https://52dzs50o5i.execute-api.us-east-1.amazonaws.com/golang-image-api
