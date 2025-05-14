package config

const LoginPageURL = "/user/login"

type Credentials struct {
	Login    string
	Password string
}

func GetExistingCredentials() Credentials {
	return Credentials{
		Login:    "test_login",
		Password: "123456",
	}
}

func GetNonExistingCredentials() Credentials {
	return Credentials{
		Login:    "test_invalid12ss3123",
		Password: "1234561212s3sda",
	}
}
