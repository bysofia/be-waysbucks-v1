package transactiondto

import "waysbucks/models"

type TransactionResponse struct {
	ID      int                         `json:"id"`
	Users   models.UsersProfileResponse `json:"userOrder"`
	Status  string                      `json:"status"`
	Product []models.ProductResponse    `json:"order"`
}
