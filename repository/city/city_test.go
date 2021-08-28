package city

import (
	"testing"
)

func Test_repository_PopulateCities(t *testing.T) {
	tests := []struct {
		name string
		r    *repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{}
			r.PopulateCities()
		})
	}
}

func Test_repository_FindCity(t *testing.T) {
	type args struct {
		id int32
	}
	tests := []struct {
		name    string
		r       *repository
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{}
			got, err := r.FindCity(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.FindCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.FindCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
