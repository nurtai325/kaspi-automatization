package messaging

import (
	"fmt"
	"strings"

	"github.com/nurtai325/kaspi/mailing/internal/models"
)

const (
	reviewLinkBase = "https://kaspi.kz/shop/review/productreview?rating=5"
)

func NewOrderMessage(name, orderCode, shop, phone string, entries []models.Entry) string {
	text := `Қайырлы күн, %s!
%s магазиніне тапсырыс бергеніңіз үшін рақмет
Сіздің тапсырысыңыз:   %s.
Жеткізу көрсетілген күні жүзеге асырылады.
Тапсырыс нөмірі: %s
Жұмыс уақыты: 10:00-20:00 дейын
Сұрақтар бойнша осы %s номерге хабарласыңыз

Тапсырысыңызды тез арада жинап, сізге жібереміз.
Саудаңыз сәтті болсын!


*

Добрый день, %s!
Спасибо за заказ в магазине %s
Вы заказали: %s.
Доставка будет осуществлена в указанную дату.
Номер заказа: %s
График работы с 10:00 до 20:00
По всем вопросам обращайтесь по %s этому номеру 
В ближайшее время мы соберём заказ и отправим вам.
Хороших покупок!`
	message := fmt.Sprintf(
		text,
		name,
		shop,
		parseEntries(entries, true),
		orderCode,
		"+"+phone,
		name,
		shop,
		parseEntries(entries, false),
		orderCode,
		"+"+phone,
	)
	return message
}

func CompletedOrderMessage(name, orderCode, productCode, shop string) string {
	reviewLink := fmt.Sprintf("%s&orderCode=%s&productCode=%s", reviewLinkBase, orderCode, productCode)
	text := `Қайырлы күн, %s!
%s дүкенінен сатып алуыңызбен құттықтаймыз!
Сізге бәрі ұнады деп үміттенеміз.
Сілтеме арқылы өтіп, біздің дүкеннің атауын көрсете отырып, пікір қалдыра аласыз ба, бұл біз үшін маңызды⤵️:
%s
Сілтеме белсенді болуы үшін жауап ретінде бірдеңе жазыңыз.

🔸🔸🔸

Добрый день, %s!
Поздравляем с покупкой с магазина %s!
Мы надеемся, что вам все понравилось.
Если вам не сложно, пожалуйста оставьте отзыв с указанием названия нашего магазина перейдя по ссылке⤵️:
%s
Чтобы ссылка стала активной, напишите пожалуйста что-нибудь в ответ.`
	message := fmt.Sprintf(text, name, shop, reviewLink, name, shop, reviewLink)
	return message
}

func parseEntries(entries []models.Entry, kz bool) string {
	parsed := strings.Builder{}
	length := len(entries)

	for i, entry := range entries {
		parsedEntry := ""
		if kz {
			parsedEntry = fmt.Sprintf(
				"%s, %d дана, бағасы: %d",
				entry.ProductName,
				entry.Quantity,
				entry.Price,
			)
		} else {
			parsedEntry = fmt.Sprintf(
				"%s, количество: %d, цена: %d",
				entry.ProductName,
				entry.Quantity,
				entry.Price,
			)
		}
		parsed.WriteString(parsedEntry)
		if i+1 != length {
			parsed.WriteString(",")
		}
	}
	parsed.WriteString(".")

	return parsed.String()
}
