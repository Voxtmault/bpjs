package logger

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"
)

type RequestLogger struct {
	MariaDB *sql.DB
	RedisDB *redis.Client
}

var requestLogger *RequestLogger

func InitRequestLogger() *RequestLogger {
	requestLogger = &RequestLogger{
		MariaDB: storage.GetDBConnection(),
		RedisDB: storage.GetRedisCon(),
	}
	return requestLogger
}

func GetRequestLogger() *RequestLogger {
	return requestLogger
}

// LogEggressRequest logs the request that is sent to the external service
func (l *RequestLogger) LogEggressRequest(ctx context.Context, request *http.Request, response *http.Response) (err error) {
	if l.MariaDB == nil {
		// this shold not happen
		slog.Warn("mariadb connection is nil")
		l.MariaDB = storage.GetDBConnection()
	}

	tx, err := l.MariaDB.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return eris.Wrap(err, "failed to begin transaction")
	}

	var requestBody []byte
	if request.Body != nil {
		requestBody, err = io.ReadAll(request.Body)
		if err != nil {
			tx.Rollback()
			return eris.Wrap(err, "failed to read request body")
		}
	}
	if len(requestBody) == 0 {
		requestBody = []byte("{}")
	}
	var responseBody []byte
	if response.Body != nil {
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			tx.Rollback()
			return eris.Wrap(err, "failed to read response body")
		}
	}
	if len(responseBody) == 0 {
		responseBody = []byte("{}")
	}

	// Before marshalling the request header, remove sensitive data first
	request.Header.Del("User_key")
	request.Header.Del("X-Cons-Id")

	requestHeader, err := json.Marshal(request.Header)
	if err != nil {
		tx.Rollback()
		return eris.Wrap(err, "failed to marshal request header")
	}

	statement := `
	INSERT INTO bpjs_egress_log (request_url, request_http_method, request_body, request_header,
								 response_status_code, response_body, created_by)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	if _, err = tx.ExecContext(ctx, statement, request.URL.String(), request.Method, string(requestBody), string(requestHeader), response.StatusCode, responseBody, "0"); err != nil {
		tx.Rollback()
		return eris.Wrap(err, "failed to insert data")
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return eris.Wrap(err, "failed to commit transaction")
	}

	return
}
