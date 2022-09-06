package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang_web_programming/internal/dto"
	"golang_web_programming/internal/repository"
	"testing"
)

var (
	dumyID       = uuid.New().String()
	dumyUsername = "jenny"
	dumyType     = "naver"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(repository.NewRepository(map[string]repository.Membership{}))
		req := dto.CreateRequest{"jenny", "naver"}

		res, err := app.Create(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("internal 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(repository.NewRepository(map[string]repository.Membership{}))
		req := dto.CreateRequest{
			UserName:       dumyUsername,
			MembershipType: dumyType,
		}
		res, _ := app.Create(req)

		name, mType := "zico", "payco"
		u := dto.UpdateRequest{
			ID: res.ID,
			UpdateRequestBody: dto.UpdateRequestBody{
				UserName:       name,
				MembershipType: mType,
			},
		}
		updated, err := app.Update(u)

		assert.Nil(t, err)
		assert.Equal(t, res.ID, updated.ID)
		assert.Equal(t, name, updated.UserName)
		assert.Equal(t, mType, updated.MembershipType)
	})

	t.Run("If the membership username is duplicated > exception", func(t *testing.T) {
		app := NewApplication(repository.NewRepository(map[string]repository.Membership{
			dumyID: *repository.CreateMembership(dumyID, dumyUsername, dumyType),
		}))
		u := dto.UpdateRequest{
			ID: dumyID,
			UpdateRequestBody: dto.UpdateRequestBody{
				UserName:       dumyUsername,
				MembershipType: dumyType,
			},
		}

		actual, err := app.Update(u)

		assert.ErrorAs(t, err, &repository.ErrDuplicateName)
		assert.Empty(t, actual)
	})

}

func TestDelete(t *testing.T) {
	t.Run("Delete membership", func(t *testing.T) {
		repository := repository.NewRepository(map[string]repository.Membership{
			dumyID: *repository.CreateMembership(dumyID, dumyUsername, dumyType),
		})
		app := NewApplication(repository)

		actual := app.Delete(dumyID)

		assert.Nil(t, actual)
		assert.Nil(t, repository.Get(dumyID))
	})

	t.Run("If entered iD does not exist > exception", func(t *testing.T) {
		app := NewApplication(repository.NewRepository(map[string]repository.Membership{
			dumyID: *repository.CreateMembership(dumyID, dumyUsername, dumyType),
		}))

		actual := app.Delete(uuid.New().String())

		assert.ErrorAs(t, actual, &repository.ErrNotExist)
	})
}

func TestGet(t *testing.T) {
	t.Run("Get exist Membership", func(t *testing.T) {
		repository := repository.NewRepository(map[string]repository.Membership{
			dumyID: *repository.CreateMembership(dumyID, dumyUsername, dumyType),
		})
		app := NewApplication(repository)

		actual, err := app.Get(dumyID)

		assert.Nil(t, err)
		assert.Equal(t, dumyID, actual.ID)
		assert.Equal(t, dumyUsername, actual.UserName)
		assert.Equal(t, dumyType, actual.MembershipType)
	})

	t.Run("If Not exist Membership > return empty struct", func(t *testing.T) {
		repository := repository.NewRepository(map[string]repository.Membership{})
		app := NewApplication(repository)

		actual, err := app.Get(dumyID)

		assert.Nil(t, err)
		assert.Empty(t, actual)
	})
}
