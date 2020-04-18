package database

type MongoDBHandler interface {
	Put([]byte, []byte, ...interface{}) error
	Get(string, ...interface{}) ([]byte, error)
	Scan()
}
