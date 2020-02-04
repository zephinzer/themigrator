package migration

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"
)

func New(uuid string, filepath string) (*Migration, error) {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	content := strings.
		ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						string(fileContent),
						"\n", " ",
					),
					"\r", "",
				),
				"\t", " ",
			),
			"  ", " ",
		)
	contentHash := fmt.Sprintf("%x", sha256.Sum256([]byte(content)))
	return &Migration{
		UUID:        uuid,
		Content:     content,
		ContentHash: contentHash,
	}, nil
}
