package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
	"git.druva.org/druva.com/DI-Master/apimocker"
	"git.druva.org/druva.com/DI-Master/config"
	"git.druva.org/druva.com/DI-Master/rerror"
	"git.druva.org/druva.com/DI-Master/restlib/authjwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestGetTotalBackupDataForDayshandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetTotalBackupDataForDayshandler", GetTotalBackupDataForDayshandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{
		{
			name:       "Positive | GetTotalBackupDataForDayshandler Success case",
			body:       `{"startDate":0,"endDate":100}`,
			statuscode: 200,
			customerId: 103,
		},
		{
			name:       "Negative |  GetTotalBackupDataForDayshandler failure case",
			body:       ``,
			statuscode: 400,
			customerId: 103,
		},
		{
			name:       "Negative | GetTotalBackupDataForDayshandler failure empty customer_id",
			body:       `{"startDate":0,"endDate":100}`,
			statuscode: 504,
			customerId: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetTotalBackupDataForDayshandler", strings.NewReader(tt.body))

			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got testName:%s  %d\n", tt.statuscode, tt.name, w.Code)
			}
		})
	}

}

func TestGetFileTypesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	r.POST("/GetFileTypesHandler", GetFileTypesHandler)
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{
		{
			name:       "Positive | GetFileTypesHandler Success case",
			body:       `{"date":0}`,
			statuscode: 200,
			customerId: 305,
		},
		{
			name:       "Negative | GetFileTypesHandler with empty body",
			body:       ``,
			statuscode: 400,
			customerId: 305,
		},
		{
			name:       "Negative | GetFileTypesHandler with empty jwt",
			body:       `{"date":0}`,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetFileTypesHandler", strings.NewReader(tt.body))

			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)
			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestGetFileExtensionsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetFileExtensionsHandler", GetFileExtensionsHandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{
		{
			name:       "Positive | valid file id",
			body:       `{"fileId":1}`,
			statuscode: 200,
			customerId: 305,
		},
		{
			name:       "Negative | empty body",
			body:       ``,
			statuscode: 400,
			customerId: 305,
		},
		{
			name:       "Negative | empty jwt",
			body:       `{}`,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetFileExtensionsHandler", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)
			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestGetDataByResourceTypehandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetDataByResourceTypehandler", GetDataByResourceTypehandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{

		{
			name:       "Positive | invalid File id",
			body:       `{}`,
			statuscode: 200,
			customerId: 201,
		},
		{
			name:       "Negative | empty body",
			body:       ``,
			statuscode: 400,
			customerId: 201,
		},
		{
			name:       "Negative | invalid jwt",
			body:       `{}`,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetDataByResourceTypehandler", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestGetDataSourceGroupshandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetDataSourceGroupshandler", GetDataSourceGroupshandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{
		{
			name:       "Positive | Success case",
			body:       ``,
			statuscode: 200,
			customerId: 201,
		},
		{
			name:       "Positive | Success case",
			body:       ``,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetDataSourceGroupshandler", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)
			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestGetFileTypesAndExtensionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetFileTypesAndExtensionHandler", GetFileTypesAndExtensionHandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{
		{
			name:       "Positive | Success case get API ",
			body:       ``,
			statuscode: 200,
			customerId: 305,
		},
		{
			name:       "Negative | invalid jwt",
			body:       ``,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetFileTypesAndExtensionHandler", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)
			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestGetSourceListhandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetSourceListhandler", GetSourceListhandler)
	cqeClient = apimocker.MockCQEClient{}
	remoteClient = apimocker.NewMockPhoenixClient()
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
	}{
		{
			name:       "Negative | empty body",
			body:       ``,
			statuscode: 400,
			customerId: 302,
		},
		{
			name:       "Positive | Success Case",
			body:       `{"sourceId":1}`,
			statuscode: 200,
			customerId: 302,
		},
		{
			name:       "Negative | invalid Source ID",
			body:       `{"sourceId":1}`,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetSourceListhandler", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)
			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestGGetDistributionBySizeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/GetDistributionBySizeHandler", GetDistributionBySizeHandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
		customerId int
		flag       bool
	}{
		{
			name:       "Negative | Empty request body",
			body:       ``,
			statuscode: 400,
			customerId: 445,
		},
		{
			name:       "Positive| Success Case",
			body:       `{"date":1111}`,
			statuscode: 200,
			customerId: 445,
		},
		{
			name:       "Positive| Success Case",
			body:       `{"date":1111}`,
			statuscode: 504,
			customerId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			req, err := http.NewRequest(http.MethodPost, "/GetDistributionBySizeHandler", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			req.Header.Set("Authontication", jwt)

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}

}

func TestDistributionByModifiedMonthhandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/distributionByModifiedMonth", DistributionByModifiedMonthhandler)
	cqeClient = apimocker.MockCQEClient{}
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
	}{
		{
			name:       "Negative | Empty request body",
			body:       ``,
			statuscode: 400,
		},
		{
			name:       "Positive| Success Case",
			body:       `{"month":11111}`,
			statuscode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodPost, "/distributionByModifiedMonth", strings.NewReader(tt.body))
			jwt := CreateTestJWT(665)
			req.Header.Set("Authontication", jwt)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}
}

func TestGetOrganizationshandler(t *testing.T) {
	remoteClient = apimocker.NewMockPhoenixClient()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/getOrganizations", GetOrganizationshandler)
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	RemoteLicensingClient = apimocker.MockLicenseServiceJob{}

	defer ctrl.Finish()
	tests := []struct {
		name       string
		body       string
		statuscode int
	}{
		{
			name:       "Negative | Empty request body",
			body:       ``,
			statuscode: 400,
		},
		{
			name:       "Positive| Success Case",
			body:       `{"date":1111}`,
			statuscode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodPost, "/getOrganizations", strings.NewReader(tt.body))
			jwt := CreateTestJWT(765)
			req.Header.Set("Authontication", jwt)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Check to see if the response was what you expected
			if w.Code != tt.statuscode {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.statuscode, w.Code)
			}
		})
	}
}

func TestErrorResponse(t *testing.T) {
	testerr := rerror.New("test")
	rootError := make(map[string]string)
	rootError["error"] = testerr.Error()
	tests := []struct {
		name   string
		err    error
		output error
	}{
		{
			name:   "Positive | Success Case",
			err:    rerror.Error(rerror.InternalError),
			output: rerror.Error(rerror.InternalError),
		},
		{
			name:   "Positive | Error type Error ",
			err:    testerr,
			output: rerror.Error(rerror.InternalError, nil, rootError),
		},
		{
			name:   "Positive | Error type Rerror",
			err:    rerror.Error(rerror.InvalidArguments),
			output: rerror.Error(rerror.InternalError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ErrorResponse(tt.err)
			if !reflect.DeepEqual(output, tt.output) {
				t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", output, tt.output)
			}
		})
	}
	output := ErrorResponse(nil)
	if !reflect.DeepEqual(output, rerror.Error(rerror.InternalError)) {
		t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", output, rerror.Error(rerror.InternalError))
	}
}

func CreateTestJWT(id int) string {
	var jwtdata authjwt.Jwtstruct
	adminType := 2
	if id == 100 {
		adminType = 3 //non druva admin
	}
	jwtdata = authjwt.Jwtstruct{AdminData: authjwt.AdminData{"GlobalAdminID5", "Email5@druva.com", "Name5", adminType, "Timezone5", "Customer5", "cust1", "Accessor", true, false},
		SessionData: authjwt.SessionData{104, "TokenSecret5", 1115.16},
		ProductData: authjwt.ProductData{{[]int{12289, id, 7}, "Secret5"}, {[]int{8193, id, 7}, "Secret5"}},
		ClientData:  authjwt.ClientData{"ClientID5", "ClientName5", "Customer5", 106, "GlobalCustomerID5", "APIKey5"},
		Expiry:      time.Now().Unix() + 2000000,
		Signature:   "Signature5"}

	jwtString, _ := authjwt.GenerateJWT(getMap(jwtdata))
	return jwtString
}
func getMap(jwtinput authjwt.Jwtstruct) map[string]interface{} {
	jwtJson, err := json.Marshal(jwtinput)
	if err != nil {
		return nil
	}
	jwtMap := make(map[string]interface{})
	err = json.Unmarshal(jwtJson, &jwtMap)
	return jwtMap
}

func TestGetCustomerAndProductIds(t *testing.T) {
	RemoteLicensingClient = apimocker.NewMockLicenseServiceJob(false)

	tests := []struct {
		name       string
		customerId int
		flag       bool
		lFlag      bool
	}{

		{
			name:       "Positive| Success Case",
			customerId: 111,
			flag:       false,
			lFlag:      false,
		},
		{
			name:       "Positive| Success Case",
			customerId: 11,
			flag:       false,
			lFlag:      false,
		},
		{
			name:       "Negative| invalid jwt Case",
			customerId: 0,
			flag:       true,
			lFlag:      false,
		},
		{
			name:       "Negative| doesn't have Licence",
			customerId: 111,
			flag:       true,
			lFlag:      true,
		},
		{
			name:       "Negative| product admin does not have api access",
			customerId: 100, // it is for Product admin
			flag:       true,
			lFlag:      false,
		},		
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jwt string
			if tt.customerId > 0 {
				jwt = CreateTestJWT(tt.customerId)
			}
			if tt.lFlag {
				RemoteLicensingClient = apimocker.NewMockLicenseServiceJob(true)
			}
			//jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcm9kdWN0X2RhdGEiOlt7InNlY3JldCI6ImhhdGlqdUJYZlNxRDBya0Q2N0ZLVHBVK2NtcXFyUXF1aUE1MkZvWWNjTkRIYng4MDFwNjNWb3BiZHJXVkFBa0lBZ1ZDRGd0M1FEd3lvdDl6Wk82UkVmaUtBaCtRZ1Fua0I4K2xJWWdIcXpSQVRWVVZXRjUvRWs4M3NXYURKYzc0Z29nV1lIYld5YVpSU0IrWXdtcUY2eXBwa0RuUENtVkJTcm1uRHRvU0FjOWpMblJGcUpqM3d0dHBFNng5Y0dJb0toYU0vN1JGWlRZVE9TSDRqU3FOKy9DTkJmUzUvSEpDM0EveWhRZGE3MmZLWmliNjFHVlYwNmtTdDc1WUsweW1KTis5NSs1S2ozN3NncngvMUZpbXFzRy9GZmJEWWQydXdoSVlOcnFhb214bC90VFFFMUptRzVhdVhKQlRWdnh4IiwiaWRzIjpbNDA5Nyw0LDBdfSx7InNlY3JldCI6ImhhdDZ4YVQrb2plVGlyZVVEcVVNdUJFQ1ZRWW4xWk5WNUNpdU9RL3NSZVY4WDJmd1d5SjhvaTBUQ1Z3d0I4Z0JJYk5aQWdsZmNPYXozMThNUUV2UXRiT243T3VKdnVxZEZkamdNT2NPR2llOUxMeDJhVjdUWFRpTXA2eWplbUpleXZIRDZOeU1LS3lFMUdDbTdNUmVuQW9OM2Q3UWRnUGQ5UVpha2wxYTh4R3dVZURmQXV5OFoxT3dGb08zclpGSU5iVXdOVjI1UUZ6Nm9XM1lKV3R1cGFMcERLZ2pBODRmWWRNRDE5dVlxeEpSRUl1NEh5RlllQWk0aEdGcTdVWkI1YUtFREdWYTVsU0xCMnBPNkJpV1Z4QVhuV1pYYmt6Z0NKcG9EdlJhQjhseWk2NDVjOTVxTWlmSDJMN0FpRGFuIiwiaWRzIjpbMTIyODksMiwwXX1dLCJleHBpcnkiOjE1Njg3ODc5NzIsImNsaWVudF9kYXRhIjp7ImN1c3RvbWVyIjoibmFzXzEiLCJjbGllbnRfbmFtZSI6ImNyZWRzIiwicm9sZSI6MiwiY2xpZW50X2lkIjoiN1crd1pLUi9STndqT25yNUovT1dQZDk0N0hsYi9SRzIiLCJhcGlfa2V5IjpudWxsLCJnbG9iYWxfY3VzdG9tZXJfaWQiOiI5ZDY0OWY2NC00MTBmLTQyN2UtOTYyMC1kZGY4Yjc2ZDhmMzktNCJ9LCJzaWduYXR1cmUiOiJ1NkY3OWwzTFFSOXR2ejZ5Q29qdkkxUHB4WHlXZjJpK2UxN2p5RDUyM1JrPSIsImFkbWluX2RhdGEiOnt9LCJzZXNzaW9uX2RhdGEiOnsic2Vzc2lvbl9jcmVhdGlvbl90aW1lIjoxNTY4NzAxNTA2LjQxMDU0NCwic2Vzc2lvbl90b2tlbl9leHBpcnkiOjE1Njg3MDMzNzIsInRva2VuX3NlY3JldCI6IkNaUnpWQ0VBTl9pbTVzcE5ETk1Od3c9PSJ9fQ.ye6MdXH7BvuU0sgYVI4UcaNoQFJErHl9fo_ED9jVgjc"
			c := CreateContext(jwt)
			customerProductIds, err := GetCustomerAndProductIds(c)
			if err != nil && !tt.flag {
				t.Fatalf("Failded to Read JWT: %v\n", err)
			} else if tt.flag && err == nil {
				t.Fatalf("Failed to Read jwt: %v\n", err)
			}
			if customerProductIds[config.ProductIdPhoenix] != tt.customerId && !tt.flag {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

		})
	}
}

// function to create the gin.Context for testing
func CreateContext(val string) *gin.Context {

	var c gin.Context
	req, _ := http.NewRequest(http.MethodPost, "/Test", strings.NewReader("{}"))
	c.Request = req
	c.Request.Header.Set("Authontication", val)

	return &c
}
