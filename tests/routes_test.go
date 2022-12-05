package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/harsh/project/internal/service"
)

func Test_GetUsers(t *testing.T) {

	//
	r := mux.NewRouter()
	r.HandleFunc("/api/", service.GetUsers)
	req, _ := http.NewRequest(http.MethodGet, "/api/", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetUser(t *testing.T) {

	type tests struct {
		name string
		id   string
		err  error
	}

	testCases := []tests{
		{"valid id", "1", nil},
		{"invalid id", "a", fmt.Errorf("invalid id")},
		{"id not found", "100", fmt.Errorf("id not found")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := mux.NewRouter()
			r.HandleFunc("/api/{id}", service.GetUser)
			req, _ := http.NewRequest(http.MethodGet, "/api/"+tc.id, nil)
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))
			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
				t.Errorf("Expected error %v, got %v", tc.err, string(body))
			}
		})
	}
}

func TestCreateUser(t *testing.T) {

	type tests struct {
		name string
		user []byte
		err  error
	}

	testCases := []tests{
		{"valid user", []byte(`{"name":"harsh","email":"harsh@harsh.com","password":"123456"}`), nil},
		{"invalid user", []byte(`{"name":"harsh","email":90,"password":""}`), fmt.Errorf("invalid user")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := mux.NewRouter()
			r.HandleFunc("/api/", service.CreateUser)
			req, _ := http.NewRequest(http.MethodPost, "/api/", bytes.NewBuffer(tc.user))
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))
			if res.StatusCode != http.StatusCreated {
				t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.StatusCode)
				t.Errorf("Expected error %v, got %v", tc.err, string(body))
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {

	type tests struct {
		name string
		id   string
		user []byte
		err  error
	}

	testCases := []tests{
		{"valid user", "1", []byte(`{"name":"harsh verma","email":"harsh@harsh.com","password":"123456"} `), nil},
		{"invalid user", "a", []byte(`{"name":"harsh verma","email":90,"password":""}`), fmt.Errorf("invalid user")},
		{"id not found", "100", []byte(`{"name":"harsh verma","email":","password":""}`), fmt.Errorf("id not found")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := mux.NewRouter()
			r.HandleFunc("/api/{id}", service.UpdateUser)
			req, _ := http.NewRequest(http.MethodPut, "/api/"+tc.id, bytes.NewBuffer(tc.user))
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))
			if res.StatusCode != http.StatusCreated {
				t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.StatusCode)
				t.Errorf("Expected error %v, got %v", tc.err, string(body))
			}
		})
	}
}
