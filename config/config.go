package config

import (
	"ag/api"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DefaultProvider string                    `yaml:"default_provider"`
	Providers       map[string]ProviderConfig `yaml:"providers"`
}

type ProviderConfig struct {
	Type     string `yaml:"type"`
	Endpoint string `yaml:"endpoint"`
	APIKey   string `yaml:"api_key"`
	Model    string `yaml:"model"`
}

var config *Config

func getConfigPaths() []string {
	// 获取用户主目录
	home, err := os.UserHomeDir()
	if err != nil {
		return []string{}
	}

	// 按照优先级返回配置路径
	return []string{
		filepath.Join(home, ".local", "bin", "ag", "config.yaml"), // 主配置路径
		filepath.Join(home, ".config", "ag", "config.yaml"),       // XDG 配置路径
		"/etc/ag/config.yaml",   // 系统级配置
		"config.yaml",           // 当前目录
		"./config/config.yaml",  // config 子目录
		"../config.yaml",        // 上一级目录
		"../config/config.yaml", // 上一级目录的 config 子目录
	}
}

func Load() error {
	// 尝试多个可能的配置文件路径
	configPaths := getConfigPaths()

	var data []byte
	var err error

	// 尝试读取配置文件
	for _, path := range configPaths {
		data, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		return fmt.Errorf("无法找到配置文件，尝试了以下路径：\n%s", strings.Join(configPaths, "\n"))
	}

	// 解析YAML
	config = &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 初始化API提供商
	for name, cfg := range config.Providers {
		switch cfg.Type {
		case "volcengine":
			api.RegisterProvider(name, api.NewVolcEngineClient(cfg.APIKey))
		case "openai":
			api.RegisterProvider(name, api.NewOpenAIClient(cfg.APIKey))
		case "deepseek":
			api.RegisterProvider(name, api.NewDeepSeekClient(cfg.APIKey))
		default:
			return fmt.Errorf("未知的提供商类型: %s", cfg.Type)
		}
	}

	return nil
}

func GetDefaultProvider() string {
	if config == nil {
		return ""
	}

	return config.DefaultProvider
}

func GetProviderConfig(name string) *ProviderConfig {
	if config == nil {
		return nil
	}
	if cfg, ok := config.Providers[name]; ok {
		return &cfg
	}
	return nil
}
