package repository

import "time"

const (
	userID        = 101
	postID        = 111
	touristSpotID = 121
	comicID       = 131
	vlogID        = 141
	featureID     = 151
	categoryID    = 161
	lcategoryID   = 171
	reviewID      = 181
	hashtagID     = 191
	innID         = 201

	mockReviewPerPage = 10
	mockReviewPage    = 1
)

var (
	categoryIDs     = []int{1501, 1502}
	addedCategoryID = 1503

	bodies    = []string{"b1601", "b1602"}
	addedBody = "1603"

	lcategoryIDs     = []int{1701, 1702}
	addedLcategoryID = 1703

	touristSpotIDs     = []int{1801, 1802}
	addedTouristSpotID = 1803

	postIDs     = []int{1901, 1902}
	addedPostID = 1903

	hashtagIDs     = []int{2001, 2002}
	addedHashtagID = 2003

	sampleTime = time.Date(2020, 7, 7, 10, 0, 0, 0, time.Local)
)
