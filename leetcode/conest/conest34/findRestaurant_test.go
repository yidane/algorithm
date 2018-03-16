package conest34

import (
	"reflect"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}}, want: []string{"Shogun"}},
		{name: "2", args: args{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, list2: []string{"KFC", "Shogun", "Burger King"}}, want: []string{"Shogun"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
