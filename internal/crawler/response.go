package crawler

import "simplecrawler/internal/utils"

type Response struct {
	URL          string       `json:"url"`
	Code         ResponseCode `json:"code"`
	ErrorMessage *string      `json:"error_message"`
	Object       *string      `json:"object"`
}

type ResponseCode uint16

const (
	InternalErrorCode ResponseCode = 200
	HTTPErrorCode     ResponseCode = 201
	ParsingErrorCode  ResponseCode = 202
	SuccessCode       ResponseCode = 100
)

func ErrorResponse(url string, code ResponseCode, errorMsg string) Response {
	return Response{
		URL:          url,
		Code:         code,
		ErrorMessage: utils.StringPtr(errorMsg),
	}
}

func SuccessResponse(url string, code ResponseCode, obj string) Response {
	return Response{
		URL:    url,
		Code:   code,
		Object: utils.StringPtr(obj),
	}
}
