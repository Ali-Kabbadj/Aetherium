package api

import "aetherium.com/user-service/app/features/signup"

type Resolver struct {
	SignUpService *signup.SignUpService
}