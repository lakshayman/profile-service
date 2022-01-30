package main

import (
	"fmt"
	// "context"
	// "log"

	// firebase "firebase.google.com/go"
	// "google.golang.org/api/option"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var (
	secretCache, _ = secretcache.New()
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, _ := secretCache.GetSecretString("Password")
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
