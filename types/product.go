package types
type Product struct {
	Id          string
	Name        string
	Description string
	RetailerId string
	Url string
	Price       string
	Currency    string
	IsHidden   string
}

type ProductImg struct {
	Url string
}
