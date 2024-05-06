package webhooks

import "log"

type WebHookBody struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Event   string `json:"event"`
	Time    string `json:"time"`
}

func (w *WebHookBody) Print() {
	log.Println("ID: ", w.ID)
	log.Println("Country: ", w.Country)
	log.Println("Event: ", w.Event)
	log.Println("Time: ", w.Time)
}
