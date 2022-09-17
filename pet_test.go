package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const validNamePet, validNameOwner = "Name Pet", "Name owner"

var (
	validUUID = uuid.New().String()
	validAge  = 10
	owner     = &Owner{
		ID:   ID(uuid.New().String()),
		Name: validNameOwner,
		Age:  33,
	}
)

func TestValidate_WithoutBuilders(t *testing.T) {
	data := []struct {
		name    string
		pet     *Pet
		isValid bool
	}{
		{
			name: "Valid pet",
			pet: &Pet{
				ID:    ID(validUUID),
				Name:  validNamePet,
				Age:   Age(validAge),
				Owner: owner,
			},
			isValid: true,
		},
		{
			name: "Valid pet with empty owner",
			pet: &Pet{
				ID:    ID(validUUID),
				Name:  validNamePet,
				Age:   Age(validAge),
				Owner: nil,
			},
			isValid: true,
		},
		{
			name: "Invalid id",
			pet: &Pet{
				ID:    "invalidID",
				Name:  validNamePet,
				Age:   Age(validAge),
				Owner: owner,
			},
			isValid: false,
		},
		{
			name: "Invalid name",
			pet: &Pet{
				ID:    ID(validUUID),
				Name:  "invalid_name",
				Age:   Age(validAge),
				Owner: owner,
			},
			isValid: false,
		},
		{
			name: "Invalid age",
			pet: &Pet{
				ID:    ID(validUUID),
				Name:  validNamePet,
				Age:   -10,
				Owner: owner,
			},
			isValid: false,
		},
		{
			name: "Invalid owner",
			pet: &Pet{
				ID:   ID(validUUID),
				Name: validNamePet,
				Age:  Age(validAge),
				Owner: &Owner{
					ID:   owner.ID,
					Name: "invalid_owner_name",
					Age:  owner.Age,
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

func TestValidate_WithBuilders(t *testing.T) {
	data := []struct {
		name    string
		pet     *Pet
		isValid bool
	}{
		{
			name:    "Valid pet",
			pet:     PetBuilder(WithID(uuid.New().String()), WithName(validNamePet), WithAge(validAge), WithOwner(owner)),
			isValid: true,
		},
		{
			name:    "Valid pet with empty owner",
			pet:     PetBuilder(WithID(uuid.New().String()), WithName(validNamePet), WithAge(validAge)),
			isValid: true,
		},
		{
			name:    "Invalid id",
			pet:     PetBuilder(WithID("invalidID"), WithName(validNamePet), WithAge(validAge), WithOwner(owner)),
			isValid: false,
		},
		{
			name:    "Invalid name",
			pet:     PetBuilder(WithID(uuid.New().String()), WithName("invalid_name"), WithAge(validAge), WithOwner(owner)),
			isValid: false,
		},
		{
			name:    "Invalid age",
			pet:     PetBuilder(WithID(uuid.New().String()), WithName(validNamePet), WithAge(-10), WithOwner(owner)),
			isValid: false,
		},
		{
			name: "Invalid owner",
			pet: PetBuilder(WithID(uuid.New().String()), WithName(validNamePet), WithAge(-10), WithOwner(
				&Owner{
					ID:   ID(uuid.New().String()),
					Name: "invalid_owner_name",
					Age:  33,
				})),
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

func TestValidate_WithBuilders_ObjectMother(t *testing.T) {
	data := []struct {
		name    string
		pet     *Pet
		isValid bool
	}{
		{
			name:    "Valid pet",
			pet:     NewPetBuilder(),
			isValid: true,
		},
		{
			name:    "Valid pet with empty owner",
			pet:     NewPetWithoutOwnerBuilder(),
			isValid: true,
		},
		{
			name:    "Invalid id",
			pet:     NewPetBuilder(WithID("invalidID")),
			isValid: false,
		},
		{
			name:    "Invalid name",
			pet:     NewPetBuilder(WithName("invalid_name")),
			isValid: false,
		},
		{
			name:    "Invalid age",
			pet:     NewPetBuilder(WithAge(-3)),
			isValid: false,
		},
		{
			name: "Invalid owner",
			pet: NewPetBuilder(WithOwner(&Owner{
				ID:   ID(uuid.New().String()),
				Name: "invalid_owner_name",
				Age:  33,
			})),
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
