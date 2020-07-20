package wordpress

type (
	CfProject struct {
		ID            int                 `json:"id"`
		Date          JSTTime             `json:"date"`
		Modified      UTCTime             `json:"modified"`
		Slug          URLEscapedString    `json:"slug"`
		Status        Status              `json:"status"`
		Type          string              `json:"type"`
		Link          string              `json:"link"`
		Title         Text                `json:"title"`
		Content       ProtectableText     `json:"content"`
		Author        int                 `json:"author"`
		FeaturedMedia int                 `json:"featured_media"`
		Categories    []int               `json:"categories"`
		Attributes    CfProjectAttributes `json:"acf"`
	}

	CfProjectAttributes struct {
		PhotoGallery struct {
			Thumbnails [][]*PhotoGalleryItem `json:"thumbnails"`
		} `json:"photo_gallery"`
		Summary   string    `json:"summary"`
		GoalPrice IntString `json:"goal_price"`
		Deadline  JSTTime   `json:"deadline"`
	}

	CfReturnGift struct {
		ID            int                    `json:"id"`
		Date          JSTTime                `json:"date"`
		Modified      UTCTime                `json:"modified"`
		Slug          URLEscapedString       `json:"slug"`
		Status        Status                 `json:"status"`
		Type          string                 `json:"type"`
		Link          string                 `json:"link"`
		Title         Text                   `json:"title"`
		Content       ProtectableText        `json:"content"`
		Author        int                    `json:"author"`
		FeaturedMedia int                    `json:"featured_media"`
		Attributes    CfReturnGiftAttributes `json:"acf"`
	}

	CfReturnGiftAttributes struct {
		CfProject    RelatedPost     `json:"cf_project"`
		SortOrder    IntString       `json:"sort_order"`
		GiftType     GiftType        `json:"gift_type"`
		Price        IntString       `json:"price"`
		FullAmount   IntString       `json:"full_amount"`
		Deadline     NullableJSTTime `json:"deadline"`
		DeliveryDate string          `json:"delivery_date"`
		IsCancelable bool            `json:"is_cancelable"`
	}
)
