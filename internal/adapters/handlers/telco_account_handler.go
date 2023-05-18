package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ZAF07/telco/internal/core/domain"
	"github.com/ZAF07/telco/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// Handler (controller) would unmarshal request objects into required DTO for services (rpc)
// The handlers are the driving adapters. They take in requests objects from the outside world and triggers the internal services to execute their business logic. It is not tightly coupled with the internal services because it's dependency is abstracted into an interface allowing us to make changes to the handlers without worrying about propagating breakages

// Hanlders are the adapters to the core services, they implement the ports.service

type TelcoAccountHandler struct {
	TelcoService ports.ITelcoAccountService
}

func NewTelcoAccountHandler(svc ports.ITelcoAccountService) *TelcoAccountHandler {
	return &TelcoAccountHandler{
		TelcoService: svc,
	}
}

func (t *TelcoAccountHandler) CreateTelcoAccount(g *gin.Context) {
	ctx := context.Background()

	account := domain.TelcoAccount{
		Name:          "zaffere",
		SIMNumber:     1234,
		TelcoStatus:   "NORMAL",
		AccountActive: true,
	}

	success, err := t.TelcoService.CreateTelcoAccount(ctx, account)
	if err != nil {
		log.Fatalf("create accont service error : %+v", err)
	}

	g.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "Account creation success",
		"data": success,
	})
}
