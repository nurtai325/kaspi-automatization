package queue

type OrderQueue interface {
	Add(id string) error
	Remove(id string) error
	Get() ([]string, error)
	Count() (int, error)
	GetPage(page int, size int) (string, error)
}
