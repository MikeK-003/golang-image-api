# golang-image-api

This repo contains the source code for a serverless API I have made with AWS.

The API has 3 parts: AWS Lambda, a S3 Bucket, and an API Gateway trigger for the Lambda function.

On the S3 bucket I've uploaded an image. The Lambda function lists that object as an url and performs a HTTP get request to obtain the image, and returns it as bytes.
The final response displays the image, and an API Gateway is added as a trigger to call the function and display the response.

Final API Gateway trigger endpoint: https://52dzs50o5i.execute-api.us-east-1.amazonaws.com/golang-image-api

To reproduce, you must have an AWS account, Golang installed, and Visual Studio Code.

First, download all the contents of this repository as a ZIP and extract it wherever you want.

Then, open AWS and create a S3 bucket. It can be named anything, but open the .go file and change the string var "bucketName" to your bucket's name. 
Also, change the string var "regionName" to your bucket's region.
Upload an image to it under any name you want.

Next, build the .go file to upload to AWS Lambda. If you are on Windows, type `$Env:GOOS = "linux"` in the VS terminal and then type `go build main.go`.
Take the new file and send it into a .zip file.

Open AWS Lambda and create a new function. Change its runtime to Go 1.x and upload its source code from the new .zip file. Change the runtime handler to "main".

Within the S3 bucket, disable blocking public access. 
Create a bucket policy that allows for the functions s3:GetObject and s3:ListBucket to work on the resources of the bucket and the picture.

Open AWS API Gateway and create a HTTP API. Go back to AWS Lambda and add in the new API Gateway as the function's trigger.

The resulting API endpoint should execute the Lambda function, which then calls on the bucket, obtains the file's url, and responds with the image in its response body.

![image](https://user-images.githubusercontent.com/102551944/196841774-abebebd9-cb75-4b93-90fe-5bbe914abcd3.png)
