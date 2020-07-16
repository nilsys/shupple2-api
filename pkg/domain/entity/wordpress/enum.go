package wordpress

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(undefined, japan, world, theme)
*/
type CategoryType int

/*
ENUM(user = 1, post, location, movie, comic, feature, category, location__cat, revision, cf_project, cf_return_gift)
*/
type EntityType int

/*
ENUM(publish = 1, future, draft, pending, private, trash)
*/
type Status int

/*
ENUM(ReservedTicket = 1, Other)
*/
type GiftType int
