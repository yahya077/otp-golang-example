package smsProvider

import "fmt"

type MockedSmsSenderProviderClient struct {
}

// SendOtp ISmsProvider interfaces function
func (m MockedSmsSenderProviderClient) SendOtp(phone string, code string) error {
	// send sms with your provider here
	fmt.Printf("Provider is sent Otp code to your phone.\n Phone: %s, Otp: %s \n", phone, code)
	return nil
}
