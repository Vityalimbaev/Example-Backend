package tests

import (
	"encoding/json"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetContentActionResponseView struct {
	IsSuccess bool                   `json:"is_success,omitempty"`
	Data      []entity.ContentAction `json:"data"`
}

type CreateContentActionResponseView struct {
	IsSuccess bool `json:"is_success,omitempty"`
	Data      int  `json:"data"`
}

func (suite *APITestSuite) TestCreateContentAction() {
	//SETUP
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData:   entity.ContentAction{Title: "someTitle1"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      8,
			},
		},
		{
			ExceptedError: false,
			RequestData:   entity.ContentAction{Title: "someTitle2"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      float64(9),
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/create_content_action", testCase.RequestData, suite.PrivateRequest)

		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		actualResponse := entity.ResponseView{
			IsSuccess: false,
			Data:      0,
		}
		err = json.Unmarshal(respBytes, &actualResponse)

		r.NoError(err)
		r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		r.Condition(func() (success bool) {
			return actualResponse.Data.(float64) > 0
		})

	}
}

func (suite *APITestSuite) TestDeleteContentAction() {
	//SETUP
	r := suite.Require()

	id := suite.saveContentActions(entity.ContentAction{
		Title: "testTitle3",
	})

	testCases := []TestCase{
		{
			RequestMethod: http.MethodDelete,
			RequestURL:    "/delete_content_action",
			ExceptedError: false,
			RequestData:   id,
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      struct{}{},
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)
		respBytes, statusCode, err := suite.deleteByIdRequest(testCase.RequestURL, id, suite.PrivateRequest)

		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		actualResponse := entity.ResponseView{}
		err = json.Unmarshal(respBytes, &actualResponse)

		r.NoError(err)
		r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
	}
}

func (suite *APITestSuite) TestGetContentActions() {
	//SETUP

	r := suite.Require()
	testCases := []TestCase{
		{
			RequestMethod: http.MethodDelete,
			RequestURL:    "/get_content_actions",
			ExceptedError: false,
			RequestData:   nil,
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      []entity.ContentAction{},
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)
		respBytes, statusCode, err := suite.getRequest(testCase.RequestURL, nil, suite.PrivateRequest)

		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		actualResponse := entity.ResponseView{}
		err = json.Unmarshal(respBytes, &actualResponse)

		r.NoError(err)
		r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		r.NotEmpty(actualResponse.Data)
	}
}

func (suite *APITestSuite) TestUpdateContentAction() {
	//SETUP
	r := suite.Require()

	cases := []TestCase{
		{
			RequestMethod: http.MethodPost,
			RequestURL:    "/update_content_action",
			ExceptedError: false,
			RequestData: entity.ContentAction{
				Id:    1,
				Title: "NewTitle",
			},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      struct{}{},
			},
		},
	}

	//REQUEST
	for i, testCase := range cases {
		logrus.Infof("Testing case %d : %v ", i, testCase)
		respBytes, statusCode, err := suite.postRequest(testCase.RequestURL, testCase.RequestData, suite.PrivateRequest)

		if !testCase.ExceptedError {
			r.NoError(err)
			r.Equal(http.StatusOK, statusCode)

			actualResponse := entity.ResponseView{}
			err = json.Unmarshal(respBytes, &actualResponse)

			r.NoError(err)
			r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
			r.NotNil(actualResponse.Data)

			for _, action := range suite.getContentActions() {
				if action == testCase.RequestData {
					continue
				}
			}
		}
	}

}

func (suite *APITestSuite) getContentActions() []entity.ContentAction {
	respBytes, statusCode, err := suite.getRequest("/get_content_actions", nil, suite.PrivateRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := GetContentActionResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}

func (suite *APITestSuite) saveContentActions(contentAction entity.ContentAction) int {
	respBytes, statusCode, err := suite.postRequest("/create_content_action", contentAction, suite.PrivateRequest)

	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := CreateContentActionResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}
