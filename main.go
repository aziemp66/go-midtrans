package main

import (
	"fmt"

	midtrans "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func main() {
	midtrans.ServerKey = "your server key"

	var s snap.Client
	s.New(midtrans.ServerKey, midtrans.Sandbox)
	s.Options.SetPaymentAppendNotification("https://example.com/append_notification")
	s.Options.SetPaymentOverrideNotification("https://example.com/override_notification")

	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "order-6969609",
			GrossAmt: 55000,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "item-1",
				Name:  "Adidas f50",
				Price: 55000,
				Qty:   1,
			},
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, _ := s.CreateTransactionUrl(req)
	fmt.Println("Response :", snapResp)

}
