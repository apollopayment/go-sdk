package requester

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apollopayment/go-sdk/types/responses"
	"io"
	"net/http"
)

func (r *Requester) Request(ctx context.Context, method string, payload interface{}, response interface{}) responses.BaseResponse {
	var resStruct responses.BaseResponse

	mapPayload, err := convertStructToMap(payload)

	if err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on convert request payload: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	mapPayload["nonce"] = r.noncer.GetNonce()

	sign, body, err := r.MakeSignFromStruct(mapPayload)

	if err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on generate sign: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	fmt.Println("body", string(body))

	url := fmt.Sprintf("https://app.apollopayment.io/api-gateway/%s", method)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))

	if err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on make request: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "Apollopayment_SDK")

	req.Header.Set("x-api-public-key", r.publicKey)
	req.Header.Set("x-api-signature", sign)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on send request: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)

	if err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on read response body: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	if resBytes == nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "INTERNAL_ERROR",
			Message: "Receive empty response",
			Code:    -1,
		}

		return resStruct
	}

	fmt.Println(string(resBytes))

	if err = json.Unmarshal(resBytes, &resStruct); err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on parse response body: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	if !resStruct.Success {
		if resStruct.Error == nil {
			resStruct.Success = false
			resStruct.Error = &responses.BaseResponseError{
				Name:    "INTERNAL_ERROR",
				Message: "Internal Server Error",
				Code:    -1,
			}

			return resStruct
		}

		return resStruct
	}

	if err = convertAnyToStruct(resStruct.Response, response); err != nil {
		resStruct.Success = false
		resStruct.Error = &responses.BaseResponseError{
			Name:    "SDK_ERROR",
			Message: "Error on parse response body: " + err.Error(),
			Code:    -1,
		}

		return resStruct
	}

	resStruct.Response = response

	return resStruct
}

func convertStructToMap(payload interface{}) (map[string]any, error) {
	/*
		плохой ход, но он самый простой...
	*/

	bytesPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	var mapPayload map[string]any

	if err = json.Unmarshal(bytesPayload, &mapPayload); err != nil {
		return nil, err
	}

	return mapPayload, nil
}

func convertAnyToStruct(payload interface{}, response interface{}) error {
	/*
		плохой ход, но он самый простой...
	*/

	bytesPayload, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytesPayload, response); err != nil {
		return err
	}

	return nil
}
