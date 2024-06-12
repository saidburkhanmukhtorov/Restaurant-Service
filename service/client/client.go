package client // Assuming your client package

import (
	"fmt"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientStruct struct {
	Payment payment.PaymentServiceClient
}

func NewClient() (*ClientStruct, error) {

	paymentConn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to payment service: %v", err)
	}
	defer paymentConn.Close()
	return &ClientStruct{
		Payment: payment.NewPaymentServiceClient(paymentConn),
	}, nil
}
