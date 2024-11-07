package authservice

import (
	"Http-gateway/internal/config"
	"Http-gateway/internal/data/gen/auth"
	"Http-gateway/pkg/singleton"
	grpcconnection "Http-gateway/pkg/utils/grpc-connection"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Login godoc
// @Summary Вход
// @Description Авторизация пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param req body auth.LoginRequest true "Login data"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /login [post]
func Login(c echo.Context) error {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")

	con, err := grpcconnection.ConnectGrpcService(cfg.GRPC.AuthPort)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to connect to gRPC service")
	}
	//⣿⠟⢋⣉⢙⣿⣿⡉⢉⣙⣿⠏⢉⠉⢹⣿⡿⠋⠙⣿⣿⡏⠙⠋⣙⣿⡿⠋⠙⣿
	//⣿⡀⠘⣋⣽⣿⡟⠀⣼⣿⠏⢠⣶⣶⣿⠏⣠⣬⠀⣿⠟⢁⡄⢸⣿⠋⢠⡄⠀⣿
	//⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠹⠏⢉⣿⠏⣉⣉⣹⣿⣏⠉⣉⣹⣿⣿⣿⣿⣿⣿⣿
	//⣿⣿⣿⣿⣿⣿⣿⣿⠏⣰⡆⢠⣿⠏⠐⠒⢒⣾⣿⠃⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿
	//⣿⣿⣿⡿⠋⣉⣉⣹⣿⠏⣉⠉⣹⣿⡿⠋⠉⣿⣿⣍⠉⣉⣽⡿⠉⠽⢿⣿⣿⣿
	//⣿⣿⣿⣧⡀⢉⣩⣿⠋⣰⣶⣾⣿⢋⣤⣅⢀⣿⣿⡃⣸⣿⣿⠁⠒⢀⣼⣿⣿⣿
	//⣿⣿⣿⡏⠉⠭⢹⣿⡿⠉⠭⢭⣿⡟⢩⡟⠉⣽⡟⠉⠽⠋⣹⡿⠋⣩⡍⢹⣿⣿
	//⣿⣿⡟⠀⠒⠀⣼⣟⠁⠀⠒⣾⣿⣷⣶⢀⣼⣿⢁⣴⡀⣼⣿⣇⡈⢉⣠⣾⣿⣿
	//⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛
	//⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⠖⡒⠒⠖⣋⡀⡄⠀⠀⠀⠀⠀⠀⠀⠀⣠⠄⡀
	//⠀⠀⠀⠀⣀⣤⣶⣿⣿⣿⡀⠛⠀⠀⠱⠛⠀⠽⠀⠀⠀⠀⢀⡤⠚⠁⠔⠁⠀⠀
	//⠀⠀⣰⣿⣿⠉⡉⣿⣿⣽⣮⣥⣾⣿⣮⣷⡾⠀⠀⠀⢠⠎⠀⠀⠀⡑⢄⠀⠀
	//⢀⠎⠸⣿⣿⣦⠑⠑⠒⠒⠒⠒⠒⠒⡉⠀⠀⠀⠀⠀⣇⠀⣈⠍⡲⣀⡄⠇⠀
	//⣿⣶⣤⡈⠛⠿⠿⢿⣿⠿⠿⢿⣛⡯⣲⣄⠀⠀⠀⠀⠈⠓⠮⣔⠒⢏⠀⠀⠀
	//⣿⣿⣿⣿⣷⣶⣄⠸⠟⠛⠛⠻⢿⣷⣿⣿⣷⡀⠀⠀⠀⠀⠀⠀⠈⠛⠚⠀⠀
	authClient := auth.NewAuthClient(con)
	req := new(auth.LoginRequest)

	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	//⠀⠀⠀⢠⠤⠤⠀⡤⠤⠄⠀⡤⢤⠀⢀⣤⠀⢠⡄⠀⡄⠀⠀⠀
	//⠀⠀⠀⢸⠒⠒⠀⡗⠒⣦⠀⡇⢸⠀⡼⠼⣆⢸⡗⠒⡇⠀⠀⠀
	//⠀⠀⠀⠘⠛⠛⠀⠛⠛⠁⠘⠁⠘⠀⠁⠀⠘⠈⠁⠀⠃⠀⠀
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	res, err := authClient.Login(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	//: ⠄⠄⠄⠄ ⠄⠄⠄⠄ ⠄⠄⠄⠄
	// ⠄⠄⡔⠙⠢⡀⠄⠄⠄⢀⠼⠅⠈⢂⠄⠄⠄⠄
	// ⠄⠄⡌⠄⢰⠉⢙⢗⣲⡖⡋⢐⡺⡄⠈⢆⠄⠄⠄
	// ⠄⡜⠄⢀⠆⢠⣿⣿⣿⣿⢡⢣⢿⡱⡀⠈⠆⠄⠄
	// ⠄⠧⠤⠂⠄⣼⢧⢻⣿⣿⣞⢸⣮⠳⣕⢤⡆⠄⠄
	// ⢺⣿⣿⣶⣦⡇⡌⣰⣍⠚⢿⠄⢩⣧⠉⢷⡇⠄⠄
	// ⠘⣿⣿⣯⡙⣧⢎⢨⣶⣶⣶⣶⢸⣼⡻⡎⡇⠄⠄
	// ⠄⠘⣿⣿⣷⡀⠎⡮⡙⠶⠟⣫⣶⠛⠧⠁⠄⠄⠄
	// ⠄⠄⠘⣿⣿⣿⣦⣤⡀⢿⣿⣿⣿⣄⠄⠄⠄⠄⠄
	// ⠄⠄⠄⠈⢿⣿⣿⣿⣿⣷⣯⣿⣿⣷⣾⣿⣷⡄⠄
	// ⠄⠄⠄⠄⠄⢻⠏⣼⣿⣿⣿⣿⡿⣿⣿⣏⢾⠇⠄
	// ⠄⠄⠄⠄⠄⠈⡼⠿⠿⢿⣿⣦⡝⣿⣿⣿⠷⢀⠄
	// ⠄⠄⠄⠄⠄⠄⡇⠄⠄⠄⠈⠻⠇⠿⠋⠄⠄⢘⡆
	// ⠄⠄⠄⠄⠄⠄⠱⣀⠄⠄⠄⣀⢼⡀⠄⢀⣀⡜⠄
	// ⠄⠄⠄⠄⠄⠄⠄⢸⣉⠉⠉⠄⢀⠈⠉⢏⠁⠄⠄
	// ⠄⠄⠄⠄⠄⠄⡰⠃⠄⠄⠄⠄⢸⠄⠄⢸⣧⠄⠄
	// ⠄⠄⠄⠄⠄⣼⣧⠄⠄⠄⠄⠄⣼⠄⠄⡘⣿⡆⠄
	// ⠄⠄⠄⢀⣼⣿⡙⣷⡄⠄⠄⠄⠃⠄⢠⣿⢸⣿⡀
	// ⠄⠄⢀⣾⣿⣿⣷⣝⠿⡀⠄⠄⠄⢀⡞⢍⣼⣿⠇
	// ⠄⠄⣼⣿⣿⣿⣿⣿⣷⣄⠄⠄⠠⡊⠴⠋⠹⡜⠄
	// ⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⡆⣤⣾⣿⣿⣧⠹⠄⠄
	// ⠄⠄⢿⣿⣿⣿⣿⣿⣿⣿⢃⣿⣿⣿⣿⣿⡇⠄⠄
	// ⠄⠄⠐⡏⠉⠉⠉⠉⠉⠄⢸⠛⠿⣿⣿⡟⠄⠄⠄
	// ⠄⠄⠄⠹⡖⠒⠒⠒⠒⠊⢹⠒⠤⢤⡜⠁⠄⠄⠄
	// ⠄⠄⠄⠄⠱⠄⠄⠄⠄⠄⢸⠄⠄⠄⡖⠄⠄⠄⠄
	return c.JSON(http.StatusOK, res)
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя с указанными данными
// @Tags auth
// @Accept json
// @Produce json
// @Param req body auth.RegisterRequest true "Register data"
// @Success 200 {object} auth.RegisterResponse
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /register [post]
func Register(c echo.Context) error {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")
	c.Logger().Infof("Loaded config: %+v", cfg)

	con, err := grpcconnection.ConnectGrpcService(cfg.GRPC.AuthPort)

	if err != nil {
		c.Logger().Error("Failed to connect to gRPC service: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to connect to gRPC service")
	}
	c.Logger().Info("Successfully connected to gRPC service")

	authClient := auth.NewAuthClient(con)

	req := new(auth.RegisterRequest)

	if err = c.Bind(req); err != nil {
		c.Logger().Error("Failed to bind request: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	c.Logger().Infof("Request bound successfully: %+v", req)

	res, err := authClient.Register(context.Background(), req)

	if err != nil {
		c.Logger().Error("gRPC Register call failed: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Infof("gRPC Register response: %+v", res)

	return c.JSON(http.StatusOK, res)
}

// GetMe godoc
// @Summary Получить информацию о пользователе
// @Description Получение данных текущего пользователя по токену аутентификации
// @Tags auth
// @Accept json
// @Produce json
// @Security JWTBearer
// @Success 200 {object} auth.GetMeResponse
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /me [get]
func GetMe(c echo.Context) error {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")

	con, err := grpcconnection.ConnectGrpcService(cfg.GRPC.AuthPort)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to connect to gRPC service")
	}
	authClient := auth.NewAuthClient(con)
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
	}

	req := &auth.GetMeRequest{
		Auth_JWT_Header: authHeader,
	}
	res, err := authClient.GetMe(context.Background(), req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// Validate godoc
// @Summary Валидация токена
// @Description Проверка валидности токена JWT
// @Tags auth
// @Accept json
// @Produce json
// @Security JWTBearer
// @Success 200 {object} auth.ValidateResponse
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /validate [post]
func Validate(c echo.Context) error {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")

	con, err := grpcconnection.ConnectGrpcService(cfg.GRPC.AuthPort)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to connect to gRPC service")
	}
	authClient := auth.NewAuthClient(con)

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
	}

	req := &auth.ValidateRequest{
		Auth_JWT_Header: authHeader,
	}
	res, err := authClient.Validate(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// Refresh godoc
// @Summary Обновление токена
// @Description Обновляет токены доступа и рефреш токены
// @Tags auth
// @Accept json
// @Produce json
// @Security JWTBearer
// @Success 200 {object} auth.RefreshResponse
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /refresh [post]
func Refresh(c echo.Context) error {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")
	con, err := grpcconnection.ConnectGrpcService(cfg.GRPC.AuthPort)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to connect to gRPC service")
	}
	authClient := auth.NewAuthClient(con)
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
	}

	req := &auth.RefreshRequest{
		RefreshToken: authHeader,
	}
	res, err := authClient.Refresh(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
