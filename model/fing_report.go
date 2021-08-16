package model

//FingReport struct
type FingReport struct {
	BotDateTime    string
	NetworkAddress string
	NetworkFamily  string
	NetworkBearer  string
	Hosts          []Host
}

//Host struct
type Host struct {
	Name            string
	Hostname        string
	HardwareAddress string
	Address         string
	Vendor          string
	State           string
	LastChangeTime  string
}
