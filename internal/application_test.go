package internal

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// week1 - 과제3
func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)
		res, err := app.Create(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		res, err := app.Create(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		res, err := app.Create(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		types := []string{"naver", "toss", "payco"}
		for i := 0; i < len(types); i++ {
			req := CreateRequest{strings.Join([]string{"jenny", strconv.Itoa(i)}, "-"), types[i]}
			res, err := app.Create(req)

			assert.Nil(t, err)
			assert.NotEmpty(t, res.ID)
			assert.Equal(t, req.MembershipType, res.MembershipType)
		}

		req := CreateRequest{"jenny", "iamport"}
		res, err := app.Create(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("멤버십 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		req := UpdateRequest{"1", "jenny", "toss"}
		res, err := app.Update(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		_, _ = app.Create(CreateRequest{"mike", "naver"})
		req := UpdateRequest{"2", "jenny", "naver"}
		res, err := app.Update(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		req := UpdateRequest{"", "jenny", "naver"}
		res, err := app.Update(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		req := UpdateRequest{"1", "", "naver"}
		res, err := app.Update(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		req := UpdateRequest{"1", "mike", ""}
		res, err := app.Update(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		req := UpdateRequest{"1", "mike", "kakao"}
		res, err := app.Update(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		err := app.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		err := app.Delete("")
		assert.NotNil(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		err := app.Delete("1")
		assert.NotNil(t, err)
	})
}
