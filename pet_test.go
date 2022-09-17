package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate_OK_WithoutBuilders(t *testing.T) {
	owner := &Owner{
		ID:   ID(uuid.New().String()),
		Name: "Name owner",
		Age:  33,
	}

	data := []struct {
		name    string
		pet     *Pet
		isValid bool
	}{
		{
			name: "Valid pet",
			pet: &Pet{
				ID:    ID(uuid.New().String()),
				Name:  "Name pet",
				Age:   10,
				Owner: owner,
			},
			isValid: true,
		},
		{
			name: "Valid pet with empty owner",
			pet: &Pet{
				ID:    ID(uuid.New().String()),
				Name:  "Name pet",
				Age:   10,
				Owner: nil,
			},
			isValid: true,
		},
		{
			name: "Invalid id",
			pet: &Pet{
				ID:    "invalidID",
				Name:  "Name pet",
				Age:   10,
				Owner: owner,
			},
			isValid: false,
		},
		{
			name: "Invalid name",
			pet: &Pet{
				ID:    ID(uuid.New().String()),
				Name:  "invalid_name",
				Age:   10,
				Owner: owner,
			},
			isValid: false,
		},
		{
			name: "Invalid age",
			pet: &Pet{
				ID:    ID(uuid.New().String()),
				Name:  "Name pet",
				Age:   -10,
				Owner: owner,
			},
			isValid: false,
		},
		{
			name: "Invalid owner",
			pet: &Pet{
				ID:   ID(uuid.New().String()),
				Name: "Name pet",
				Age:  10,
				Owner: &Owner{
					ID:   ID(uuid.New().String()),
					Name: "invalid_owner_name",
					Age:  33,
				},
			},
			isValid: false,
		},
	}

	for _, tt := range data {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pet.Validate()
			isValid := err == nil
			fmt.Printf("%+v", tt.pet)
			assert.Equal(t, tt.isValid, isValid)
			if !tt.isValid {
				assert.ErrorIs(t, err, ValidationError)
			}
		})
	}
}
