package domain

// Domain model specifically for telco account

type TelcoAccount struct {
	ID            int
	SIMNumber     int
	Name          string
	TelcoStatus   string
	AccountActive bool
}
