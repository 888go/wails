package options

import (
	"testing"
)

func TestMergeDefaultsWH(t *testing.T) {
	tests := []struct {
		name       string
		appoptions *App
		wantWidth  int
		wantHeight int
	}{
		{
			name:       "No width and height",
			appoptions: &App{},
			wantWidth:  1024,
			wantHeight: 768,
		},
		{
			name: "Basic width and height",
			appoptions: &App{
				X宽度:  800,
				X高度: 600,
			},
			wantWidth:  800,
			wantHeight: 600,
		},
		{
			name: "With MinWidth and MinHeight",
			appoptions: &App{
				X宽度:     200,
				X最小宽度:  800,
				X高度:    100,
				X最小高度: 600,
			},
			wantWidth:  800,
			wantHeight: 600,
		},
		{
			name: "With MaxWidth and MaxHeight",
			appoptions: &App{
				X宽度:     900,
				X最大宽度:  800,
				X高度:    700,
				X最大高度: 600,
			},
			wantWidth:  800,
			wantHeight: 600,
		},
		{
			name: "With MinWidth more than MaxWidth",
			appoptions: &App{
				X宽度:    900,
				X最小宽度: 900,
				X最大宽度: 800,
				X高度:   600,
			},
			wantWidth:  800,
			wantHeight: 600,
		},
		{
			name: "With MinHeight more than MaxHeight",
			appoptions: &App{
				X宽度:     800,
				X高度:    700,
				X最小高度: 900,
				X最大高度: 600,
			},
			wantWidth:  800,
			wantHeight: 600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeDefaults(tt.appoptions)
			if tt.appoptions.X宽度 != tt.wantWidth {
				t.Errorf("MergeDefaults().Width =%v, want %v", tt.appoptions.X宽度, tt.wantWidth)
			}
			if tt.appoptions.X高度 != tt.wantHeight {
				t.Errorf("MergeDefaults().Height =%v, want %v", tt.appoptions.X高度, tt.wantHeight)
			}
		})
	}
}
