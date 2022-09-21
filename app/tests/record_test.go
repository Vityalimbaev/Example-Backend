package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type GetRecordResponseView struct {
	IsSuccess bool            `json:"is_success,omitempty"`
	Data      []entity.Record `json:"data"`
}

type CreateRecordResponseView struct {
	IsSuccess bool `json:"is_success,omitempty"`
	Data      int  `json:"data"`
}

func (suite *APITestSuite) TestCreateRecord() {
	//SETUP
	r := suite.Require()

	testCases := []TestCase{
		{
			ExceptedError: false,
			RequestData: entity.Record{
				ArchivedDate:   time.Date(2022, 02, 11, 11, 11, 11, 11, time.UTC).Unix(),
				Branch:         "sadko",
				Pcode:          123456,
				LastTreat:      time.Date(2022, 02, 11, 11, 11, 11, 11, time.UTC).Unix(),
				ContentStateId: 1,
				BoxId:          1,
			},
			ExpectedResponseData: entity.ResponseView{
				IsSuccess: true,
				Data:      0,
			},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)

		respBytes, statusCode, err := suite.postRequest("/create_record", testCase.RequestData, suite.PrivateRequest)

		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		actualResponse := entity.ResponseView{}
		err = json.Unmarshal(respBytes, &actualResponse)

		r.NoError(err)
		r.Equal(testCase.ExpectedResponseData.(entity.ResponseView).IsSuccess, actualResponse.IsSuccess)
		r.Condition(func() (success bool) {
			return actualResponse.Data.(float64) > 0
		})

		time.Sleep(5 * time.Second)

		records := suite.getRecord(entity.RecordSearchParams{
			Id: int(actualResponse.Data.(float64)),
		})

		r.Condition(func() (success bool) {
			fmt.Println(len(records))
			return len(records) > 0
		})

		r.Equal(testCase.RequestData.(entity.Record).Pcode, records[0].Pcode)
		r.Equal(testCase.RequestData.(entity.Record).ContentStateId, records[0].ContentStateId)
		r.Equal(testCase.RequestData.(entity.Record).ArchivedDate, records[0].ArchivedDate)
		r.Equal(testCase.RequestData.(entity.Record).LastTreat, records[0].LastTreat)
		r.Equal(testCase.RequestData.(entity.Record).Branch, records[0].Branch)
		r.Equal(testCase.RequestData.(entity.Record).BoxId, records[0].BoxId)
	}
}

func (suite *APITestSuite) TestDeleteRecord() {
	//SETUP
	r := suite.Require()

	id := suite.saveRecords(entity.Record{
		ArchivedDate:   time.Date(2022, 02, 11, 11, 11, 11, 11, time.UTC).Unix(),
		Branch:         "sadkoTestDelete",
		Pcode:          1234568,
		LastTreat:      time.Date(2022, 02, 11, 11, 11, 11, 11, time.UTC).Unix(),
		ContentStateId: 1,
		BoxId:          1,
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

func (suite *APITestSuite) TestGetRecords() {
	//SETUP

	r := suite.Require()
	testCases := []TestCase{
		{
			RequestMethod:        http.MethodDelete,
			RequestURL:           "/get_records",
			ExceptedError:        false,
			RequestData:          struct{}{},
			ExpectedResponseData: GetRecordResponseView{},
		},
	}

	for i, testCase := range testCases {
		logrus.Infof("Testing case %d : %v ", i, testCase)
		respBytes, statusCode, err := suite.postRequest(testCase.RequestURL, struct{}{}, suite.PrivateRequest)

		r.NoError(err)
		r.Equal(http.StatusOK, statusCode)

		actualResponse := GetRecordResponseView{}
		err = json.Unmarshal(respBytes, &actualResponse)

		r.NoError(err)
		r.NotEmpty(actualResponse.Data)
	}
}

func (suite *APITestSuite) TestUpdateRecords() {
	//SETUP
	r := suite.Require()

	cases := []TestCase{
		{
			RequestMethod: http.MethodPost,
			RequestURL:    "/update_record",
			ExceptedError: false,
			RequestData: entity.Record{
				Id:     1,
				Branch: "NewBranchName",
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

			for _, state := range suite.getRecord(entity.RecordSearchParams{
				Id: 1,
			}) {
				r.True(state.Branch == testCase.RequestData.(entity.Record).Branch)
			}
		}
	}
}

func (suite *APITestSuite) getRecord(recordSearchParams entity.RecordSearchParams) []entity.Record {
	fmt.Printf("Get Record by : %v", recordSearchParams)
	respBytes, statusCode, err := suite.postRequest("/get_records", recordSearchParams, suite.PrivateRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := GetRecordResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}

func (suite *APITestSuite) saveRecords(record entity.Record) int {
	respBytes, statusCode, err := suite.postRequest("/create_record", record, suite.PrivateRequest)

	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, statusCode)

	resp := CreateRecordResponseView{}
	err = json.Unmarshal(respBytes, &resp)

	suite.Require().NoError(err)

	return resp.Data
}
