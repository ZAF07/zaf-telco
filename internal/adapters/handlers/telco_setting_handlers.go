package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ZAF07/telco/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type TelcoSettingHandler struct {
	TelcoService ports.ITelcoSettingService
}

func NewTelcoSettingHandler(svc ports.ITelcoSettingService) *TelcoSettingHandler {
	return &TelcoSettingHandler{
		TelcoService: svc,
	}
}

func (t *TelcoSettingHandler) GetTelcoSetting(g *gin.Context) {
	ctx := context.Background()

	result, err := t.TelcoService.GetTelcoSetting(ctx, "test")
	if err != nil {
		log.Fatalf("error getting result from setting service : %+v", err)
	}

	g.JSON(http.StatusOK, map[string]string{
		"msg":  "Success",
		"data": result,
	})
}
