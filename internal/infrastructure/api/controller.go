package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type ExplicaServer struct {
}

func NewExplicaServer() *ExplicaServer {
	return &ExplicaServer{}
}
func (api *ExplicaServer) Register(server *echo.Echo) {
	server.POST("/upload", api.Upload)
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
		fmt.Println("missing file")
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
		fmt.Println("fail to open file")
		return nil, err
	}
	defer src.Close()

	var buf bytes.Buffer

	if _, err := io.Copy(&buf, src); err != nil {
		fmt.Println("fail to read file")
		return nil, err
	}
	return buf.Bytes(), nil
}
