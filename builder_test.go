package main

import "github.com/google/uuid"

type petOption func(pet *Pet)

func PetBuilder(opts ...petOption) *Pet {
	pet := &Pet{}
	for _, opt := range opts {
		opt(pet)
	}
	return pet
}

func WithID(id string) func(pet *Pet) {
	return func(pet *Pet) {
		pet.ID = ID(id)
	}
}

func WithName(name string) func(pet *Pet) {
	return func(pet *Pet) {
		pet.Name = Name(name)
	}
}

func WithAge(age int) func(pet *Pet) {
	return func(pet *Pet) {
		pet.Age = Age(age)
	}
}

func WithOwner(owner *Owner) func(pet *Pet) {
	return func(pet *Pet) {
		pet.Owner = owner
	}
}

func NewPetBuilder(opts ...petOption) *Pet {
	pet := &Pet{
		ID:   ID(uuid.New().String()),
		Name: "Dummy Pet Name",
		Age:  10,
		Owner: &Owner{
			ID:   ID(uuid.New().String()),
			Name: "Dummy Owner Name",
			Age:  33,
		},
	}
	for _, opt := range opts {
		opt(pet)
	}
	return pet
}

func NewPetWithoutOwnerBuilder(opts ...petOption) *Pet {
	pet := &Pet{
		ID:   ID(uuid.New().String()),
		Name: "Dummy Pet Name",
		Age:  10,
	}
	for _, opt := range opts {
		opt(pet)
	}
	return pet
}
