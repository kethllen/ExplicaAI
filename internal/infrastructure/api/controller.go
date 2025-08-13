package api

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kethllen/explicaAi/internal/infrastructure/log"
	"github.com/labstack/echo/v4"
)

type ExplicaServer struct {
}

func NewExplicaServer() *ExplicaServer {
	return &ExplicaServer{}
}

func (api *ExplicaServer) Register(server *echo.Echo) {
	server.POST("/upload", api.Upload)
	server.GET("/summaries", api.ListSummaries)
	server.GET("/summaries/:externalId", api.GetSummaryByExternalId)
	server.DELETE("/summaries/:externalId", api.DeleteSummaryByExternalId)
}

func (api *ExplicaServer) ListSummaries(c echo.Context) error {

	//ctx := c.Request().Context()

	//todo: init get flow
	return c.JSON(http.StatusOK, nil)
}

func (api *ExplicaServer) GetSummaryByExternalId(c echo.Context) error {

	//ctx := c.Request().Context()
	//externalID := c.Param("externalId")
	// parsedExternalID, err := uuid.Parse(externalID)
	// if err != nil {
	// 	return echo.ErrBadRequest
	// }
	//todo: init get flow
	return c.JSON(http.StatusOK, nil)
}

func (api *ExplicaServer) DeleteSummaryByExternalId(c echo.Context) error {

	//ctx := c.Request().Context()
	//externalID := c.Param("externalId")
	// parsedExternalID, err := uuid.Parse(externalID)
	// if err != nil {
	// 	return echo.ErrBadRequest
	// }
	//todo: init delete flow
	return c.JSON(http.StatusOK, map[string]string{
		"message": "summary has been removed",
	})
}

func (api *ExplicaServer) Upload(c echo.Context) error {

	ctx := c.Request().Context()
	_, err := api.getFileFromRequest(ctx, c)
	if err != nil {
		return echo.ErrBadRequest
	}
	//todo: init flow
	return c.JSON(http.StatusCreated, nil)
}

func (api *ExplicaServer) getFileFromRequest(ctx context.Context, c echo.Context) ([]byte, error) {
	file, err := c.FormFile("file")
	if err != nil {
		log.LogError(ctx, "missing file", err)
		return nil, errors.New("missing file")
	}
	allowedWxtensions := map[string]bool{
		".mp3":  true,
		".mp4":  true,
		".mpeg": true,
		".mpga": true,
		".m4a":  true,
		".wav":  true,
		".webm": true,
	}

	fileEstension := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedWxtensions[fileEstension] {
		return nil, errors.New("invalid file")
	}
	src, err := file.Open()
	if err != nil {
		log.LogError(ctx, "fail to open file", err)
		return nil, err
	}
	defer src.Close()

	var buf bytes.Buffer

	if _, err := io.Copy(&buf, src); err != nil {
		log.LogError(ctx, "fail to read file", err)
		return nil, err
	}
	return buf.Bytes(), nil
}
