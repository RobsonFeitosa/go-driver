package bucket

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/RobsonFeitosa/go-driver/internal/bucket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para o serviço S3
type mockS3 struct {
	mock.Mock
}

func (m *mockS3) Download(dst string, src string) (*os.File, error) {
	args := m.Called(dst, src)
	return args.Get(0).(*os.File), args.Error(1)
}

func (m *mockS3) Upload(file []byte, key string) error {
	args := m.Called(file, key)
	return args.Error(0)
}

func (m *mockS3) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestDownloadAndUpload(t *testing.T) {
	// Configuração do Mock
	mockS3 := &mockS3{}
	sess := bucket.New(mockS3)

	// Simulação do Download
	fileContent := []byte("conteúdo do arquivo")
	mockS3.On("Download", "/tmp/teste.txt", "src/teste.txt").Return(ioutil.NopCloser(bytes.NewReader(fileContent)), nil)

	// Teste de Download
	file, err := sess.Download("src/teste.txt", "/tmp/teste.txt")
	assert.NoError(t, err)
	assert.NotNil(t, file)

	// Simulação do Upload
	mockS3.On("Upload", fileContent, "dst/teste.txt").Return(nil)

	// Teste de Upload
	err = sess.Upload(bytes.NewReader(fileContent), "dst/teste.txt")
	assert.NoError(t, err)

	// Verificação de que os métodos do Mock foram chamados corretamente
	mockS3.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	// Configuração do Mock
	mockS3 := &mockS3{}
	sess := bucket.NewAwsSession(mockS3)

	// Simulação da deleção
	mockS3.On("Delete", "file/to/delete.txt").Return(nil)

	// Teste de Delete
	err := sess.Delete("file/to/delete.txt")
	assert.NoError(t, err)

	// Verificação de que o método do Mock foi chamado corretamente
	mockS3.AssertExpectations(t)
}
