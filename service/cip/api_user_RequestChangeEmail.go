package cip

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
)

/*
RequestChangeEmail
An email with an activation link is sent to the user’s new email address. The user must click the link in the email within seven days to complete the email address change, or the link will expire.

	body - New email address of the user.
	id - Identifier of the user to change email address.
*/
func (a *APIClient) RequestChangeEmail(body types.ChangeEmailRequest, id string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/users/{id}/email/requestChange"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode >= 400 {
			var v types.ErrorResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			if len(v.Errors) > 0 {
				if v.Errors[0].Meta.Reason != "" {
					newErr.error = v.Errors[0].Message + ": " + v.Errors[0].Meta.Reason
				} else {
					newErr.error = v.Errors[0].Message
				}
			}
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
