package wordpress

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(undefined, japan, world, theme)
*/
type CategoryType int

/*
ENUM(user = 1, post, location, movie, comic, feature, category, location__cat)
*/
type EntityType int

/*
ENUM(publish = 1, future, draft, pending, private)
*/
type Status int
