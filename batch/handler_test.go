package batch

import (
	"bytes"
	"strings"
	"testing"
)

func TestBatch_parseIfStatement(t *testing.T) {
	type fields struct {
		directory       string
		batchFiles      []string
		variables       map[string]string
		currentFile     string
		currentFileLine int
		arguments       map[string]string
		continueTo      string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "if variable 1st part",
			fields: fields{variables: map[string]string{"&amg": "xxx"}},
			args:   args{line: "IF \"%&amg%\"==\"xxx\" GOTO a5w"},
			want:   "GOTO a5w",
		},
		{
			name:   "if not variable 1st part",
			fields: fields{variables: map[string]string{"&amg": "xxx"}},
			args:   args{line: "IF NOT \"%&amg%\"==\"xxx\" GOTO a5w"},
			want:   "",
		},
		{
			name: "if variable 1st & 2nd part",
			fields: fields{variables: map[string]string{
				"&amg": "nhk",
				"&l5w": "nhk",
			}},
			args: args{line: "IF \"%&amg%\"==\"%&l5w%\" GOTO a5w"},
			want: "GOTO a5w",
		},
		{
			name: "if not variable 1st & 2nd part",
			fields: fields{variables: map[string]string{
				"&amg": "nhk",
				"&l5w": "nhk",
			}},
			args: args{line: "IF NOT \"%&amg%\"==\"%&l5w%\" GOTO a5w"},
			want: "",
		},
		{
			name:   "if argument 1st part equals string",
			fields: fields{arguments: map[string]string{"%1": "xxx"}},
			args:   args{line: "IF \"%1\"==\"xxx\" GOTO a5w"},
			want:   "GOTO a5w",
		},
		{
			name:   "if not argument 1st part equals string",
			fields: fields{arguments: map[string]string{"%1": "xxx"}},
			args:   args{line: "IF NOT \"%1\"==\"xxx\" GOTO a5w"},
			want:   "",
		},
		{
			name: "if argument 1st part equals variable",
			fields: fields{
				arguments: map[string]string{"%1": "xxx"},
				variables: map[string]string{"&amg": "xxx"},
			},
			args: args{line: "IF \"%1\"==\"%&amg%\" GOTO a5w"},
			want: "GOTO a5w",
		},
		{
			name: "if argument 1st part not equals variable",
			fields: fields{
				arguments: map[string]string{"%1": "xxx"},
				variables: map[string]string{"&amg": "yyy"},
			},
			args: args{line: "IF \"%1\"==\"%&amg%\" GOTO a5w"},
			want: "",
		},
		{
			name: "if not argument 1st part equals variable",
			fields: fields{
				arguments: map[string]string{"%1": "xxx"},
				variables: map[string]string{"&amg": "xxx"},
			},
			args: args{line: "IF NOT \"%1\"==\"%&amg%\" GOTO a5w"},
			want: "",
		},
		{
			name: "if not argument 1st part not equals variable",
			fields: fields{
				arguments: map[string]string{"%1": "xxx"},
				variables: map[string]string{"&amg": "yyy"},
			},
			args: args{line: "IF NOT \"%1\"==\"%&amg%\" GOTO a5w"},
			want: "GOTO a5w",
		},
		{
			name:   "if nested",
			fields: fields{arguments: map[string]string{"%1": "turtle"}},
			args:   args{line: "IF \"%1\"==\"turtle\" IF \"%&l5w%\"==\"a3w\" IF \"%&&a5w%\"==\"b7w\" GOTO v7e"},
			want:   "IF \"%&l5w%\"==\"a3w\" IF \"%&&a5w%\"==\"b7w\" GOTO v7e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bat := &Batch{
				directory:  tt.fields.directory,
				batchFiles: tt.fields.batchFiles,
				variables:  tt.fields.variables,
				arguments:  tt.fields.arguments,
				continueTo: tt.fields.continueTo,
			}
			if got := bat.handleIf(tt.args.line); got != tt.want {
				t.Errorf("handleIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatch_handleEcho(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
	type fields struct {
		directory       string
		batchFiles      []string
		variables       map[string]string
		currentFile     string
		currentFileLine int
		arguments       map[string]string
		continueTo      string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "echo",
			fields: fields{},
			args:   args{line: "ECHO a test"},
			want:   "a test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Batch{
				directory:  tt.fields.directory,
				batchFiles: tt.fields.batchFiles,
				variables:  tt.fields.variables,
				arguments:  tt.fields.arguments,
				continueTo: tt.fields.continueTo,
			}
			b.handleEcho(tt.args.line)
			got := strings.TrimSpace(buf.String())
			if got != tt.want {
				t.Errorf("handleEcho() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatch_handleSet(t *testing.T) {
	type fields struct {
		directory       string
		batchFiles      []string
		variables       map[string]string
		currentFile     string
		currentFileLine int
		arguments       map[string]string
		continueTo      string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{
			name:   "set value",
			fields: fields{variables: map[string]string{}},
			args:   args{line: "SET &l5w=m7w"},
			want:   map[string]string{"&l5w": "m7w"},
		},
		{
			name: "set variable",
			fields: fields{
				variables: map[string]string{
					"&l5w": "m7w",
				},
			},
			args: args{line: "SET &l5w=%&l5w%"},
			want: map[string]string{"&l5w": "m7w"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Batch{
				directory:  tt.fields.directory,
				batchFiles: tt.fields.batchFiles,
				variables:  tt.fields.variables,
				arguments:  tt.fields.arguments,
				continueTo: tt.fields.continueTo,
			}
			b.handleSet(tt.args.line)
			for wantKey := range tt.want {
				got := b.variables[wantKey]
				if tt.want[wantKey] != got {
					t.Errorf("handleSet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
