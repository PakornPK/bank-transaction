package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimmiepr/bank-transaction/internal/model"
)

type Wallet struct {
	ID      uint    `json:"id,omitempty"`
	Balance float32 `json:"balance,omitempty"`
}

func (s *Service) GetBalance(c *gin.Context) {
	var wallet model.Wallet
	s.DB.First(&wallet, 1)

	c.JSON(200, Wallet{
		ID:      wallet.ID,
		Balance: wallet.Balance,
	})
}

type Transaction struct {
	ID     uint    `json: "id"`
	Typ    string  `json: "typ"`
	Amount float32 `json: "amount"`
}

func (s *Service) CreateTransaction(c *gin.Context) {
	var tx Transaction
	if err := c.BindJSON(&tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not bind JSON"})
		return
	}
	if tx.Amount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "amount cannot less than 0"})
		return
	}

	var wallet model.Wallet
	s.DB.First(&wallet, tx.ID)

	switch tx.Typ {
	case "deposit":
		wallet.Balance += tx.Amount
	case "withdraw":
		wallet.Balance -= tx.Amount
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tx type did not defined"})
		return
	}

	s.DB.Save(&wallet)

	c.JSON(200, Wallet{
		ID:      wallet.ID,
		Balance: wallet.Balance,
	})
}
