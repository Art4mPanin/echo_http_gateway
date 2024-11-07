package infoservice

import (
	"Http-gateway/internal/config"
	"Http-gateway/internal/data/gen/info"
	"Http-gateway/pkg/singleton"
	grpcconnection "Http-gateway/pkg/utils/grpc-connection"
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
	"net/http"
)

// CreateInfo godoc
// @Summary Создание информации
// @Description Возвращает строку и айди
// @Tags info
// @Accept json
// @Produce json
// @Security JWTBearer
// @Param smtrequest body info.CreateInfoRequest true "Запрос создания информации"
// @Success 200 {object} info.CreateInfoResponse
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /info [post]
func CreateInfo(c echo.Context) error {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")
	con, err := grpcconnection.ConnectGrpcService(cfg.GRPC.InfoPort)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to connect to gRPC service")
	}
	infoClient := info.NewInfoClient(con)
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
	}
	req := new(info.CreateInfoRequest)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + authHeader,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := infoClient.CreateInfo(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
