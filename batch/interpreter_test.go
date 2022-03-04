package batch

import "testing"

func TestBatch_checkBatchFile(t *testing.T) {
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
		name          string
		fields        fields
		args          args
		want          string
		wantIsBatch   bool
		wantArguments map[string]string
	}{
		{
			name: "literal arguments",
			fields: fields{
				batchFiles: []string{"ASK.BAT"},
				arguments:  map[string]string{},
			},
			args:        args{line: "ASK owner about turtle"},
			want:        "ASK.BAT",
			wantIsBatch: true,
			wantArguments: map[string]string{
				"%1": "owner",
				"%2": "about",
				"%3": "turtle",
			},
		}, {
			name: "sequential arguments",
			fields: fields{
				batchFiles: []string{"ASK.BAT"},
				arguments: map[string]string{
					"%1": "owner",
					"%2": "about",
					"%3": "turtle",
				}},
			args:        args{line: "ASK %1 %2 %3"},
			want:        "ASK.BAT",
			wantIsBatch: true,
			wantArguments: map[string]string{
				"%1": "owner",
				"%2": "about",
				"%3": "turtle",
			},
		}, {
			name: "non sequential arguments",
			fields: fields{
				batchFiles: []string{"ASK.BAT"},
				arguments: map[string]string{
					"%1": "owner",
					"%2": "about",
					"%3": "turtle",
				}},
			args:        args{line: "ASK %1 %3 %4"},
			want:        "ASK.BAT",
			wantIsBatch: true,
			wantArguments: map[string]string{
				"%1": "owner",
				"%2": "turtle",
				"%3": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Batch{
				directory:       tt.fields.directory,
				batchFiles:      tt.fields.batchFiles,
				variables:       tt.fields.variables,
				currentFile:     tt.fields.currentFile,
				currentFileLine: tt.fields.currentFileLine,
				arguments:       tt.fields.arguments,
				continueTo:      tt.fields.continueTo,
			}
			got, got1 := b.checkBatchFile(tt.args.line)
			if got != tt.want {
				t.Errorf("checkBatchFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantIsBatch {
				t.Errorf("checkBatchFile() got isBatch = %v, want %v", got1, tt.wantIsBatch)
			}
			for wantArgName := range tt.wantArguments {
				if tt.wantArguments[wantArgName] != b.arguments[wantArgName] {
					t.Errorf(
						"checkBatchFile() got argument '%v' with value=%v, want argument '%v' with value=%v",
						wantArgName,
						b.arguments[wantArgName],
						wantArgName,
						tt.wantArguments[wantArgName],
					)

				}
			}
		})
	}
}
