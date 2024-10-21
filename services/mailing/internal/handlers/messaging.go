package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/messaging"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
	qrcode "github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

func HandleConnectQrcode(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.Form.Get("id")
	clientId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	phone := r.Form.Get("phone")
	if phone == "" {
		http.Error(w, "phone form value is not present", http.StatusBadRequest)
		return
	}

	qr, err := getQr(clientId, phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "qrcode.html", struct {
		Qr string
	}{Qr: qr})
	if err != nil {
		log.Println(err)
	}
}

func getQr(clientId int, phone string) (string, error) {
	dbConn := db.New()
	container := sqlstore.NewWithDB(dbConn, "postgres", nil)
	err := container.Upgrade()
	if err != nil {

	}

	device := container.NewDevice()
	client := whatsmeow.NewClient(device, nil)

	qrChan, _ := client.GetQRChannel(context.Background())
	err = client.Connect()
	if err != nil {
		panic(err)
	}

	conf := config.New()
	qrAsset := fmt.Sprintf("assets/%s.png", phone)

	for evt := range qrChan {
		if evt.Event != "code" {
			continue
		}

		err = qrcode.WriteFile(evt.Code, qrcode.Medium, 512, conf.WORK_DIR+"/"+qrAsset)
		if err != nil {
			return "", err
		}

		go handleQrScan(client, qrChan, clientId, phone)
		break
	}

	return qrAsset, nil
}

func handleQrScan(client *whatsmeow.Client, qrChan <-chan whatsmeow.QRChannelItem, clientId int, phone string) {
	for res := range qrChan {
		if res.Error != nil {
			return
		}
		if res == whatsmeow.QRChannelSuccess {
			break
		}
	}

	repo := repositories.NewClient()
	err := repo.ConnectWh(clientId)
	if err != nil {
		log.Println(err)
	}

	messaging.Add(phone, client)
}
