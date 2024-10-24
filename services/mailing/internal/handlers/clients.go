package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

type ClientData struct {
	Clients []models.Client
}

func HandleClientsView(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewClient()
	clients, err := repo.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientData := ClientData{Clients: clients}
	err = templates.ExecuteTemplate(w, "customers.html", clientData)
	if err != nil {
		log.Println(err)
	}
}

func HandleAddClientView(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "add-customer.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func HandleAddClient(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")
	phone := r.Form.Get("phone")
	token := r.Form.Get("token")

	repo := repositories.NewClient()
	err = repo.Insert(models.Client{
		Name:      name,
		Phone:     phone,
		Token:     token,
		Connected: false,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleExtendClientDate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rawId := r.Form.Get("id")
	rawDuration := r.Form.Get("duration")
	unit := r.Form.Get("unit")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	duration, err := strconv.Atoi(rawDuration)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if unit == "" || (unit != "months" && unit != "days") {
		log.Println(err)
		http.Error(w, "Ай немесе күнді таңдаңыз", http.StatusInternalServerError)
		return
	}

	repo := repositories.NewClient()
	err = repo.Extend(id, duration, unit)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleDeactivate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rawId := r.Form.Get("id")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repo := repositories.NewClient()
	err = repo.Deactivate(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
