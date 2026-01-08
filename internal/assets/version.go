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

		copyErr := func() error {
			_, err := io.Copy(hasher, file)
			return err
		}()
		closeErr := file.Close()
		if copyErr != nil {
			return "", fmt.Errorf("read %s: %w", path, copyErr)
		}
		if closeErr != nil {
			return "", fmt.Errorf("close %s: %w", path, closeErr)
		}
	}

	sum := hex.EncodeToString(hasher.Sum(nil))
	if len(sum) > shortHashLength {
		sum = sum[:shortHashLength]
	}

	return sum, nil
}
