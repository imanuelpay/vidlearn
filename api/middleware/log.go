package middleware

import (
	"fmt"
	"net/http"
	"time"
	"vidlearn-final-projcect/util"

	logService "vidlearn-final-projcect/business/log"
	logRepository "vidlearn-final-projcect/repository/log"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func makeLogEntry(c echo.Context) *log.Entry {
	return log.WithFields(log.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"ip":         c.Request().RemoteAddr,
		"host":       c.Request().Host,
		"method":     c.Request().Method,
		"path":       c.Request().URL.Path,
		"user_agent": c.Request().UserAgent(),
		"status":     c.Response().Status,
	})
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	log.Println(report.Code, report.Message.(string))
}

func Logger(e *echo.Echo, dbcon *util.DatabaseConnection) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			data := makeLogEntry(c).Data
			makeLogEntry(c).Info()

			logger := &logService.Log{
				At:        data["at"].(string),
				Host:      data["host"].(string),
				Method:    data["method"].(string),
				Path:      data["path"].(string),
				IP:        data["ip"].(string),
				Status:    data["status"].(int),
				UserAgent: data["user_agent"].(string),
			}

			logPremitRepository := logRepository.LogRepository(dbcon)
			_, err := logPremitRepository.CreateLog(logger)
			if err != nil {
				return err
			}

			return next(c)
		}
	})
}
