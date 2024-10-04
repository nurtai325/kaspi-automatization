package messaging

import "fmt"

const (
	reviewLinkBase = "https://kaspi.kz/shop/review/productreview?rating=5"
)

func NewOrderMessage(name, orderCode string) string {
	message := fmt.Sprintf("hello, %s \norder code: %s \n", name, orderCode)
	return message
}

func CompletedOrderMessage(name, orderCode, productCode string) string {
	reviewLink := fmt.Sprintf("%s&orderCode=%s&productCode=%s", reviewLinkBase, orderCode, productCode)
	message := fmt.Sprintf("hello, %s, please leave a review here: %s\n", name, reviewLink)
	return message
}
