package config

type Order struct {
	Login    string
	Password string
	Name     string
	Email    string
	Address  string
	Note     string
}

var ExistingUserToOrderData = Order{
	Login:    GetExistingCredentials().Login,
	Password: GetExistingCredentials().Password,
	Name:     "Name",
	Email:    "example@example.com",
	Address:  "Йошкар-Ола",
	Note:     "example",
}

var NonExistingUserToOrderData = Order{
	Login:    "fdsdlfnlsdnkfkdslflsdvsdf",
	Password: GetNonExistingCredentials().Password,
	Name:     "fdsdlfnlsdnkfkdslfasdsadfl",
	Email:    "fdsdlfnlsdnkfkdslfl@example.com",
	Address:  "Йошкар-Ола",
	Note:     "example",
}
