package caesar

import "testing"

func TestCaesar_rotateChar(t *testing.T) {
	type fields struct {
		Key int
	}
	type args struct {
		char uint8
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "convert a",
			fields: fields{
				Key: 1,
			},
			args:   args{
				char: 'a',
				key: 1,
			},
			want:   "b",
		},
		{
			name:   "convert z",
			fields: fields{
				Key: 2,
			},
			args:   args{
				char: 'z',
				key: 2,
			},
			want:   "b",
		},
		{
			name:   "decrypt z",
			fields: fields{
				Key: -2,
			},
			args:   args{
				char: 'b',
				key: -2,
			},
			want:   "z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Caesar{
				Key: tt.fields.Key,
			}
			if got := c.rotateChar(tt.args.char, tt.args.key); got != tt.want {
				t.Errorf("rotateChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaesar_Crypt(t *testing.T) {
	type fields struct {
		Key int
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "crypt basic text",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "abc de",
			},
			want:   "bcd ef",
		},
		{
			name:   "crypt basic text",
			fields: fields{
				Key: 2,
			},
			args:   args{
				text: "azb",
			},
			want:   "cbd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Caesar{
				Key: tt.fields.Key,
			}
			if got := c.Crypt(tt.args.text); got != tt.want {
				t.Errorf("Crypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaesar_Decrypt(t *testing.T) {
	type fields struct {
		Key int
	}
	type args struct {
		cryptedText string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test basic decrypt",
			fields: fields{
				Key: 2,
			},
			args:   args{
				cryptedText: "cbd",
			},
			want:   "azb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Caesar{
				Key: tt.fields.Key,
			}
			if got := c.Decrypt(tt.args.cryptedText); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}