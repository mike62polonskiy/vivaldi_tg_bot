package main

//Group ...
type Group struct {
	GroupAfishaID string `json:"group_afisha_id"`
	GroupCity     string `json:"group_city"`
	GroupDomain   string `json:"group_domain"`
	GroupID       string `json:"group_id"`
	GroupName     string `json:"group_name"`
	GroupTAG      string `json:"group_tag"`
	GroupURL      string `json:"group_url"`
	ID            int `json:"id"`
}

//Event ...
type Event struct {
	EventContacts   string `json:"event_contacts"`
	EventDateTime   string `json:"event_datetime"`
	EventDesription string `json:"event_description"`
	EventDomain     string `json:"event_domain"`
	EventImage      string `json:"event_image"`
	EventName       string `json:"event_name"`
	EventURL        string `json:"event_url"`
	ID              int `json:"id"`
}


const (
	msgTemplate = `
Органазаторы в вашем городе:
-----------------------
{{range .}}
Город: {{.GroupCity}}
Название группы vk: {{.GroupName}}
Ссылка на группу организатора: {{.GroupURL}}
Тэг группы: {{.GroupTAG}}
{{end}}`

	eventMsgTemplate = `
Органазаторы в вашем городе:
-----------------------
{{range .}}
Название ивента: {{.EventName}}
Ссылка на ивент(vkontakte): {{.EventURL}}
Описание ивента: {{.EventDesription}}
Дата события: {{.EventDateTime}}
{{end}}`
) 