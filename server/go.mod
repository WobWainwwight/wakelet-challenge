module wakelet-challenge/server

go 1.15

replace wakelet-challenge/nasa => ../nasa

replace wakelet-challenge/events-repository => ../events-repository

require (
	github.com/aws/aws-sdk-go v1.35.32 // indirect
	wakelet-challenge/events-repository v0.0.0-00010101000000-000000000000
	wakelet-challenge/nasa v0.0.0-00010101000000-000000000000
)
