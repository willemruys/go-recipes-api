package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"recipes-api.com/m/models"
)

func TestRecipeRoutes(t *testing.T) {
	db:= models.SetupMockModels()
	ts := httptest.NewServer(setupServer(db))
	/* Create recipes */
	createRecipeValuesComplete := map[string]string{"Title": "my test recipe", "Ingredients": "my test ingredients"}

	jsonCreateRecipeValuesComplete, _ := json.Marshal(createRecipeValuesComplete)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/recipes", ts.URL), bytes.NewBuffer(jsonCreateRecipeValuesComplete))

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTc3ODY5NjEsInVzZXJfaWQiOjM5fQ.V82INybiYgFeTtYzHnMlfr3SQs8A9ZNTkf0aMtmEiNw")

	completeResp, err := http.DefaultClient.Do(req)

	if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if completeResp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", completeResp.StatusCode)
    }

	/* create invalid recipe */

	createRecipeValuesIncomplete := map[string]string{"Title": "my test recipe"}

	jsonCreateRecipeValuesInComplete, _ := json.Marshal(createRecipeValuesIncomplete)

	incompleteResp, err := http.Post(fmt.Sprintf("%s/recipes", ts.URL), "application/json", bytes.NewBuffer(jsonCreateRecipeValuesInComplete))

    if incompleteResp.StatusCode != 400 {
        t.Fatalf("Expected bad request %v", incompleteResp.StatusCode)
    }

}