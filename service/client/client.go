package client // Assuming your client package

import (
	"log"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientStruct struct {
	Payment payment.PaymentServiceClient
}

func NewClient() (*ClientStruct, error) {

	paymentConn, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
		return nil, err
	}
	return &ClientStruct{
		Payment: payment.NewPaymentServiceClient(paymentConn),
	}, nil
}
