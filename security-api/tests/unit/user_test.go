package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"security-api/controllers"
	userModel "security-api/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Variables
var (
	uc controllers.UserController
)


// -----------------------
// MODEL TEST
// -----------------------

func TestUserModel(t *testing.T) {
	testModel := userModel.User{
		UserID: 10,
		Name: "Keylor",
		Email: "test@test.com",
		Password: "admin1234",
		IsActive: true,
	}

	if testModel.UserID != 10 {
		t.Errorf("\n UserID %v expected to be 10\n", testModel.UserID)
	}

	if testModel.Name != "Keylor" {
		t.Errorf("\n Name %v expected to be Keylor\n", testModel.Name)
	}

	if testModel.Email != "test@test.com" {
		t.Errorf("\n Email %v expected to be test@test.com\n", testModel.Email)
	}

	if testModel.Password != "admin1234" {
		t.Errorf("\n Password %v expected to be admin1234\n", testModel.Password)
	}

	if testModel.IsActive != true {
		t.Errorf("\n IsActive %v expected to be true\n", testModel.IsActive)
	}
}

// ---------------------
// TEST GIN
// ---------------------

// test
func TestGetUsers(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := GetTestGinContext(w)

	u := url.URL{
		RawPath: "http://localhost:9090/v1/user/getall",
	}

	MockJsonGet(ctx, u)

	uc.GetAll(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code)

	//got, _ := strconv.Atoi(w.Body.String())

	//assert.Equal(t, 1, got)
}

// Mock Gin Context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)	// create test gin context (mock)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{},
	}

	return ctx
}

// test mock GET
func MockJsonGet(c *gin.Context, u url.URL) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	//c.Request.RequestURI = u.RawPath
	c.Request.URL.RawPath = u.RawPath
}

