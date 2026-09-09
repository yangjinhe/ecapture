// Copyright 2022 CFC4N <cfc4n.cs@gmail.com>. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package writers

// KeylogWriter writes keylog to a standalone file.
type KeylogWriter struct {
	*FileWriter
}

func (w *KeylogWriter) Name() string {
	return "keylog_writer"
}

func (w *KeylogWriter) Flush() error {
	return w.FileWriter.Flush()
}

func NewKeylogWriter(fw *FileWriter) *KeylogWriter {
	return &KeylogWriter{
		FileWriter: fw,
	}
}

func (w *KeylogWriter) Write(p []byte) (n int, err error) {
	// Create a copy to avoid modifying the provided buffer
	data := make([]byte, len(p)+1)
	copy(data, p)
	data[len(p)] = '\n'
	return w.FileWriter.Write(data)
}
