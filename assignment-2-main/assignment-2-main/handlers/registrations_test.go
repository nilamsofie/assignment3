package handlers_test

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/handlers"
	"assignment-2/webhooks"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"strconv"
	"testing"

	"net/http"
)

type TestData struct {
	Method       string
	Id           string
	RequestBody  io.Reader
	Statuscode   int
	ResponseBody string
}

func post(body map[string]interface{}) error {
	// Parse body to JSON
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req := httptest.NewRequest("POST", "/dashboard/v1/registrations", bytes.NewBuffer(bodyJSON))
	rec := httptest.NewRecorder()

	handlers.RegistrationsPost(rec, req)

	if rec.Code == http.StatusOK {
		return nil
	}

	return errors.New("Failed to post.")
}

func getAll() error {
	req := httptest.NewRequest("GET", "/dashboard/v1/registrations", nil)
	rec := httptest.NewRecorder()

	handlers.GetAllRegistrations(rec, req)

	if rec.Code == http.StatusOK {
		var body []map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &body)
		if err != nil {
			return err
		}
		if body[0]["country"] == "Norway" && body[1]["country"] == "Germany" {
			return nil
		}
	}

	return errors.New("Failed to get all.")
}

func getById(id int) error {
	idStr := strconv.Itoa(id)

	req := httptest.NewRequest("GET", "/dashboard/v1/registrations/"+idStr, nil)
	req.SetPathValue("id", idStr)

	rec := httptest.NewRecorder()

	handlers.GetRegistrationsById(rec, req)

	if rec.Code == http.StatusOK {
		var body map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &body)
		if err != nil {
			return err
		}
		if body["country"] != "Germany" {
			return errors.New("failed to get by id")
		}
	}

	return nil
}

func delete(id int) error {
	idStr := strconv.Itoa(id)

	req := httptest.NewRequest("DELETE", "/dashboard/v1/registrations/"+idStr, nil)
	req.SetPathValue("id", idStr)

	rec := httptest.NewRecorder()

	handlers.RegistrationsDelete(rec, req)

	if rec.Code != http.StatusNoContent {
		return errors.New("Failed to delete.")
	}

	return nil
}

func stringPointer(str string) *string {
	return &str
}

func testBody(index int) *bytes.Buffer {
	var m map[string]interface{}

	switch index {
	case 0:
		//Body with missing/typo in field
		m = map[string]interface{}{
			"counry": "Norway",
			"features": map[string]interface{}{
				"area":             true,
				"capital":          true,
				"coordinates":      false,
				"population":       false,
				"precipitation":    true,
				"targetCurrencies": []string{"EUR", "USD", "SE"},
				"temperature":      true,
				"map":              true,
			},
		}
	case 1:
		//Body with empty country field
		m = map[string]interface{}{
			"country": "",
			"features": map[string]interface{}{
				"area":             true,
				"capital":          true,
				"coordinates":      false,
				"population":       false,
				"precipitation":    true,
				"targetCurrencies": []string{"EUR", "USD", "SE"},
				"temperature":      true,
				"map":              true,
			},
		}
	case 2:
		//Empty body
		m = map[string]interface{}{}
	case 3:
		//Body with no valid fields
		m = map[string]interface{}{
			"country": "Norway",
			"featwures": map[string]interface{}{
				"aread":             true,
				"capwital":          true,
				"coordindates":      false,
				"populadtion":       false,
				"precidpitation":    true,
				"tarwgetCurrencies": []string{"EUR", "USD", "SE"},
				"tempderature":      true,
				"mdap":              true,
			},
		}
	case 4:
		//BOdy with missing value of field
		m = map[string]interface{}{
			"country": "Norway",
			"features": map[string]interface{}{
				"area":             nil,
				"capital":          true,
				"coordinates":      false,
				"population":       false,
				"precipitation":    true,
				"targetCurrencies": []string{"EUR", "USD", "SE"},
				"temperature":      true,
				"map":              true,
			},
		}
	case 5:
		//BOdy with both isoCode and country
		m = map[string]interface{}{
			"country": "Norway",
			"isoCode": "NO",
			"features": map[string]interface{}{
				"area":             nil,
				"capital":          true,
				"coordinates":      false,
				"population":       false,
				"precipitation":    true,
				"targetCurrencies": []string{"EUR", "USD", "SE"},
				"temperature":      true,
				"map":              true,
			},
		}
	default:
		break
	}

	body, err := json.Marshal(m)
	if err != nil {
		log.Printf("Error når parser test body: %v", err)
	}
	return bytes.NewBuffer(body)

}

