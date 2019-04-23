package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift-metal3/kni-installer/pkg/asset"
	"github.com/openshift-metal3/kni-installer/pkg/asset/templates/content"
)

const (
	kubeSystemConfigmapRootCAFileName = "kube-system-configmap-root-ca.yaml.template"
)

var _ asset.WritableAsset = (*KubeSystemConfigmapRootCA)(nil)

// KubeSystemConfigmapRootCA is the constant to represent contents of kube-system-configmap-root-ca.yaml.template file.
type KubeSystemConfigmapRootCA struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *KubeSystemConfigmapRootCA) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *KubeSystemConfigmapRootCA) Name() string {
	return "KubeSystemConfigmapRootCA"
}

// Generate generates the actual files by this asset
func (t *KubeSystemConfigmapRootCA) Generate(parents asset.Parents) error {
	fileName := kubeSystemConfigmapRootCAFileName
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
func (t *KubeSystemConfigmapRootCA) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *KubeSystemConfigmapRootCA) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeSystemConfigmapRootCAFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
