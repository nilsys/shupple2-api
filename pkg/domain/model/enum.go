package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Male = 1, Female)
*/
type Gender int

/*
ENUM(NEW = 1, RANKING, RECOMMEND)
*/
type MediaSortBy int

/*
ENUM(NEW = 1, RECOMMEND)
*/
type ReviewSortBy int

/*
ENUM(RANKING = 1, RECOMMEND)
*/
type UserSortBy int

/*
ENUM(Area = 1, SubArea, SubSubArea, TouristSpot, HashTag, User)
*/
type SuggestionType int

/*
ENUM(BUISINESS = 1, COUPLE, FAMILY, FRIEND, ONLY, WITHCHILD)
*/
type AccompanyingType int

/*
ENUM(Undefined, Japan, World)
*/
type AreaGroup int

/*
ENUM(Undefined, Area, SubArea, SubSubArea)
sub_sub_areaの子カテとその子孫がundefinedに分類される
*/
type AreaCategoryType int

/*
ENUM(Undefined, Theme = 1, SubTheme)
sub_themeの子カテとその子孫がundefinedに分類される
*/
type ThemeCategoryType int

/*
ENUM(Undefined, SpotCategory = 1, SubSpotCategory)
*/
type SpotCategoryType int

/*
ENUM(Undefined, Style, Scene, Gourmet, LifeStyle, Activity, Sport)
*/
type InterestGroup int

/*
ENUM(FAVORITE = 1, COMMENT, REPLY, FOLLOW, TAGGED)
*/
type NoticeActionType int

/*
ENUM(POST = 1, VLOG, REVIEW, COMMENT, REPLY, USER)
*/
type NoticeActionTargetType int

/*
ENUM(Transaction = 1)
*/
type ContextKey int

/*
ENUM(REVIEW = 1, COMMENT, REPLY)
*/
type ReportTargetType int

/*
ENUM(UNKNOWN =1, SEXUAL, INAPPROPRIATE)
*/
type ReportReasonType int
