package models

type Address struct {
	ID       int
	Country  string
	City     string
	Street   string
	home     int
	apartmen int
}

type AddressList struct {
	Addresses []Address
}
