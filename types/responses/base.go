package responses

type BaseResponseError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type BaseResponse struct {
	Success   bool               `json:"success"`
	Response  any                `json:"response"`
	Error     *BaseResponseError `json:"error"`
	RequestId string             `json:"requestId"`
}

type BaseResponseGeneric[T any] struct {
	Success   bool               `json:"success"`
	Error     *BaseResponseError `json:"error"`
	RequestId string             `json:"requestId"`
	Response  T                  `json:"response"`
}

func ConvertBase[T interface{}](res BaseResponse) BaseResponseGeneric[T] {
	if !res.Success {
		return BaseResponseGeneric[T]{
			Success:   false,
			Error:     res.Error,
			RequestId: res.RequestId,
		}
	}

	if res.Response == nil {
		return BaseResponseGeneric[T]{
			Success: true,
		}
	}

	_res, ok := res.Response.(T)

	if !ok {
		return BaseResponseGeneric[T]{
			Success: true,
			Error: &BaseResponseError{
				Name:    "SDK_ERROR",
				Message: "Cast response type error",
				Code:    -1,
			},
		}
	}

	return BaseResponseGeneric[T]{
		Success:  true,
		Response: _res,
	}
}
