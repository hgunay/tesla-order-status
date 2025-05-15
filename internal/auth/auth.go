// Author: Hakan Gunay
// Date: 2025-05-15

package auth

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"tesla-order-status/internal/order"
)

const (
	ClientID            = "ownerapi"
	RedirectURI         = "https://auth.tesla.com/void/callback"
	AuthURL             = "https://auth.tesla.com/oauth2/v3/authorize"
	TokenURL            = "https://auth.tesla.com/oauth2/v3/token"
	Scope               = "openid email offline_access"
	CodeChallengeMethod = "S256"
)

func GenerateCodeVerifierAndChallenge() (string, string) {
	verifierBytes := make([]byte, 32)
	_, err := rand.Read(verifierBytes)
	if err != nil {
		panic(err)
	}
	codeVerifier := base64.RawURLEncoding.EncodeToString(verifierBytes)
	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.RawURLEncoding.EncodeToString(hash[:])
	return codeVerifier, codeChallenge
}

func GetAuthCode(codeChallenge string) string {
	state := "state"
	params := url.Values{}
	params.Add("client_id", ClientID)
	params.Add("redirect_uri", RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", Scope)
	params.Add("state", state)
	params.Add("code_challenge", codeChallenge)
	params.Add("code_challenge_method", CodeChallengeMethod)

	fullAuthURL := fmt.Sprintf("%s?%s", AuthURL, params.Encode())
	fmt.Println("Tarayıcıda oturum açın ve yönlendirme URL'sini buraya yapıştırın:")
	_ = exec.Command("open", fullAuthURL).Start()
	fmt.Println(fullAuthURL)

	fmt.Print("\nYönlendirme URL'sini girin: ")
	reader := bufio.NewReader(os.Stdin)
	redirectedURL, _ := reader.ReadString('\n')
	redirectedURL = strings.TrimSpace(redirectedURL)

	parsedURL, err := url.Parse(redirectedURL)
	if err != nil {
		panic("URL parse edilemedi")
	}
	return parsedURL.Query().Get("code")
}

func ExchangeCodeForTokens(authCode, codeVerifier string) order.Tokens {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", ClientID)
	data.Set("code", authCode)
	data.Set("redirect_uri", RedirectURI)
	data.Set("code_verifier", codeVerifier)

	resp, err := http.PostForm(TokenURL, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var tokens order.Tokens
	_ = json.Unmarshal(body, &tokens)
	tokens.CreatedAt = time.Now().Unix()
	return tokens
}

func IsTokenValid(token string) bool {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return false
	}
	decoded, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(decoded, &payload); err != nil {
		return false
	}
	if exp, ok := payload["exp"].(float64); ok {
		return int64(exp) > time.Now().Unix()
	}
	return false
}

func RefreshTokens(refreshToken string) order.Tokens {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", ClientID)
	data.Set("refresh_token", refreshToken)
	resp, err := http.PostForm(TokenURL, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var tokens order.Tokens
	_ = json.Unmarshal(body, &tokens)
	tokens.CreatedAt = time.Now().Unix()
	return tokens
}

func LoadTokensFromFile(tokenPath string) (order.Tokens, error) {
	file, err := os.Open(tokenPath)
	if err != nil {
		return order.Tokens{}, err
	}
	defer file.Close()
	var tokens order.Tokens
	err = json.NewDecoder(file).Decode(&tokens)
	return tokens, err
}

func SaveTokensToFile(tokens order.Tokens, tokenPath string) {
	file, _ := os.Create(tokenPath)
	defer file.Close()
	_ = json.NewEncoder(file).Encode(tokens)
}
