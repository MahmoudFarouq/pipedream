package main

import (
	"reflect"
	"testing"
)

func TestEnvironment_GetByPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := map[string]struct {
		e       Environment
		args    args
		want    interface{}
		wantErr bool
	}{
		"notfound": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				path: "asd",
			},
			want:    nil,
			wantErr: true,
		},
		"nested notfound": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				path: "work.badpeople",
			},
			want:    nil,
			wantErr: true,
		},
		"simple found": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				path: "name",
			},
			want:    "mahmoud",
			wantErr: false,
		},
		"nested found": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				path: "work.name",
			},
			want:    "Flash",
			wantErr: false,
		},
		"not very nested found": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				path: "work.name.first",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.e.GetByPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironment_Render(t *testing.T) {
	type args struct {
		input string
	}
	tests := map[string]struct {
		e       Environment
		args    args
		want    string
		wantErr bool
	}{
		"simple: no templates": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				input: "name",
			},
			want:    "name",
			wantErr: false,
		},
		"simple: replace single string": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				input: "{{name}}",
			},
			want:    "mahmoud",
			wantErr: false,
		},
		"simple: replace single with nesting": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				input: "{{work.name}}",
			},
			want:    "Flash",
			wantErr: false,
		},
		"simple: replace multipe with nesting": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name": "Flash",
				},
			},
			args: args{
				input: "my name is {{name}} and i work at {{work.name}}",
			},
			want:    "my name is mahmoud and i work at Flash",
			wantErr: false,
		},
		"if the result is int": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name":  "Flash",
					"stNum": 10,
				},
			},
			args: args{
				input: "my name is {{name}} and i work at {{work.name}} which is in st num {{work.stNum}}",
			},
			want:    "my name is mahmoud and i work at Flash which is in st num 10",
			wantErr: false,
		},
		"if the result is map it should be json": {
			e: map[string]interface{}{
				"name": "mahmoud",
				"age":  15,
				"work": map[string]interface{}{
					"name":  "Flash",
					"stNum": 10,
				},
			},
			args: args{
				input: "{{work}}",
			},
			want:    `{"name":"Flash","stNum":10}`,
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.e.Render(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Render() got = %v, want %v", got, tt.want)
			}
		})
	}
}
