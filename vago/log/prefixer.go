package log

type Prexifer interface {
	prefix(string) string
}
