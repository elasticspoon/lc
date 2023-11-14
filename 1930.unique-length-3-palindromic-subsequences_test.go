package main

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "aabca", want: 3},
		{s: "adc", want: 0},
		{s: "bbcbaba", want: 4},
		{s: "huafuhlavlucuharcifuhasdjhflaksjfhdlkjashflkjsahlkjfhslkjdfhlksjdfhlkasjhglkhhgsfkjgsufgyaugefkuyagfkuygasefkcyugbafygsaefycgasgfasefgcsjugycfasuygfasgfjasgyfjgbyuascfusyagebfgcfjgbseaycfyusgcfgycjfefjsegycfusaygefusyagcfsakycfksauycfksaegbyfcksbyugafksybgecfsaukygcefuyksagcfuysagcfuysakcfuyksacyuksfcskfcysacfuksayfcukysgacfuskyfbuyskfcyukasfbksayfbcskaycfsfbskjycfeskycfbksycfbksajcsayfbksacsacfsbkcsacfsacfysfcygsfcygsafcygsacgfyscgfysycgfjsagecfeygcfsjgycfjgyscfysgefjsgacfyegscfajbsgycfysagbcfjasgecfysgefycsajfcgsajfgbcsjgyfjsgbfcjsgyecfyjgsebfcyjsbfygsabcyasjcfgasegycfyesbafsygcfseygfcsyegf", want: 191},
		{s: "abcdefghijklmnopqrstuvwxyz", want: 0},
		{s: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", want: 676},
	}

	for _, tt := range tests {
		got := countPalindromicSubsequence(tt.s)
		if got != tt.want {
			t.Errorf("countPalindromicSubsequence(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
