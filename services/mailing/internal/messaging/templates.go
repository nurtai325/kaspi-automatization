package messaging

import (
	"fmt"

	"github.com/nurtai325/kaspi/mailing/internal/models"
)

const (
	reviewLinkBase = "https://kaspi.kz/shop/review/productreview?rating=5"
)

func NewOrderMessage(name, orderCode string, entries []models.Entry) string {
	text := `Сәлеметсіз бе, %s! 
Cіздің тапсырысыңыздың номері: %s.
Жақын арада тапсырысты қаптап сізге қарай жібереміз.

Здраствуйте, %s!
Ваш номер заказа: %s.
В скором времени упакуем и отправим его вам.`
	message := fmt.Sprintf(text, name, orderCode, name, orderCode)
	return message
}

func CompletedOrderMessage(name, orderCode, productCode string) string {
	reviewLink := fmt.Sprintf("%s&orderCode=%s&productCode=%s", reviewLinkBase, orderCode, productCode)
	text := `Қайырлы күн, %s!
Сатып алуыңызбен құттықтаймыз!
Сізге бәрі ұнады деп үміттенеміз.
Төмендегі сілтеме арқылы оң пікір қалдыра аласыз ба🤗

Добрый день, %s!
Поздравляем с покупкой!
Мы надеемся, что вам все понравилось.
Можете уделить минуту и оставить положительный отзыв 🤗
%s`
	message := fmt.Sprintf(text, name, name, reviewLink)
	return message
}
