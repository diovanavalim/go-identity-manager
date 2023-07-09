package anchorhandle

import (
	"fmt"
	"log"
	"os/exec"
)

type AnchorHandle interface {
	GenerateKey(keyLength int) string
	GenerateSeed() string
}

type Service struct{}

func (s Service) GenerateKey(keyLength int) string {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("openssl rand -base64 %d", keyLength))

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(out[:len(out)-1])
}

func (s Service) GenerateSeed() string {
	key := s.GenerateKey(32)

	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s\" | fold -w 32 | head -n 1", key))

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprintf("out %s", string(out)))
		log.Fatal(err.Error())
	}

	return string(out[:len(out)-1])
}

func (s Service) SeedAsBytes(seed string) []byte {
	return []byte(seed)
}

func (s Service) GenerateDID(seed string) {
	//seedAsBytes := s.SeedAsBytes(seed)
	//vk := byte(nacl.)
}
