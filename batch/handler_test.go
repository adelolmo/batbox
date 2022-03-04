package batch

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bat := &Batch{
				directory:       tt.fields.directory,
				batchFiles:      tt.fields.batchFiles,
				variables:       tt.fields.variables,
				currentFile:     tt.fields.currentFile,
				currentFileLine: tt.fields.currentFileLine,
				arguments:       tt.fields.arguments,
				continueTo:      tt.fields.continueTo,
			}
			if got := bat.parseIfStatement(tt.args.line); got != tt.want {
				t.Errorf("parseIfStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
