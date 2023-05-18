package domain

//  Domain model specifically for telco settings

type TelcoSettings struct {
	ParentalControlActive bool
	SpamSMSActive         bool
	RoamingActive         bool
}
