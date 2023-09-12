package boot

import "os"

type CType string

const (
	Loc = CType("local")
	Dev = CType("dev")
	Pro = CType("prod")

	envKey = "boot"
)

func SetEnv(e CType) error {
	return os.Setenv(envKey, string(e))
}

func GetEnv() string {
	return os.Getenv(envKey)
}
