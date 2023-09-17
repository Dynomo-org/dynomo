package usecase

type BuildKeystoreParam struct {
	AppID         string
	FullName      string
	Organization  string
	Country       string
	Alias         string
	KeyPassword   string
	StorePassword string
}
