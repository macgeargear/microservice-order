package payment

import (
	"context"

	"github.com/huseyinbabal/microservices-proto/golang/payment"
	"github.com/macgeargear/microservices-order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	paymet payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption // Data model for connection configurations

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // This is for disabling TLS
	conn, err := grpc.Dial(paymentServiceUrl, opts...)                            // Connects to service

	if err != nil {
		return nil, err
	}

	defer conn.Close()                       // Always close the connection before quitting the function
	client := payment.NewPaymentClient(conn) // Initializes the new payment stub instance

	return &Adapter{paymet: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.paymet.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})

	return err
}
