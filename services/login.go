package services

import "tangy-inn/tangy-inn_backend/models"

// Login Interface
type Login interface {
	FindByUser(models.LoginDetail) (*models.User, error)
}