func badrequestTests() error {

	//Testing bad requests
	var tests []TestData = []TestData{
		{Method: "GET", Id: "0", Statuscode: http.StatusBadRequest, ResponseBody: constants.InvalidIDMessage + "\n"},
		{Method: "GET", Id: "345", Statuscode: http.StatusNotFound, ResponseBody: constants.DataNotFound + "\n"},
		{Method: "GET", Id: "gege", Statuscode: http.StatusBadRequest, ResponseBody: constants.InvalidIDMessage + "\n"},
		{Method: "DELETE", Id: "-12", Statuscode: http.StatusBadRequest, ResponseBody: constants.InvalidIDMessage + "\n"},
		{Method: "DELETE", Id: "423", Statuscode: http.StatusNotFound, ResponseBody: constants.DataNotFound + "\n"},
		{Method: "DELETE", Id: "hihi", Statuscode: http.StatusBadRequest, ResponseBody: constants.InvalidIDMessage + "\n"},
		{
			Method:       "PUT",
			Id:           "1",
			Statuscode:   http.StatusBadRequest,
			ResponseBody: constants.InvalidRequestBodyMessage + "\n",
			RequestBody:  testBody(0),
		},
		{
			Method:       "PUT",
			Id:           "2",
			Statuscode:   http.StatusBadRequest,
			ResponseBody: constants.InvalidRequestBodyMessage + "\n",
			RequestBody:  testBody(1),
		},
		{
			Method:       "PUT",
			Id:           "2",
			Statuscode:   http.StatusBadRequest,
			ResponseBody: constants.InvalidRequestBodyMessage + "\n",
			RequestBody:  testBody(2),
		},
		{
			Method:       "POST",
			Id:           "",
			Statuscode:   http.StatusBadRequest,
			ResponseBody: constants.InvalidRequestBodyMessage + "\n",
			RequestBody:  testBody(0),
		},
		{
			Method:       "POST",
			Id:           "",
			Statuscode:   http.StatusBadRequest,
			ResponseBody: constants.InvalidRequestBodyMessage + "\n",
			RequestBody:  testBody(1),
		},
		{
			Method:       "POST",
			Id:           "",
			Statuscode:   http.StatusBadRequest,
			ResponseBody: "invalid registration body, must have either country or isoCode, but not both\n",
			RequestBody:  testBody(5),
		},
	}

	for _, test := range tests {
		var path string = "/dashboard/v1/registrations"
		if test.Method != "POST" {
			path += "/"
			path += test.Id
		}

		req := httptest.NewRequest(test.Method, path, test.RequestBody)

		if test.Method != "POST" {
			req.SetPathValue("id", test.Id)
		}

		rec := httptest.NewRecorder()

		switch test.Method {
		case http.MethodPost:
			handlers.RegistrationsPost(rec, req)
		case http.MethodGet:
			handlers.GetRegistrationsById(rec, req)
		case http.MethodDelete:
			handlers.RegistrationsDelete(rec, req)
		case http.MethodPut:
			handlers.RegistrationsPut(rec, req)
			//Patch not implemented for in-memory database
			// case "PATCH":
			// 	handlers.RegistrationsPatch(rec, req)
		}

		if rec.Code != test.Statuscode {
			fmt.Println(rec.Body.String(), rec.Code)
			return errors.New(test.Method + ": unexpected statuscode")
		}

		if rec.Body.String() != test.ResponseBody {
			fmt.Println(rec.Body.String(), rec.Code, test.Id)
			return errors.New(test.Method + ": unexpected responsebody")
		}
	}
	return nil
}

// Tests posting two dashboard configuarations, getting, and deleting
func TestRegistrations(t *testing.T) {
	// initialize in-memory database
	db := database.InitializeInMemoryDatabase()
	handlers.Db = &db
	webhooks.Db = &db

	postBody1 := map[string]interface{}{
		"country": nil,
		"isoCode": stringPointer("NO"),
		"features": map[string]interface{}{
			"area":             true,
			"capital":          true,
			"coordinates":      false,
			"population":       false,
			"precipitation":    true,
			"targetCurrencies": []string{"EUR", "USD", "SE"},
			"temperature":      true,
			"map":              true,
		},
	}

	postBody2 := map[string]interface{}{
		"country": stringPointer("Germany"),
		"isoCode": nil,
		"features": map[string]interface{}{
			"area":             false,
			"capital":          true,
			"coordinates":      true,
			"population":       true,
			"precipitation":    false,
			"targetCurrencies": []string{"CAD", "JPY", "NOK"},
			"temperature":      false,
			"map":              false,
		},
	}

	err := post(postBody1)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}

	err = post(postBody2)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}

	err = getAll()
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}

	err = getById(2)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}

	err = delete(1)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}

	err = delete(2)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}

	err = badrequestTests()
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error: %v", err)
	}
}
