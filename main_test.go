package main

import (
	"basic-gin/controller"
	"basic-gin/handle"
	"basic-gin/model"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddData(t *testing.T) {

	routeTest := gin.Default()

	m := handle.NewMember()

	routeTest.GET("/member", controller.AllData(m))
	routeTest.POST("/member", controller.AddData(m))

	t.Run("GET Data", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/member", nil)
		res := httptest.NewRecorder()

		routeTest.ServeHTTP(res, req)

		response := res.Result()
		body, _ := ioutil.ReadAll(response.Body)

		result := model.Member{}
		_ = json.Unmarshal(body, &result)

		expect := model.Member{}

		assert.Equal(t, expect, result)

	})

	t.Run("Add Data", func(t *testing.T) {

		data := `{
			"name":"charan",
			"phone":"0123456789",
			"email":"charan@ru.ac.th"
		}`

		req := httptest.NewRequest("POST", "/member", strings.NewReader(data))
		res := httptest.NewRecorder()

		routeTest.ServeHTTP(res, req)

		response := res.Result()
		body, _ := ioutil.ReadAll(response.Body)

		result := model.Member{}
		_ = json.Unmarshal(body, &result)

		expect := model.Member{
			Name:  "charan",
			Phone: "0123456789",
			Email: "charan@ru.ac.th",
		}

		assert.Equal(t, expect, result)
	})
}
