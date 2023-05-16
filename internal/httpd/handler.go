package httpd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pandodao/botastic/api"
	"github.com/pandodao/botastic/internal/llms"
	"github.com/pandodao/botastic/models"
	"github.com/pandodao/botastic/pkg/chanhub"
	"github.com/pandodao/botastic/storage"
	"go.uber.org/zap"
)

type TurnTransmitter interface {
	GetTurnsChan() chan<- *models.Turn
}

type Handler struct {
	logger          *zap.Logger
	llms            *llms.Handler
	sh              *storage.Handler
	hub             *chanhub.Hub
	turnTransmitter TurnTransmitter
}

func NewHandler(sh *storage.Handler, llms *llms.Handler, hub *chanhub.Hub, turnTransmitter TurnTransmitter, logger *zap.Logger) *Handler {
	return &Handler{
		logger:          logger.Named("httpd/handler"),
		llms:            llms,
		sh:              sh,
		hub:             hub,
		turnTransmitter: turnTransmitter,
	}
}

func (h *Handler) respErr(c *gin.Context, statusCode int, err error, codes ...api.ErrorCode) {
	code := statusCode
	if len(codes) > 0 {
		code = int(codes[0])
	}
	c.JSON(code, api.NewErrorResponse(code, err.Error()))
}

func (h *Handler) respData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, api.NewSuccessResponse(data))
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.String(200, "OK")
}

func (h *Handler) ListModels(c *gin.Context) {
	h.respData(c, api.ListModelsResponse{
		ChatModels:      h.llms.ChatModels(),
		EmbeddingModels: h.llms.EmbeddingModels(),
	})
}