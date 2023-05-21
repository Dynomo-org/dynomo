package usecase

import "dynapgen/repository/redis"

type GenerateStoreParam struct {
	AppID         string
	FullName      string
	Organization  string
	Country       string
	Alias         string
	KeyPassword   string
	StorePassword string
}

type Keystore struct {
	Status       uint8
	URL          string
	ErrorMessage string
}

func convertKeystoreFromRepo(keystore redis.Keystore) Keystore {
	return Keystore{
		Status:       uint8(keystore.Status),
		URL:          keystore.URL,
		ErrorMessage: keystore.ErrorMessage,
	}
}
