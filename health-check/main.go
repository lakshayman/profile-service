package main

import (
	"fmt"
	// "context"
	// "log"

	// firebase "firebase.google.com/go"
	// "google.golang.org/api/option"

	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func getSecret() string {
	secretName := "Password"
	region := "us-east-2"

	//Create a Secrets Manager client
	sess, err := session.NewSession()
	if err != nil {
		// Handle session creation error
		fmt.Println(err.Error())
		return "a"
	}
	svc := secretsmanager.New(sess,
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return ""
	}

	// Decrypts secret using the associated KMS key.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
		fmt.Println("here :" + secretName + " " + secretString)
		return secretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return ""
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
		return decodedBinarySecret
	}

	// Your code goes here.
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var result string = getSecret()
	// ctx := context.Background()
	// sa := option.WithCredentialsJSON([]byte(os.Getenv("firestoreCred")))
	// app, err := firebase.NewApp(ctx, nil, sa)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// client, err := app.Firestore(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// client.Collection("users").Add(ctx, map[string]interface{}{
	// 	"first": "Ada",
	// 	"last":  "Lovelace",
	// 	"born":  1815,
	// })
	// defer client.Close()

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Got It : " + result),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
