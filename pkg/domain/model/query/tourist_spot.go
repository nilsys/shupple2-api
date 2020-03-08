package query

type (
	// repositoryで使用するクエリ発行用構造体
	FindTouristSpotListQuery struct {
		AreaID         int
		SubAreaID      int
		SubSubAreaID   int
		SpotCategoryId int
		ExcludeSpotIDs []int
		Limit          int
		OffSet         int
	}

	FindRecommendTouristSpotListQuery struct {
		ID                    int
		TouristSpotCategoryID int
		Limit                 int
		OffSet                int
	}
)
