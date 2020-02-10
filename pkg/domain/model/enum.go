package model

//go:generate go-enum -f=enum.go --marshal

/*
ENUM(Male, Female)
*/
type Gender int

/*
ENUM(NEW, RANKING)
*/
type SortBy int
