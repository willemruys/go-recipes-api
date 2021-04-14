package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/gorm"
	"recipes-api.com/m/auth"
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

type ListResponse struct {
	gorm.Model
	Title 		string  `json:"title"`
	Description string	`json:"description"`
	UserID 		uint64	`json:"userId"`	
}

type ResponseCreateRecipe struct {
	Recipe 		RecipeResponse 	`json:"recipe"`
	Response 	string 			`json:"response"`
}

type ResponseCreateList struct {
	List ListResponse `json:"list"`
}

type GetListResponse struct {
	List 		ListResponse 		`json:"list"`
	Recipes 	[]RecipeResponse 	`json:"recipes"`
}

type UpdateListResponse struct {
	List ListResponse `json:"list"`
}

func TestRunMainTests(t *testing.T) {

	db := models.SetupMockModels()
	ts := httptest.NewServer(setupServer(db))

	/* login */
	loginResUser1 := CreateUser(db, ts, t)
	loginResUser2 := CreateUserTwo(db, ts, t)

	/* user */
	UpdateUserPersonalDetails(ts, t, loginResUser1)
	UpdateUserPersonalDetailsUnauthorized(ts, t, loginResUser1)

	/* recipe */
	recipe := CreateRecipe(ts, t, loginResUser1)
	GetRecipe(ts, t, loginResUser1, recipe)
	CreateIncompleteRecipe(ts, t, loginResUser1)

	/* list */
	createListResponse := CreateListTest(ts, t, loginResUser1)
	GetListAfterCreate(ts, t, loginResUser1, createListResponse)
	updateListResponse := UpdateList(ts, t, loginResUser1, createListResponse)
	GetListAfterUpdate(ts, t, loginResUser1, updateListResponse)
	UpdateListUnauthorized(ts, t, loginResUser2, updateListResponse)
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

func CreateUserTwo(db *gorm.DB, ts *httptest.Server, t *testing.T) LoginResponse {

	/* Create user */
	createUser := map[string]string{"Email": "testuser2@gmail.com", "Username": "TestUser2", "Password": "password"}

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
	loginUser := map[string]string{"Email": "testuser2@gmail.com", "Password": "password"}

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


func UpdateUserPersonalDetails(ts *httptest.Server, t *testing.T, loginRes LoginResponse) {

	updateUserDetails := map[string]string{"Email": "updatedUser@gmail.com", "Username": "updatedUsername"}

	jsonUpdateUserDetails, _ := json.Marshal(updateUserDetails)
	userId, err := auth.ExtractTokenIDFromToken(loginRes.Token)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/user/personal-details/%v", ts.URL, userId), bytes.NewBuffer(jsonUpdateUserDetails))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }
}

func UpdateUserPersonalDetailsUnauthorized(ts *httptest.Server, t *testing.T, loginRes LoginResponse) {

	updateUserDetails := map[string]string{"Email": "updatedUserUnauth@gmail.com", "Username": "updatedUserUnauth"}

	jsonUpdateUserDetails, _ := json.Marshal(updateUserDetails)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/user/personal-details/%v", ts.URL, 9999), bytes.NewBuffer(jsonUpdateUserDetails))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 401 {
        t.Fatalf("Expected status code 401, got %v", completeResp.StatusCode)
    }
}

/* recipes test */
func CreateRecipe(ts *httptest.Server, t *testing.T, loginRes LoginResponse) ResponseCreateRecipe {

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


/* List test */
func CreateListTest(ts *httptest.Server, t *testing.T, loginRes LoginResponse) ResponseCreateList {

	createListInput := map[string]string{"Title": "my test list", "Description": "My test description"}

	jsonCreateListInput, _ := json.Marshal(createListInput)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/list", ts.URL), bytes.NewBuffer(jsonCreateListInput))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }

	var decoder = json.NewDecoder(completeResp.Body)

	var listResponse ResponseCreateList
	if err := decoder.Decode(&listResponse); err != nil {
		t.Fatalf("Error decoding json")
	}

	return listResponse

}

func GetListAfterCreate(ts *httptest.Server, t *testing.T, loginRes LoginResponse, listResponse ResponseCreateList) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/list/%v", ts.URL, listResponse.List.ID), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }

	var decoder = json.NewDecoder(completeResp.Body)

	var getListResponse GetListResponse
	if err := decoder.Decode(&getListResponse); err != nil {
		t.Fatalf("Error decoding json")
	}

	if getListResponse.List.ID != listResponse.List.ID {
		t.Fatalf("Expected matching list ids")
	}

	if getListResponse.List.Title != listResponse.List.Title {
		t.Fatalf("Expected matching list titles")
	}

	if getListResponse.List.Description != listResponse.List.Description {
		t.Fatalf("Expected matching list descriptions")
	}

}

func UpdateList(ts *httptest.Server, t *testing.T, loginRes LoginResponse, listResponse ResponseCreateList) UpdateListResponse {
	
	createListInput := map[string]string{"Title": "my test list adjusted", "Description": "My test description adjusted"}

	jsonCreateListInput, _ := json.Marshal(createListInput)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/list/%v", ts.URL, listResponse.List.ID), bytes.NewBuffer(jsonCreateListInput))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }

	var decoder = json.NewDecoder(completeResp.Body)

	var updateListResponse UpdateListResponse
	if err := decoder.Decode(&updateListResponse); err != nil {
		t.Fatalf("Error decoding json")
	}

	return updateListResponse

}


func GetListAfterUpdate(ts *httptest.Server, t *testing.T, loginRes LoginResponse, listResponse UpdateListResponse) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/list/%v", ts.URL, listResponse.List.ID), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }

	var decoder = json.NewDecoder(completeResp.Body)

	var getListResponse GetListResponse
	if err := decoder.Decode(&getListResponse); err != nil {
		t.Fatalf("Error decoding json")
	}

	if getListResponse.List.ID != listResponse.List.ID {
		t.Fatalf("Expected matching list ids")
	}

	if getListResponse.List.Title != listResponse.List.Title {
		t.Fatalf("Expected matching list titles")
	}

	if getListResponse.List.Description != listResponse.List.Description {
		t.Fatalf("Expected matching list descriptions")
	}

}

func UpdateListUnauthorized(ts *httptest.Server, t *testing.T, loginRes LoginResponse, listResponse UpdateListResponse) {

	createListInput := map[string]string{"Title": "my test list adjusted", "Description": "My test description adjusted"}

	jsonCreateListInput, _ := json.Marshal(createListInput)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/list/%v", ts.URL, listResponse.List.ID), bytes.NewBuffer(jsonCreateListInput))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginRes.Token))

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err.Error())
    }

	if completeResp.StatusCode != 401 {
        t.Fatalf("Expected status code 401, got %v", completeResp.StatusCode)
    }

}