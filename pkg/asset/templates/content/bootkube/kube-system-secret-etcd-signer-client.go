package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift-metal3/kni-installer/pkg/asset"
	"github.com/openshift-metal3/kni-installer/pkg/asset/templates/content"
)

const (
	kubeSystemSecretEtcdSignerClientFileName = "kube-system-secret-etcd-signer-client.yaml.template"
)

var _ asset.WritableAsset = (*KubeSystemSecretEtcdSignerClient)(nil)

// KubeSystemSecretEtcdSignerClient is the constant to represent contents of kube-system-secret-etcd-signer-client.yaml.template file.
type KubeSystemSecretEtcdSignerClient struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *KubeSystemSecretEtcdSignerClient) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *KubeSystemSecretEtcdSignerClient) Name() string {
	return "KubeSystemSecretEtcdSignerClient"
}

// Generate generates the actual files by this asset
func (t *KubeSystemSecretEtcdSignerClient) Generate(parents asset.Parents) error {
	fileName := kubeSystemSecretEtcdSignerClientFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *KubeSystemSecretEtcdSignerClient) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *KubeSystemSecretEtcdSignerClient) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeSystemSecretEtcdSignerClientFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
