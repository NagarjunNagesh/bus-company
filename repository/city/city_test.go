package city

import (
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
)

var barce = "Barcelona"

func setup() {
	Cities = []*city.City{
		{
			ID:   1,
			Name: barce,
		},
		{
			ID:   2,
			Name: "Seville",
		},
		{
			ID:   3,
			Name: "Madrid",
		},
		{
			ID:   4,
			Name: "Valencia",
		},
		{
			ID:   5,
			Name: "Andorra la Vella",
		},
		{
			ID:   6,
			Name: "Malaga",
		},
	}
}

func Test_repository_FindCity(t *testing.T) {
	setup()
	type args struct {
		id int32
	}
	type testCases struct {
		name    string
		r       *repository
		args    args
		want    *string
		wantErr bool
	}
	tests := []testCases{
		{
			name: "FindData",
			r:    &repository{},
			args: args{
				id: 1,
			},
			want:    &barce,
			wantErr: false,
		},
		{
			name: "FindData",
			r:    &repository{},
			args: args{
				id: 7,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{}
			got, err := r.FindCity(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.FindCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && *got != *tt.want {
				t.Errorf("repository.FindCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
