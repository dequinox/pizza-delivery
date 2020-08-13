package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Pizza struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Cost        int    `json:"cost"`
}

type OrderDetail struct {
	PizzaID string `json:"pizzaID"`
	Amount  int    `json:"amount"`
}

type Order struct {
	ID        string        `json:"id"`
	AccountID string        `json:"accountID"`
	Cost      int           `json:"cost"`
	Details   []OrderDetail `json:"orderDetails"`
}

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
	City      string `json:"city"`
	Street    string `json:"street"`
	House     string `json:"house"`
}

var orders []Order

var accounts []Account

var pizzas = []Pizza{
	{
		ID:          "1",
		Title:       "Карамельный Ананас",
		Description: "Ананас, Соус Карамельный, Соус Сырный, Сыр Моцарелла",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/KARPNAPL/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "2",
		Title:       "Пепперони",
		Description: "Пепперони, Сыр Моцарелла",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/%D0%9F%D0%95%D0%9F%D0%95/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "3",
		Title:       "Маргарита Гурме",
		Description: "Свежие томаты, Сыр Моцарелла",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/_%D0%9C%D0%90%D0%A0%D0%93%D0%93%D0%A0%D0%9C/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "4",
		Title:       "Домино'c",
		Description: "Бекон, Говядина, Грибы, Болгарский перец, Курица, Лук, Пепперони, Свежие томаты, Свинина, Сыр Моцарелла, Сыр",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/_%D0%94%D0%9E%D0%9C%D0%98%D0%9D%D0%9E%D0%A1/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "5",
		Title:       "Маргарита",
		Description: "Сыр Моцарелла",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/%D0%9C%D0%90%D0%A0%D0%93/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "6",
		Title:       "Домашняя",
		Description: "Ветчина, Свежие томаты, Соус Сырный, Сыр Моцарелла",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/_DOMSV/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "7",
		Title:       "Гавайская",
		Description: "Ананас, Ветчина, Сыр Моцарелла",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/%D0%93%D0%90%D0%92%D0%90/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
	{
		ID:          "8",
		Title:       "4 Сыра",
		Description: "Соус Карбонара, Сыр Моцарелла, Сыр Пармезан, Сыр Роккфорти, Сыр Чеддер (тёртый)",
		ImageUrl:    "https://dpr-cdn.azureedge.net/api/medium/ProductOsg/Global/4CHEESE/NULL/270x270/RU?v=b22f041db98d8069c082bf44fbba7474-1597330200000",
	},
}

func getPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizzas)
}

func getPizza(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range pizzas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var accountOrders []Order
	for _, item := range orders {
		if item.AccountID == account.ID {
			accountOrders = append(accountOrders, item)
		}
	}

	json.NewEncoder(w).Encode(accountOrders)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newOrder Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newOrder.ID = strconv.Itoa(len(orders) + 1)
	orders = append(orders, newOrder)

	json.NewEncoder(w).Encode(newOrder)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAccount Account
	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAccount.ID = strconv.Itoa(len(accounts) + 1)
	accounts = append(accounts, newAccount)

	json.NewEncoder(w).Encode(newAccount)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, item := range accounts {
		if item.Email == account.Email {
			if item.Password == account.Password {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}

	http.Error(w, err.Error(), http.StatusUnauthorized)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/pizzas", getPizzas).Methods("GET")
	router.HandleFunc("/pizzas/{id}", getPizza).Methods("GET")

	router.HandleFunc("/orders", getOrders).Methods("POST")
	router.HandleFunc("/order", createOrder).Methods("POST")

	router.HandleFunc("/account/create", createAccount).Methods("POST")
	router.HandleFunc("/account/login", Authenticate).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
