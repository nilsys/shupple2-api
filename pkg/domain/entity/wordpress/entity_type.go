package wordpress

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(post = 1, location, movie, comic, feature, category, location__cat)
*/
type EntityType int
