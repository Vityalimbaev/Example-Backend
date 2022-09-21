package tests

import (
	"encoding/json"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetContentStateResponseView struct {
	IsSuccess bool                  `json:"is_success,omitempty"`
	Data      []entity.ContentState `json:"data"`
}

type CreateContentStateResponseView struct {
	IsSuccess bool `json:"is_success,omitempty"`
	Data      int  `json:"data"`
}

func (suite *APITestSuite) TestCreateContentState() {
	//SETUP
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData:   entity.ContentState{Title: "someTitle1"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      8,
			},
		},
		{
			ExceptedError: false,
			RequestData:   entity.ContentState{Title: "someTitle2"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      float64(9),
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/create_content_state", testCase.RequestData, suite.PrivateRequest)

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

func (suite *APITestSuite) TestDeleteContentState() {
	//SETUP
	r := suite.Require()

	id := suite.saveContentStates(entity.ContentState{
		Title: "testTitle3",
	})

	testCases := []TestCase{
		{
			RequestMethod: http.MethodDelete,
			RequestURL:    "/delete_content_state",
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

func (suite *APITestSuite) TestGetContentStates() {
	//SETUP

	r := suite.Require()
	testCases := []TestCase{
		{
			RequestMethod: http.MethodDelete,
			RequestURL:    "/get_content_states",
			ExceptedError: false,
			RequestData:   nil,
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      []entity.ContentState{},
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

func (suite *APITestSuite) TestUpdateContentStates() {
	//SETUP
	r := suite.Require()

	cases := []TestCase{
		{
			RequestMethod: http.MethodPost,
			RequestURL:    "/update_content_state",
			ExceptedError: false,
			RequestData: entity.ContentState{
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

			for _, state := range suite.getContentStates() {
				if state == testCase.RequestData {
					continue
				}
			}
		}
	}

}

func (suite *APITestSuite) getContentStates() []entity.ContentState {
	respBytes, statusCode, err := suite.getRequest("/get_content_states", nil, suite.PrivateRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := GetContentStateResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}

func (suite *APITestSuite) saveContentStates(contentState entity.ContentState) int {
	respBytes, statusCode, err := suite.postRequest("/create_content_state", contentState, suite.PrivateRequest)

	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := CreateContentStateResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}
