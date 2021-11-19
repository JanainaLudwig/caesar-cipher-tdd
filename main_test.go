package main

import "testing"

func TestCrypt(t *testing.T) {
	type args struct {
		text string
		key int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "encriptação básica",
			args: args{
				text: "abc",
				key: 1,
			},
			want: "bcd",
		},
		{
			name: "encriptação básica 2",
			args: args{
				text: "abc",
				key: 2,
			},
			want: "cde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crypt(tt.args.text, tt.args.key); got != tt.want {
				t.Errorf("Crypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptAndDecrypt(t *testing.T) {
	type args struct {
		text string
		key int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "encriptação básica",
			args: args{
				text: "abc",
				key: 10,
			},
			want: "abc",
		},
		{
			name: "encriptação básica",
			args: args{
				text: "áaaaíaó",
				key: 10,
			},
			want: "aaaaiao",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypted := Crypt(tt.args.text, tt.args.key)
			t.Log("crypted: ", crypted)

			if got := Decrypt(crypted, tt.args.key); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		text string
		key  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "decriptação básica 2",
			args: args{
				text: "cde",
				key: 2,
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.text, tt.args.key); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}