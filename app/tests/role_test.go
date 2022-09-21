package tests

import (
	"encoding/json"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetRoleResponseView struct {
	IsSuccess bool          `json:"is_success,omitempty"`
	Data      []entity.Role `json:"data"`
}

type CreateRoleResponseView struct {
	IsSuccess bool `json:"is_success,omitempty"`
	Data      int  `json:"data"`
}

func (suite *APITestSuite) TestCreateRole() {
	//SETUP
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData:   entity.Role{Title: "someTitle1"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      8,
			},
		},
		{
			ExceptedError: false,
			RequestData:   entity.Role{Title: "someTitle2"},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      float64(9),
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/create_role", testCase.RequestData, suite.PrivateRequest)

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

func (suite *APITestSuite) TestDeleteRole() {
	//SETUP
	r := suite.Require()

	id := suite.saveRoles(entity.Role{
		Title: "testTitle3",
	})

	testCases := []TestCase{
		{
			RequestMethod: http.MethodDelete,
			RequestURL:    "/delete_role",
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

func (suite *APITestSuite) TestGetRoles() {
	//SETUP

	r := suite.Require()
	testCases := []TestCase{
		{
			RequestMethod: http.MethodDelete,
			RequestURL:    "/get_roles",
			ExceptedError: false,
			RequestData:   nil,
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      []entity.Role{},
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

func (suite *APITestSuite) TestUpdateRoles() {
	//SETUP
	r := suite.Require()

	cases := []TestCase{
		{
			RequestMethod: http.MethodPost,
			RequestURL:    "/update_role",
			ExceptedError: false,
			RequestData: entity.Role{
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

			for _, state := range suite.getRoles() {
				if state == testCase.RequestData {
					continue
				}
			}
		}
	}

}

func (suite *APITestSuite) getRoles() []entity.Role {
	respBytes, statusCode, err := suite.getRequest("/get_roles", nil, suite.PrivateRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := GetRoleResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}

func (suite *APITestSuite) saveRoles(Role entity.Role) int {
	respBytes, statusCode, err := suite.postRequest("/create_role", Role, suite.PrivateRequest)

	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := CreateRoleResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}
