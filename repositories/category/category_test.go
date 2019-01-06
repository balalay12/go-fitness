package category

import (
	"github.com/stretchr/testify/require"
	"go-fitness/database"
	"testing"
)

func TestService_GetCategory(t *testing.T) {
	db, err := database.NewDB("file:../../app.db??cache=shared&mode=rwc")
	require.NoError(t, err)

	service := NewCategoryService(db)
	cat, err := service.GetCategory(1)
	require.NoError(t, err)
	require.Equal(t, int64(1), cat.ID)
	require.Equal(t, "Грудь", cat.Name)
}

func TestService_GetAllCategories(t *testing.T) {
	db, err := database.NewDB("file:../../app.db??cache=shared&mode=rwc")
	require.NoError(t, err)

	service := NewCategoryService(db)
	cats, err := service.GetAllCategories()
	require.NoError(t, err)
	require.Len(t, cats, 13)
}
