package utils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestStatusCodeToHTTPStatus(t *testing.T) {
	tests := []struct {
		code plugnmeet.StatusCode
		want int
	}{
		{plugnmeet.StatusCode_UNKNOWN_STATUS, fiber.StatusInternalServerError},
		{plugnmeet.StatusCode_SUCCESS, fiber.StatusOK},
		{plugnmeet.StatusCode_PERMISSION_DENIED, fiber.StatusForbidden},
		{plugnmeet.StatusCode_NOT_FOUND, fiber.StatusNotFound},
		{plugnmeet.StatusCode_CONFLICT, fiber.StatusConflict},
		{plugnmeet.StatusCode_INVALID_TOKEN_OR_SIGNATURE, fiber.StatusUnauthorized},
		// invalid credentials share the 401 category so clients can renew them
		{plugnmeet.StatusCode_INVALID_API_KEY, fiber.StatusUnauthorized},
		{plugnmeet.StatusCode_INTERNAL_SERVER_ERROR, fiber.StatusInternalServerError},
		{plugnmeet.StatusCode_INVALID_PARAMETERS, fiber.StatusBadRequest},
		{plugnmeet.StatusCode_MISSING_REQUIRED_PARAMETER, fiber.StatusBadRequest},
		{plugnmeet.StatusCode_ROOM_NOT_FOUND, fiber.StatusNotFound},
		{plugnmeet.StatusCode_USER_NOT_FOUND, fiber.StatusNotFound},
		{plugnmeet.StatusCode_USER_BLOCKED, fiber.StatusForbidden},
		{plugnmeet.StatusCode_USER_NOT_ACTIVE, fiber.StatusForbidden},
		{plugnmeet.StatusCode_ROOM_NOT_ACTIVE, fiber.StatusConflict},
		// unknown/future enum values must not silently become 2xx
		{plugnmeet.StatusCode(999), fiber.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.code.String(), func(t *testing.T) {
			if got := StatusCodeToHTTPStatus(tt.code); got != tt.want {
				t.Errorf("StatusCodeToHTTPStatus(%v) = %d, want %d", tt.code, got, tt.want)
			}
			if status := CommonErrorHTTPStatus(tt.code); status < 400 || status > 599 {
				t.Errorf("CommonErrorHTTPStatus(%v) = %d, want a non-2xx status", tt.code, status)
			}
		})
	}
}

func TestSendCommonErrorResponses(t *testing.T) {
	app := fiber.New()
	app.Get("/protobuf", func(c fiber.Ctx) error {
		return SendCommonErrorProtobufResponse(c, "notifications.room-not-active", plugnmeet.StatusCode_ROOM_NOT_ACTIVE)
	})
	app.Get("/json", func(c fiber.Ctx) error {
		return SendCommonErrorJsonResponse(c, "notifications.room-not-found", plugnmeet.StatusCode_ROOM_NOT_FOUND)
	})
	app.Get("/success-code", func(c fiber.Ctx) error {
		// a success code must never produce a 2xx error response
		return SendCommonErrorProtobufResponse(c, "notifications.error-occurred", plugnmeet.StatusCode_SUCCESS)
	})

	t.Run("protobuf", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/protobuf", nil))
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		if res.StatusCode != fiber.StatusConflict {
			t.Errorf("status = %d, want %d", res.StatusCode, fiber.StatusConflict)
		}
		if ct := res.Header.Get("Content-Type"); ct != "application/protobuf" {
			t.Errorf("content type = %q, want application/protobuf", ct)
		}

		var body plugnmeet.CommonResponse
		if err := proto.Unmarshal(readBody(t, res), &body); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if body.Status {
			t.Error("status = true, want false")
		}
		if body.Msg != "notifications.room-not-active" {
			t.Errorf("msg = %q", body.Msg)
		}
		if body.StatusCode != plugnmeet.StatusCode_ROOM_NOT_ACTIVE {
			t.Errorf("status code = %v", body.StatusCode)
		}
	})

	t.Run("json", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/json", nil))
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		if res.StatusCode != fiber.StatusNotFound {
			t.Errorf("status = %d, want %d", res.StatusCode, fiber.StatusNotFound)
		}
		if ct := res.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("content type = %q, want application/json", ct)
		}

		var body plugnmeet.CommonResponse
		if err := protojson.Unmarshal(readBody(t, res), &body); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if body.Status {
			t.Error("status = true, want false")
		}
		if body.Msg != "notifications.room-not-found" {
			t.Errorf("msg = %q", body.Msg)
		}
		if body.StatusCode != plugnmeet.StatusCode_ROOM_NOT_FOUND {
			t.Errorf("status code = %v", body.StatusCode)
		}
	})

	t.Run("success code is not a 2xx", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/success-code", nil))
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		if res.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("status = %d, want %d", res.StatusCode, fiber.StatusInternalServerError)
		}
	})
}

func readBody(t *testing.T, res *http.Response) []byte {
	t.Helper()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	return body
}
