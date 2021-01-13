package db

import (
	"../schema"
	"reflect"
	"testing"
	"time"
)

//func TestSample_Close(t *testing.T) {
//
//	tests := []struct {
//		name string
//	}{
//
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Sample{}
//		})
//	}
//}

func TestSample_Delete(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSample_GetAll(t *testing.T) {

	want := []schema.Todo{
		{
			ID:      1,
			Title:   "Do dishes",
			Note:    "",
			DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:      2,
			Title:   "Do homework",
			Note:    "",
			DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:      2,
			Title:   "Twitter",
			Note:    "",
			DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	tests := []struct {
		name    string
		want    []schema.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Get 3 todo's",want,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample_Insert(t *testing.T) {
	type args struct {
		todo *schema.Todo
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"insert none", args{nil},0,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{}
			got, err := s.Insert(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
