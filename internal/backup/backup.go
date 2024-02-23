package backup

type Item struct {
	URL   string
	Link  string
	Title string
}

type Manager interface {
	Import([]byte) ([]Item, error)

	Export([]Item) ([]byte, error)
}
