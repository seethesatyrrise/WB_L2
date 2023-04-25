package main

import (
	"reflect"
	"testing"
)

var testMap *anagrams = mapAnagrams([]string{"пятак", "сорт", "СТОЛИК", "ток",
	"листок", "рост", "ватка", "Тяпка", "торс", "кот", "пятка", "трос", "слиток"})

func Test_findPlace(t *testing.T) {
	type args struct {
		word  string
		words []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"1", args{"столик", []string{"листок", "слиток"}}, 2, false},
		{"2", args{"столик", []string{"листок", "слиток", "столик"}}, 0, true},
		{"3", args{"слиток", []string{"листок", "столик"}}, 1, false},
		{"4", args{"слиток", []string{"листок", "листок", "листок", "листок", "столик"}}, 4, false},
		{"5", args{"пятка", []string{"пятак", "тяпка"}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findPlace(tt.args.word, tt.args.words)
			if got != tt.want {
				t.Errorf("findPlace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findPlace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_anagrams_findAnagrams(t *testing.T) {
	type fields struct {
		words map[string][]string
		aux   map[string]string
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{"1", fields{testMap.words, testMap.aux},
			args{"кто"}, []string{"кот", "ток"}},
		{"2", fields{testMap.words, testMap.aux},
			args{"рост"}, []string{"рост", "сорт", "торс", "трос"}},
		{"3", fields{testMap.words, testMap.aux},
			args{"РОСТ"}, []string{"рост", "сорт", "торс", "трос"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aMap := &anagrams{
				words: tt.fields.words,
				aux:   tt.fields.aux,
			}
			if got := aMap.findAnagrams(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
