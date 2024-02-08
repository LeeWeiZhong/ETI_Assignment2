/*
 * @File: models.message.go
 * @Description: Defines Message information will be returned to the clients
 * 
 */
package models

// Message defines the response message
type Message struct {
	Message string `json:"message" example:"message"`
}
