package wordpress

type Comic struct {
	ID       int             `json:"id"`
	Date     Time            `json:"date"`
	Modified Time            `json:"modified"`
	Slug     string          `json:"slug"`
	Status   string          `json:"status"`
	Type     string          `json:"type"`
	Link     string          `json:"link"`
	Title    Text            `json:"title"`
	Content  ProtectableText `json:"content"`
	Author   int             `json:"author"`
	Template string          `json:"template"`
}
