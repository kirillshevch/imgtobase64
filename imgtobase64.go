package imgtobase64

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func Imgtobase64(name string) (string, error) {
	file, err := os.Open(name)

	if err != nil {
		return "", fmt.Errorf("imagebase64: Can't open source file: %s", err)
	}

	defer file.Close()

	fInfo, err := file.Stat()

	if err != nil {
		return "", fmt.Errorf("imagebase64: Can't get source file info: %s", err)
	}

	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(file)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64Str, nil
}
