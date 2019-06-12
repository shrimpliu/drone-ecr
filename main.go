package main

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var (
		repo   = os.Getenv("PLUGIN_REPO")
		key    = os.Getenv("PLUGIN_ACCESS_KEY")
		secret = os.Getenv("PLUGIN_SECRET_KEY")
		region = os.Getenv("PLUGIN_REGION")
	)

	os.Setenv("AWS_ACCESS_KEY_ID", key)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secret)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		log.Fatalf("create aws session failed: %v", err)
	}

	svc := ecr.New(sess)
	username, password, registry, err := ecrGetLogin(svc)

	if err != nil {
		log.Fatalf("ecr get login failed: %v", err)
	}

	if !strings.HasPrefix(repo, registry) {
		repo = fmt.Sprintf("%s/%s", registry, repo)
	}

	os.Setenv("PLUGIN_REPO", repo)
	os.Setenv("PLUGIN_REGISTRY", registry)
	os.Setenv("DOCKER_USERNAME", username)
	os.Setenv("DOCKER_PASSWORD", password)

	cmd := exec.Command("drone-docker")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		log.Fatalf("push docker failed: %v", err)
	}
}

func ecrGetLogin(svc *ecr.ECR) (username, password, registry string, err error) {
	var result *ecr.GetAuthorizationTokenOutput
	var decoded []byte

	result, err = svc.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return
	}

	auth := result.AuthorizationData[0]
	token := *auth.AuthorizationToken
	decoded, err = base64.StdEncoding.DecodeString(token)
	if err != nil {
		return
	}

	registry = strings.TrimPrefix(*auth.ProxyEndpoint, "https://")
	creds := strings.Split(string(decoded), ":")
	username = creds[0]
	password = creds[1]
	return
}
