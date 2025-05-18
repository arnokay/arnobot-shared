package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
  Routes(group *echo.Group)
}
