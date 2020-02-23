package music_lib

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

func LoadMusicFile(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var buffer [][]byte
	var opusLen int16
	for {
		// Read opus frame length from dca file.
		if err := binary.Read(file, binary.LittleEndian, &opusLen); err != nil {
			// If this is the end of the file, just return.
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				if err := file.Close(); err != nil {
					return nil, err
				}
				return buffer, nil
			}
			if err != nil {
				return nil, errors.New("error reading from dca file: " + err.Error())
			}
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opusLen)

		// Should not be any end of file errors
		if err := binary.Read(file, binary.LittleEndian, &InBuf); err != nil {
			return nil, errors.New("error reading from dca file : " + err.Error())
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}
