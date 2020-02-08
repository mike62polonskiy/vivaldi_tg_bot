package main

import (
	"log"
	"encoding/json"
)

func msgGroups() ([]Group, error) {
    sqlState := "SELECT * FROM vk_data_grub_vkgroups"
	byteJSON, err := sqlToJSON(sqlState)
	if err != nil {
		log.Panic(err)
	}

	var result []Group

	if err := json.Unmarshal(byteJSON, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func msgEvents() ([]Event, error) {
	sqlState := "SELECT * FROM vk_data_grub_events WHERE event_datetime >= current_date AND event_place = 'DEEP'"
	byteJSON, err := sqlToJSON(sqlState)
	if err != nil {
		log.Panic(err)
	}

	var result []Event

	if err := json.Unmarshal(byteJSON, &result); err != nil {
		return nil, err
	}
	return result, nil
}