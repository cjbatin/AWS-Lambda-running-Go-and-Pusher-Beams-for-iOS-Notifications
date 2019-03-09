package main
import (
        "fmt"
        "context"
				"github.com/pusher/push-notifications-go"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Title string `json:"title"`
				Message string `json:"message"`
}

const (
  instanceId = "YOUR_INSTANCE_ID"
  secretKey  = "YOUR_SECRET_KEY"
)

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	beamsClient, _ := pushnotifications.New(instanceId, secretKey)

	publishRequest := map[string]interface{}{
		"apns": map[string]interface{}{
			"aps": map[string]interface{}{
				"alert": map[string]interface{}{
					"title": event.Title,
					"body":  event.Message,
				},
			},
		},
	}

	pubId, err := beamsClient.PublishToInterests([]string{"hello"}, publishRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publish Id:", pubId)
	}
	return fmt.Sprintf("Completed"), nil
}

func main() {
  lambda.Start(HandleRequest)
}
