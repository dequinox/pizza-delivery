package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pizza struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

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
		for item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/pizzas", getPizzas).Methods("GET")
	router.HandleFunc("/pizzas/{id}", getPizza).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
