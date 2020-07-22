package repository

import "time"

const (
	userID                 = 101
	postID                 = 111
	touristSpotID          = 121
	comicID                = 131
	vlogID                 = 141
	featureID              = 151
	areaCategoryID         = 161
	themeCategoryID        = 171
	spotCategoryID         = 181
	reviewID               = 191
	hashtagID              = 201
	innID                  = 211
	shippingAddressID      = 221
	cfProjectID            = 231
	cfProjectCommentID     = 241
	metasearchAreaID       = 251
	cardID                 = 261
	paymentID              = 271
	cfReturnGiftID         = 281
	cfReturnGiftSnapshotID = 291
	cfProjectSnapshotID    = 301
)

var (
	sampleTime          = time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local)
	areaCategoryIDs     = []int{1401, 1402}
	addedAreaCategoryID = 1403

	themeCategoryIDs     = []int{1501, 1502}
	addedThemeCategoryID = 1503

	bodies    = []string{"b1601", "b1602"}
	addedBody = "1603"

	spotCategoryIDs     = []int{1701, 1702}
	addedSpotCategoryID = 1703

	touristSpotIDs     = []int{1801, 1802}
	addedTouristSpotID = 1803

	postIDs     = []int{1901, 1902}
	addedPostID = 1903

	hashtagIDs     = []int{2001, 2002}
	addedHashtagID = 2003

	userIDs     = []int{2101, 2102}
	addedUserID = 2103

	thumbnails = []string{"thumbnail1", "thumbnail2"}
)
