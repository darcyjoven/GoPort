package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func generateSign(provider, username, realName, department string, timestamp int64, secretKey string) string {
	signStr := fmt.Sprintf("provider=%s&username=%s&realName=%s&department=%s×tamp=%d",
		provider, username, realName, department, timestamp)

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}

func loginTest(baseURL, provider, username, realName, department, sign, token string, timestamp int64) {
	params := url.Values{}
	params.Add("provider", provider)
	params.Add("username", username)
	params.Add("realName", realName)
	params.Add("department", department)
	params.Add("timestamp", fmt.Sprintf("%d", timestamp))
	params.Add("sign", sign)
	if token != "" {
		params.Add("token", token)
	}

	fullURL := fmt.Sprintf("%s/sso/login?%s", baseURL, params.Encode())
	fmt.Printf("\n正在请求: %s\n", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
}

func main() {
	baseURL := flag.String("url", "http://localhost:8888", "GVA 服务器地址")
	provider := flag.String("provider", "provider1", "SSO 提供商名称")
	username := flag.String("username", "darcy", "用户名")
	realName := flag.String("realname", "测试用户", "真实姓名")
	department := flag.String("dept", "技术部", "部门")
	token := flag.String("token", "", "SSO token (可选)")
	secretKey := flag.String("secret", "your-secret-key-1", "签名密钥")
	timestamp := flag.Int64("ts", 0, "时间戳 (0表示使用当前时间)")

	flag.Parse()

	if *timestamp == 0 {
		*timestamp = time.Now().Unix()
	}

	sign := generateSign(*provider, *username, *realName, *department, *timestamp, *secretKey)

	fmt.Println("=== SSO 测试工具 ===")
	fmt.Printf("提供商: %s\n", *provider)
	fmt.Printf("用户名: %s\n", *username)
	fmt.Printf("真实姓名: %s\n", *department)
	fmt.Printf("部门: %s\n", *department)
	fmt.Printf("时间戳: %d\n", *timestamp)
	fmt.Printf("签名: %s\n", sign)

	if *token != "" {
		fmt.Printf("SSO Token: %s\n", *token)
	}

	loginTest(*baseURL, *provider, *username, *realName, *department, sign, *token, *timestamp)
}
