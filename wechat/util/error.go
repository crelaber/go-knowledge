package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type CommonErr struct {
	apiName string
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func (c *CommonErr) Error() string {
	return fmt.Sprintf("%s Error, errcode=%d, errmsg=%s", c.apiName, c.ErrCode, c.ErrMsg)
}

func DecodeWithCommonError(response []byte, apiName string) (err error) {
	var commonErr CommonErr
	err = json.Unmarshal(response, &commonErr)
	if err != nil {
		return
	}

	commonErr.apiName = apiName
	if commonErr.ErrCode != 0 {
		return &commonErr
	}

	return nil
}

func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	if err != nil {
		return fmt.Errorf("json unmarshal error, err = %v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		fmt.Errorf("obj is invalid")
	}

	commonErr := responseObj.Elem().FieldByName("CommonError")
	if !commonErr.IsValid() || commonErr.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is valid or not struct")
	}
	errCode := commonErr.FieldByName("ErrCode")
	errMsg := commonErr.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is valid")
	}
	if errCode.Int() != 0 {
		return &CommonErr{
			apiName: apiName,
			ErrCode: errCode.Int(),
			ErrMsg:  errMsg.String(),
		}
	}
	return nil
}
