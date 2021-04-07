package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/gorm"
	"recipes-api.com/m/models"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type RecipeResponse struct {
	gorm.Model
	Title 		string 	`json:"title"`
	Ingredients string 	`json:"ingredients"`
	UserID 		uint 	`json:"UserID"`
	Likes 		int 	`json:"likes"`
}

type ResponseCreateRecipe struct {
	Recipe 		RecipeResponse 	`json:"recipe"`
	Response 	string 			`json:"response"`
}

func TestRunMainTests(t *testing.T) {

	db := models.SetupMockModels()
	ts := httptest.NewServer(setupServer(db))

	loginRes := CreateUser(db, ts, t)
	recipe := CreateRecipe(ts, t, loginRes)
	GetRecipe(ts, t, loginRes, recipe)
	CreateIncompleteRecipe(ts, t, loginRes)
}

func CreateUser(db *gorm.DB, ts *httptest.Server, t *testing.T) LoginResponse {

	/* Create user */
	createUser := map[string]string{"Email": "testuser@gmail.com", "Username": "TestUser", "Password": "password"}

	jsonCreateUserBody, _ := json.Marshal(createUser)

	req, errCreateUserReq := http.NewRequest("POST", fmt.Sprintf("%s/user", ts.URL), bytes.NewBuffer(jsonCreateUserBody))

	if errCreateUserReq != nil {
		t.Fatalf("Received error creating user. Error: %s", errCreateUserReq)
	}

	res, errCreateUser := http.DefaultClient.Do(req)

	if errCreateUser != nil {
		t.Fatalf("Received error creating user. Error: %s", errCreateUser)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, received %v", res.StatusCode)
	}
	
	
	/* Create user */
	loginUser := map[string]string{"Email": "testuser@gmail.com", "Password": "password"}

	jsonLoginUser, _ := json.Marshal(loginUser)

	req, errLoginUserReq := http.NewRequest("POST", fmt.Sprintf("%s/login", ts.URL), bytes.NewBuffer(jsonLoginUser))

	if errLoginUserReq != nil {
		t.Fatalf("Received error creating user. Error: %s", errCreateUserReq)
	}

	res, errLoginUser := http.DefaultClient.Do(req)

	if errLoginUser != nil {
		t.Fatalf("Received error creating user. Error: %s", errLoginUser)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, received %v", res.StatusCode)
	}

	var decoder = json.NewDecoder(res.Body)

	var loginRes LoginResponse

	if err := decoder.Decode(&loginRes); err != nil {
		t.Fatalf("Error decoding json")
	}

	return loginRes

}

func CreateRecipe(ts *httptest.Server, t *testing.T, loginRes LoginResponse) ResponseCreateRecipe {

	/* Create recipes */
	createRecipeValuesComplete := map[string]string{"Title": "my test recipe", "Ingredients": "my test ingredients"}

	jsonCreateRecipeValuesComplete, _ := json.Marshal(createRecipeValuesComplete)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/recipes", ts.URL), bytes.NewBuffer(jsonCreateRecipeValuesComplete))

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if res.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", res.StatusCode)
    }

	var decoder = json.NewDecoder(res.Body)

	var recipeResponse ResponseCreateRecipe
	if err := decoder.Decode(&recipeResponse); err != nil {
		t.Fatalf("Error decoding json")
	}

	return recipeResponse

}

func GetRecipe(ts *httptest.Server, t *testing.T, loginRes LoginResponse, recipe ResponseCreateRecipe) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/recipes/%v", ts.URL, recipe.Recipe.ID), nil)
	
	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }

}


func CreateIncompleteRecipe(ts *httptest.Server, t *testing.T, loginRes LoginResponse) {

	/* Create recipes */
	createRecipeValuesComplete := map[string]string{"Title": "my test recipe"}

	jsonCreateRecipeValuesComplete, _ := json.Marshal(createRecipeValuesComplete)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/recipes", ts.URL), bytes.NewBuffer(jsonCreateRecipeValuesComplete))

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if completeResp.StatusCode != 400 {
        t.Fatalf("Expected status code 400, got %v", completeResp.StatusCode)
    }
	
}