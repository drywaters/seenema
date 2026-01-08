package assets

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

const shortHashLength = 12

// Version returns a short hash representing the contents of the provided asset files.
func Version(paths ...string) (string, error) {
	if len(paths) == 0 {
		return "", fmt.Errorf("no asset paths provided")
	}

	hasher := sha256.New()
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return "", fmt.Errorf("open %s: %w", path, err)
		}

		if _, err := io.Copy(hasher, file); err != nil {
			_ = file.Close()
			return "", fmt.Errorf("read %s: %w", path, err)
		}

		if err := file.Close(); err != nil {
			return "", fmt.Errorf("close %s: %w", path, err)
		}
	}

	sum := hex.EncodeToString(hasher.Sum(nil))
	if len(sum) > shortHashLength {
		sum = sum[:shortHashLength]
	}

	return sum, nil
}
