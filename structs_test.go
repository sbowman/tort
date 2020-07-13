package tort_test

import (
	"testing"

	"github.com/sbowman/tort"
)

func TestStructNil(t *testing.T) {
	assert := tort.For(t)

	type Sample struct {
	}

	var ptr *Sample
	assert.Struct(ptr).IsNil()

	ptr = new(Sample)
	assert.Struct(ptr).IsNotNil()

	var sample Sample
	assert.Struct(sample).IsNotNil()
}

func TestStructDeepDive(t *testing.T) {
	assert := tort.For(t)

	type Planet struct {
		Name string
		PopB int8
	}

	type Address struct {
		City   string
		Planet Planet
	}

	type User struct {
		Email   string
		Address Address
	}

	user := User{
		Email: "kirk@starfleet.gov",
		Address: Address{
			City: "San Francisco",
			Planet: Planet{
				Name: "Earth",
				PopB: 6,
			},
		},
	}

	assert.Struct(user).String("Email").NotBlank()
	assert.Struct(user).String("Email").Matches(`\w+@starfleet\.gov`)
	assert.Struct(user).String("Email").Contains("starfleet.gov")

	assert.Struct(user).Struct("Address").String("City").NotEquals("New York")
	assert.Struct(user).Struct("Address").String("City").Equals("San Francisco")

	assert.Struct(user).Struct("Address").Struct("Planet").String("Name").Equals("Earth")
	assert.Struct(user).Struct("Address").Struct("Planet").Int8("PopB").Equals(6)
}
