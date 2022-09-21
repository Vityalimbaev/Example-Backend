package tests

import (
	"bytes"
	"encoding/json"
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/Vityalimbaev/Example-Backend/internal/adapter"
	"github.com/Vityalimbaev/Example-Backend/internal/adapter/database"
	"github.com/Vityalimbaev/Example-Backend/internal/adapter/server"
	"github.com/Vityalimbaev/Example-Backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

type APITestSuite struct {
	suite.Suite
	handlers           *gin.Engine
	defaultAccessToken string
	PrivateRequest     bool
	PublicRequest      bool
}

type TestCase struct {
	ExceptedError        bool
	RequestURL           string
	RequestMethod        string
	RequestData          interface{}
	ExpectedResponseData interface{}
	HttpStatusCode       int
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (suite *APITestSuite) SetupSuite() {
	suite.initApp()
}

func (suite *APITestSuite) TearDownSuite() {

}

func TestMain(m *testing.M) {
	rm := m.Run()
	os.Exit(rm)
}

func (suite *APITestSuite) initApp() {
	config.InitConfig()
	logger.SetupLogger()

	db := database.GetDbConnection(config.GetTestDbConfig())
	database.UpDBMigrations(db)

	suite.handlers = server.GetRouter()
	servicePool := adapter.InitServices(db)
	adapter.InitHandlers(suite.handlers, servicePool)

	suite.defaultAccessToken = suite.auth()
	suite.PrivateRequest = true
	suite.PublicRequest = false
}

func (suite *APITestSuite) auth() string {
	r := suite.Require()

	reqBody := map[string]interface{}{
		"email":    "admin@sadkomed.ru",
		"password": "123456",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-type", "application/json")

	resp := httptest.NewRecorder()
	suite.handlers.ServeHTTP(resp, req)

	r.Equal(resp.Result().StatusCode, 200)

	respMap := map[string]map[string]string{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	r.NoError(err)

	_ = json.Unmarshal(respBytes, &respMap)
	r.NotEqual(len(respMap["data"]["access_token"]), 0)
	return respMap["data"]["access_token"]
}

func (suite *APITestSuite) postRequest(url string, body interface{}, privateReq bool) ([]byte, int, error) {
	reqBodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api"+url, bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-type", "application/json")

	if privateReq {
		req.Header.Set("Authorization", "Bearer "+suite.defaultAccessToken)
	}

	resp := httptest.NewRecorder()
	suite.handlers.ServeHTTP(resp, req)
	respBytes, err := ioutil.ReadAll(resp.Body)
	return respBytes, resp.Result().StatusCode, err
}

func (suite *APITestSuite) getRequest(url string, body interface{}, privateReq bool) ([]byte, int, error) {
	reqBodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodGet, "/api"+url, bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-type", "application/json")

	if privateReq {
		req.Header.Set("Authorization", "Bearer "+suite.defaultAccessToken)
	}

	resp := httptest.NewRecorder()
	suite.handlers.ServeHTTP(resp, req)
	respBytes, err := ioutil.ReadAll(resp.Body)
	return respBytes, resp.Result().StatusCode, err
}

func (suite *APITestSuite) deleteByIdRequest(url string, id int, privateReq bool) ([]byte, int, error) {
	req, _ := http.NewRequest("DELETE", "/api"+url+"/"+strconv.Itoa(id), nil)
	req.Header.Set("Content-type", "application/json")

	if privateReq {
		req.Header.Set("Authorization", "Bearer "+suite.defaultAccessToken)
	}

	resp := httptest.NewRecorder()
	suite.handlers.ServeHTTP(resp, req)
	respBytes, err := ioutil.ReadAll(resp.Body)
	return respBytes, resp.Result().StatusCode, err
}
