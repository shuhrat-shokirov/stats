package stats

import (
	"testing"
	"reflect"
	"fmt"

	"github.com/shuhrat-shokirov/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   53_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "pharmacy",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "restaurant",
			Status:   types.StatusFail,
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)
	//Output:  1000000
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			Category: "car",
			Amount: 0,
		},
	}
	want := map[types.Category]types.Money{
		"car": 0,
	}
	got := CategoriesAvg(payments)
	if !reflect.DeepEqual(want,got){
		t.Errorf("want: %v, got: %v", want, got)
	}
}
