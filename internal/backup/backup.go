package backup

type BackupItem struct {
	URL   string
	Link  string
	Title string
}

type BackupManager interface {
	Import([]byte) ([]BackupItem, error)

	Export([]BackupItem) ([]byte, error)
}
