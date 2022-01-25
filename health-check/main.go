package main

import (
	"fmt"
	// "context"
	// "log"
	"os"

	// firebase "firebase.google.com/go"
	// "google.golang.org/api/option"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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
		Body:       fmt.Sprintf(os.Getenv("PipelineUserSecretKey-86CzG9I7ZMgZ")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
