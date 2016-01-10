package crud_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdate(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfile{}
	err := DB.Read(&nova, "SELECT * FROM user_profile WHERE name = 'Nova'")
	assert.Nil(t, err)

	nova.Bio = "Y O L O"
	assert.Nil(t, DB.Update(nova))

	novac := UserProfile{}
	err = DB.Read(&novac, "SELECT * FROM user_profile WHERE name = 'Nova'")
	assert.Nil(t, err)
	assert.Equal(t, novac.Bio, "Y O L O")
	assert.Equal(t, novac.Email, nova.Email)
	assert.Equal(t, novac.Id, nova.Id)
}

func TestUpdateNotMatching(t *testing.T) {
	assert.Nil(t, DB.Update(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}

func TestMustUpdate(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfile{}
	err := DB.Read(&nova, "SELECT * FROM user_profile WHERE name = 'Nova'")
	assert.Nil(t, err)

	nova.Bio = "Hola"
	assert.Nil(t, DB.MustUpdate(nova))

	novac := UserProfile{}
	err = DB.Read(&novac, "SELECT * FROM user_profile WHERE name = 'Nova'")
	assert.Nil(t, err)
	assert.Equal(t, novac.Bio, "Hola")
	assert.Equal(t, novac.Email, nova.Email)
	assert.Equal(t, novac.Id, nova.Id)
}

func TestMustUpdateNotMatching(t *testing.T) {
	assert.NotNil(t, DB.MustUpdate(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}