// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ThreadUpdate Сообщение для обновления ветки обсуждения на форуме.
// Пустые параметры остаются без изменений.
//
// swagger:model ThreadUpdate
type ThreadUpdate struct {

	// Описание ветки обсуждения.
	Message string `json:"message,omitempty"`

	// Заголовок ветки обсуждения.
	Title string `json:"title,omitempty"`
}