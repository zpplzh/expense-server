package catalog

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound  = errors.New("data not found")
	ErrDuplicate = errors.New("data duplicate")
	BadInput     = errors.New("wrong input")
	ErrAuth      = errors.New("User not exist")
)

type Service interface {
	GetCategory(ctx context.Context, input *GetCategoryInput) (*CategoryOutput, error)
	AddCategory(ctx context.Context, input *AddCategoryInput) (*CategoryOutput, error)
	DelCategory(ctx context.Context, input *DelCategoryInput) error
	//UpdateCategory(ctx context.Context, input *UpdateCategoryInput) error
	ListCategories(ctx context.Context, input *ListCategoriesInput) ([]*CategoryOutput, error)

	GetExpense(ctx context.Context, input *GetExpenseInput) (*ExpenseOutput, error)
	AddExpense(ctx context.Context, input *AddExpenseInput) (*AddExpenseOutput, error)
	ListExpense(ctx context.Context, input *ListExpensesInput) ([]*ExpenseOutput, error)
	DelExpense(ctx context.Context, input *DelExpenseInput) error
	//UpdateExpense(ctx context.Context, input *UpdateExpenseInput) error

	SignUp(ctx context.Context, input *SignUpInput) (*SignUpOutput, error)
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
	Logout(ctx context.Context, input *Logoutinput) (*Logoutoutput, error)
}

type servicedb struct {
	db *sqlx.DB
}

func NewServices(db *sqlx.DB) Service {
	return &servicedb{
		db: db,
	}
}
