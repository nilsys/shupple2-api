package entity

//go:generate go-enum -f=enum.go --marshal

/*
ENUM(Male, Female)
*/
type Gender int
