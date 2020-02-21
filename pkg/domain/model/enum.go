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

/*
ENUM(AreaGroup, Area, SubArea, SubSubArea, Theme)
*/
type CategoryType int

/*
ENUM(Area, SubArea, SubSubArea, TouristSpot, HashTag, User)
*/
type SuggestionType int
