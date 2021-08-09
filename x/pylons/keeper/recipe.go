package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

// SetRecipe set a specific recipe in the store from its ID
func (k Keeper) SetRecipe(ctx sdk.Context, recipe types.Recipe) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecipeKey+recipe.CookbookID))
	b := k.cdc.MustMarshalBinaryBare(&recipe)
	store.Set(types.KeyPrefix(recipe.ID), b)
}

// GetRecipe returns a recipe from its ID
func (k Keeper) GetRecipe(ctx sdk.Context, cookbookID string, id string) (val types.Recipe, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecipeKey+cookbookID))

	b := store.Get(types.KeyPrefix(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// GetAllRecipe returns all recipe
func (k Keeper) GetAllRecipe(ctx sdk.Context) (list []types.Recipe) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecipeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		// nolint: staticcheck
		if err != nil {
		}
	}(iterator)

	for ; iterator.Valid(); iterator.Next() {
		var val types.Recipe
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllRecipesByCookbook returns all recipes owned by cookbook
func (k Keeper) GetAllRecipesByCookbook(ctx sdk.Context, cookbookID string) (list []types.Recipe) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecipeKey+cookbookID))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		// nolint: staticcheck
		if err != nil {
		}
	}(iterator)

	for ; iterator.Valid(); iterator.Next() {
		var val types.Recipe
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}