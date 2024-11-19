package config

import (
	"os"
	"testing"
)

func TestConfig_GetEnv(t *testing.T) {

	if err := os.Setenv("BOT_TOKEN", "TEST_BOT_TOKEN"); err != nil {
		t.Errorf("Error: %v", err)
	}

	if err := os.Setenv("AI_TOKEN", "TEST_AI_TOKEN"); err != nil {
		t.Errorf("Error: %v", err)
	}

	if err := os.Setenv("CHANNEL", "123"); err != nil {
		t.Errorf("Error: %v", err)
	}

	cfg := New()
	if err := cfg.GetEnv(); err != nil {
		t.Errorf("Error: %v", err)
	}

	type fields struct {
		BotToken string
		AiToken  string
		Channel  int64
	}

	f := fields{
		BotToken: cfg.BotToken,
		AiToken:  cfg.AiToken,
		Channel:  cfg.Channel,
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Configuration", f, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnf := &Config{
				BotToken: tt.fields.BotToken,
				AiToken:  tt.fields.AiToken,
				Channel:  tt.fields.Channel,
			}
			if err := cnf.GetEnv(); (err != nil) != tt.wantErr {
				t.Errorf("GetEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	if err := os.Unsetenv("BOT_TOKEN"); err != nil {
		t.Error(err)
	}

	if err := os.Unsetenv("AI_TOKEN"); err != nil {
		t.Error(err)
	}

	if err := os.Unsetenv("CHANNEL"); err != nil {
		t.Error(err)
	}
}
