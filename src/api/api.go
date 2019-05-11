package api

type API struct {
	notify chan Notification // Sends notifications out to anyone connected to the websocket
}

func Create() API {
	api := API{
		notify: make(chan Notification),
	}

	return api
}

func (api API) QueueNotification(notif Notification) {
	api.notify <- notif
}
