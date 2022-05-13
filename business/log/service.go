package log

type Repository interface {
	GetAllLog() (logs []*Log, err error)
	CreateLog(log *Log) (*Log, error)
}
