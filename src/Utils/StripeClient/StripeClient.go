package StripeClient

import (
	"os"

	"github.com/stripe/stripe-go"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_KEY")
}
