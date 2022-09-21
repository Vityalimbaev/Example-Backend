package tests

import (
	"encoding/json"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type TokenSuccessResponse struct {
	IsSuccess bool                      `json:"is_success"`
	Data      entity.TokensResponseView `json:"data"`
}

func (suite *APITestSuite) TestLogin() {
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData:   entity.LoginForm{Email: "admin@sadkomed.ru", Password: "123456"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
			},
		},
		{
			ExceptedError: true,
			RequestData:   entity.LoginForm{Email: "admin@sadkomed.ru", Password: "12345"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: false,
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/login", testCase.RequestData, suite.PublicRequest)
		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		if !testCase.ExceptedError {
			actualResponse := TokenSuccessResponse{}
			err = json.Unmarshal(respBytes, &actualResponse)
			r.NoError(err)
			req, _ := http.NewRequest(http.MethodGet, "/api/isauth", nil)
			req.Header.Set("Content-type", "application/json")
			req.Header.Set("Authorization", "Bearer "+actualResponse.Data.AccessToken)

			resp := httptest.NewRecorder()
			suite.handlers.ServeHTTP(resp, req)
			respBytes, err = ioutil.ReadAll(resp.Body)

			err = json.Unmarshal(respBytes, &actualResponse)
			r.NoError(err)
			r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		} else {
			actualResponse := entity.ResponseView{}
			err = json.Unmarshal(respBytes, &actualResponse)
			r.NoError(err)
			r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		}
	}
}

func (suite *APITestSuite) TestSaveUser() {
	//SETUP
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData: entity.User{
				Name:         "sssadqqqqqqqqq",
				Username:     "sssadqqqqqqqqqqqq",
				Email:        "some@email.com",
				Password:     "qwerty",
				Branch:       "qwerty",
				RoleId:       1,
				RoleTitle:    "qwerty",
				ActiveStatus: true,
			},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
			},
		},
		{
			ExceptedError: true,
			RequestData: entity.User{
				Name:         "sfssssadqqqqqqqqq",
				Username:     "sfsssadqqqqqqqqqqqq",
				Email:        "some@email.com",
				Password:     "qwerty",
				Branch:       "qwerty",
				RoleId:       1,
				RoleTitle:    "qwerty",
				ActiveStatus: true,
			},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: false,
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/save-user", testCase.RequestData, suite.PrivateRequest)

		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		if !testCase.ExceptedError {
			actualResponse := struct {
				IsSuccess bool        `json:"is_success"`
				Data      entity.User `json:"data"`
			}{}
			err = json.Unmarshal(respBytes, &actualResponse)

			r.NoError(err)
			r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
			r.Condition(func() (success bool) {
				return actualResponse.Data.Id > 0
			})
		} else {
			actualResponse := entity.ResponseView{}
			err = json.Unmarshal(respBytes, &actualResponse)

			r.NoError(err)
			r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		}

	}
}

func (suite *APITestSuite) TestRefreshToken() {
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData:   entity.LoginForm{Email: "admin@sadkomed.ru", Password: "123456"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
			},
		},
	}
	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/login", testCase.RequestData, suite.PublicRequest)
		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		if !testCase.ExceptedError {
			oldResponse := TokenSuccessResponse{}
			err = json.Unmarshal(respBytes, &oldResponse)
			r.NoError(err)
			_, statusCode, err := suite.getRequest("/refresh-token", oldResponse.Data, suite.PublicRequest)
			r.NoError(err)
			r.Equal(http.StatusOK, statusCode)

			newResponse := struct {
				IsSuccess bool                      `json:"is_success"`
				Data      entity.TokensResponseView `json:"data"`
			}{}
			newBytes, statusCode, err := suite.postRequest("/login", testCase.RequestData, suite.PublicRequest)
			r.NoError(err)
			r.Equal(http.StatusOK, statusCode)
			err = json.Unmarshal(newBytes, &newResponse)

			r.NoError(err)
			r.NotEqual(oldResponse.Data.RefreshToken, newResponse.Data.RefreshToken)
			r.NotEqual(oldResponse.Data.AccessToken, newResponse.Data.AccessToken)
		} else {
			actualResponse := entity.ResponseView{}
			err = json.Unmarshal(respBytes, &actualResponse)
			r.NoError(err)
			r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		}
	}
}
