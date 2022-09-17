package main

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/google/uuid"
)

var ValidationError = errors.New("invalid property")

type ID string

func (i ID) Validate() error {
	if _, err := uuid.Parse(string(i)); err != nil {
		return fmt.Errorf("%w with value %s: %v", ValidationError, i, err)
	}
	return nil
}

type Name string

func (n Name) Validate() error {
	for _, r := range n {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return fmt.Errorf("%w with value %s", ValidationError, n)
		}
	}
	return nil
}

type Age int

func (a Age) Validate() error {
	if a <= 0 {
		return fmt.Errorf("%w with value %d", ValidationError, a)
	}
	return nil
}

type Owner struct {
	ID   ID
	Name Name
	Age  Age
}

func (o *Owner) Validate() error {
	if err := o.ID.Validate(); err != nil {
		return err
	}
	if err := o.Name.Validate(); err != nil {
		return err
	}
	if err := o.Age.Validate(); err != nil {
		return err
	}
	return nil
}

type Pet struct {
	ID    ID
	Name  Name
	Age   Age
	Owner *Owner
}

func (p *Pet) Validate() error {
	if err := p.ID.Validate(); err != nil {
		return err
	}
	if err := p.Name.Validate(); err != nil {
		return err
	}
	if err := p.Age.Validate(); err != nil {
		return err
	}

	if p.Owner != nil {
		if err := p.Owner.Validate(); err != nil {
			return fmt.Errorf("owner of pet %s is invalid: %w", p.ID, err)
		}
	}
	return nil
}
