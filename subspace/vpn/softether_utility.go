package vpn

import (
	"os/exec"
	"fmt"
	"strings"
	"unicode/utf16"
	"encoding/base64"
	md4P "gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/crypto/md4"
)

//space to $20, $ to $24
func GenerateSoftetherEncodedString(data string) (string) {
	data = strings.Replace(data, "$", "$24", -1)
	data = strings.Replace(data, " ", "$20", -1)
	return data
}

// Return base64 encoded sha0 string
func GenerateSoftetherPasswordHash(data string) (string, error) {
	data = strings.Replace(data, `"`, `\"`, -1)
	cmd := fmt.Sprintf(`echo -n "%s" | openssl sha -binary | openssl base64`, data)
	out, err := exec.Command("bash","-c",cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd), err
	}
	return strings.TrimSpace(string(out)), nil
}

// Return base64 encoded NtLm v1 hash
func GenerateNtLmPasswordHash(data string) string {
	return base64.StdEncoding.EncodeToString(ntowfv1(data))
}


/**
Copy from https://github.com/ThomsonReutersEikon/go-ntlm
Because the ntowfv1 func is internal.
*/

/********************************
 NTLM V1 Password hash functions
*********************************/
func ntowfv1(passwd string) []byte {
	return md4(utf16FromString(passwd))
}

func md4(data []byte) []byte {
	md4 := md4P.New()
	md4.Write(data)
	return md4.Sum(nil)
}

func utf16FromString(s string) []byte {
	encoded := utf16.Encode([]rune(s))
	// TODO: I'm sure there is an easier way to do the conversion from utf16 to bytes
	result := zeroBytes(len(encoded) * 2)
	for i := 0; i < len(encoded); i++ {
		result[i*2] = byte(encoded[i])
		result[i*2+1] = byte(encoded[i] << 8)
	}
	return result
}

// Create a 0 initialized slice of bytes
func zeroBytes(length int) []byte {
	return make([]byte, length, length)
}