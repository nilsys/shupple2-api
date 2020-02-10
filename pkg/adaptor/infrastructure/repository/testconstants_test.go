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
)

var (
	categoryIDs     = []int{1501, 1502}
	addedCategoryID = 1503
	// review_test
	mockReviewUserID        = 1
	mockReviewInnID         = 1
	mockReviewTouristSpotID = 1
	mockReviewHashTag       = "dummy"
	mockReviewPerPage       = 10
	mockReviewPage          = 1

	bodies    = []string{"b1601", "b1602"}
	addedBody = "1603"

	lcategoryIDs     = []int{1701, 1702}
	addedLcategoryID = 1703

	touristSpotIDs     = []int{1801, 1802}
	addedTouristSpotID = 1803

	postIDs     = []int{1901, 1902}
	addedPostID = 1903

	sampleTime = time.Date(2020, 7, 7, 10, 0, 0, 0, time.Local)
)
