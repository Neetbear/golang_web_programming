package internal

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var member = Membership{
	ID:             uuid.New().String(),
	UserName:       "jenny",
	MembershipType: "naver",
}

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}

		res, err := app.Create(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}

		_, err := app.Create(req)
		assert.Nil(t, err)

		actual, err := app.Create(req)
		assert.ErrorAs(t, err, &ErrDuplicateName)
		assert.Empty(t, actual)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}

		res, err := app.Create(req)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, res)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}

		res, err := app.Create(req)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, res)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "daum"}

		res, err := app.Create(req)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, res)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("internal 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{}))
		req := CreateRequest{member.UserName, member.MembershipType}
		res, _ := app.Create(req)

		name, mType := "zico", "payco"
		u := UpdateRequest{
			ID:             res.ID,
			UserName:       name,
			MembershipType: mType,
		}
		updated, err := app.Update(u)

		assert.Nil(t, err)
		assert.Equal(t, res.ID, updated.ID)
		assert.Equal(t, name, updated.UserName)
		assert.Equal(t, mType, updated.MembershipType)
	})

	t.Run("If the membership username is duplicated > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		u := UpdateRequest{
			ID:             member.ID,
			UserName:       member.UserName,
			MembershipType: member.MembershipType,
		}
		actual, err := app.Update(u)

		assert.ErrorAs(t, err, &ErrDuplicateName)
		assert.Empty(t, actual)
	})

	t.Run("If the membership ID is not entered > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		u := UpdateRequest{
			ID:             "",
			UserName:       "zico",
			MembershipType: "naver",
		}
		actual, err := app.Update(u)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, actual)
	})

	t.Run("If the membership username is not entered > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		u := UpdateRequest{
			ID:             member.ID,
			UserName:       "",
			MembershipType: "naver",
		}
		actual, err := app.Update(u)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, actual)
	})

	t.Run("If the membership type is not entered > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		u := UpdateRequest{
			ID:             member.ID,
			UserName:       "zico",
			MembershipType: "",
		}
		actual, err := app.Update(u)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, actual)
	})

	t.Run("If not a given membership type > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		u := UpdateRequest{
			ID:             member.ID,
			UserName:       "zico",
			MembershipType: "daum",
		}
		actual, err := app.Update(u)

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, actual)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete membership", func(t *testing.T) {
		repository := NewRepository(map[string]Membership{
			member.ID: member,
		})
		app := NewApplication(repository)

		actual := app.Delete(member.ID)

		assert.Nil(t, actual)
		assert.Nil(t, repository.Get(member.ID))
	})

	t.Run("If not entered ID > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		actual := app.Delete("")

		assert.ErrorAs(t, actual, &ErrInvalidRequest)
	})

	t.Run("If entered ID does not exist > exception", func(t *testing.T) {
		app := NewApplication(NewRepository(map[string]Membership{
			member.ID: member,
		}))

		actual := app.Delete(uuid.New().String())

		assert.ErrorAs(t, actual, &ErrNotExist)
	})
}

func TestGet(t *testing.T) {
	t.Run("Get exist Membership", func(t *testing.T) {
		repository := NewRepository(map[string]Membership{
			member.ID: member,
		})
		app := NewApplication(repository)

		actual, err := app.Get(member.ID)

		assert.Nil(t, err)
		assert.Equal(t, member.ID, actual.ID)
		assert.Equal(t, member.UserName, actual.UserName)
		assert.Equal(t, member.MembershipType, actual.MembershipType)
	})

	t.Run("If Not exist Membership > return empty struct", func(t *testing.T) {
		repository := NewRepository(map[string]Membership{})
		app := NewApplication(repository)

		actual, err := app.Get(member.ID)

		assert.Nil(t, err)
		assert.Empty(t, actual)
	})

	t.Run("If not entered ID > exception", func(t *testing.T) {
		repository := NewRepository(map[string]Membership{
			member.ID: member,
		})
		app := NewApplication(repository)

		actual, err := app.Get("")

		assert.ErrorAs(t, err, &ErrInvalidRequest)
		assert.Empty(t, actual)
	})
}
