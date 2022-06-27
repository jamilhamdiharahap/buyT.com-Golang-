package models

type Kereta1 struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga uint64 `json:"harga"`
}

type ResponseServiceKereta struct {
	User
	Kereta `json:"data"`
}
