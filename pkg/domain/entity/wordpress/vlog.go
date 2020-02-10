package wordpress

type Vlog struct {
	ID            int             `json:"id"`
	Date          Time            `json:"date"`
	Modified      Time            `json:"modified"`
	Slug          string          `json:"slug"`
	Status        string          `json:"status"`
	Type          string          `json:"type"`
	Link          string          `json:"link"`
	Title         Text            `json:"title"`
	Content       ProtectableText `json:"content"`
	Author        int             `json:"author"`
	FeaturedMedia int             `json:"featured_media"`
	Template      string          `json:"template"`
	Categories    []int           `json:"categories"`
	Attributes    VlogAttributes  `json:"acf"`
}
type Creator struct {
	ID              int    `json:"ID"`
	UserFirstname   string `json:"user_firstname"`
	UserLastname    string `json:"user_lastname"`
	Nickname        string `json:"nickname"`
	UserNicename    string `json:"user_nicename"`
	DisplayName     string `json:"display_name"`
	UserEmail       string `json:"user_email"`
	UserURL         string `json:"user_url"`
	UserRegistered  string `json:"user_registered"`
	UserDescription string `json:"user_description"`
	UserAvatar      string `json:"user_avatar"`
}
type VlogAttributes struct {
	Youtube       string  `json:"youtube"`
	Series        string  `json:"series"`
	Creator       Creator `json:"creator"`
	CreatorSns    string  `json:"creator_sns"`
	FilmEdit      string  `json:"film_edit"`
	MovieLocation []int   `json:"movie_location"`
	YearMonth     string  `json:"year_month"`
	RunTime       string  `json:"run_time"`
	MovieTimeline string  `json:"movie_timeline"`
}
