package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Transaction = 1)
*/
type ContextKey int
