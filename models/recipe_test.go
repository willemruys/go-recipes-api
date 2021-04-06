package models

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestRecipe_UpdateRecipe(t *testing.T) {
	type fields struct {
		Model       gorm.Model
		Title       string
		Ingredients string
	}
	type args struct {
		db       *gorm.DB
		recipeId string
		input    UpdateRecipe
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Recipe
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Recipe{
				Model:       tt.fields.Model,
				Title:       tt.fields.Title,
				Ingredients: tt.fields.Ingredients,
			}
			got, err := r.UpdateRecipe(tt.args.db, tt.args.recipeId, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Recipe.UpdateRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recipe.UpdateRecipe() = %v, want %v", got, tt.want)
			}
		})
	}
}
