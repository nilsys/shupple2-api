package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Male = 1, Female)
*/
type Gender int

/*
ENUM(NEW = 1, RANKING)
*/
type MediaSortBy int

/*
ENUM(NEW = 1, RECOMMEND)
*/
type ReviewSortBy int

/*
ENUM(RANKING, RECOMMEND)
*/
type UserSortBy int

/*
ENUM(AreaGroup, Area, SubArea, SubSubArea, Theme)
*/
type CategoryType int

/*
ENUM(Area = 1, SubArea, SubSubArea, TouristSpot, HashTag, User)
*/
type SuggestionType int

/*
ENUM(BUISINESS = 1, COUPLE, FAMILY, FRIEND, ONLY, WITHCHILD)
*/
type AccompanyingType int
