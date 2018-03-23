package main

import (
	"fmt"
	"errors"
	"encoding/base64"
    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/fivenp/go-identicon"
	"bytes"
	"image/png"
	"strings"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("got new request '%s'\n", request.Path)

	args := strings.Split(request.Path, "/")
	args = args[1:]

	if len(args) != 2 {
		fmt.Errorf("500 Server Error: Wrong args count (max 1 arg!)")
		return events.APIGatewayProxyResponse{}, errors.New("500 Server Error: Wrong args count (max 1 arg!)")
	} else {
		fmt.Printf("got args '%s'\n", args[1])
	}

	item := args[1]

	code := identicon.Code(item)
	size := 1024
	settings := identicon.DefaultSettings()
	img := identicon.Render(code, size, settings)

	fmt.Printf("creating identicon for '%s'\n", item)

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		fmt.Errorf("500 Server Error: Could not encode image. Details: %s", err)
		return events.APIGatewayProxyResponse{}, errors.New("500 Server Error: Could not encode image.")
	}

	// TODO - Check for ?base64 query string and provide binary return?
	response := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{"Content-Type": "image/png"}, Body: response, IsBase64Encoded: true}, nil
}

func main() {
	lambda.Start(Handler)
}
