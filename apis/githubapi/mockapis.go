package githubapi

import (
	"encoding/json"
	"fmt"
)

var mockTreeUrlGitGo = `{"sha":"e03c6e27ef88e1bd30902b7a6258f0a54e01c2c7","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/e03c6e27ef88e1bd30902b7a6258f0a54e01c2c7","tree":[{"path":".vscode","mode":"040000","type":"tree","sha":"d58bd8e47da3a8fb24ffd678916563e1cdf1ebaa","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/d58bd8e47da3a8fb24ffd678916563e1cdf1ebaa"},{"path":".vscode/settings.json","mode":"100644","type":"blob","sha":"a7e16169143195e2a711f051bdbfcfb08088db2e","size":54,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/a7e16169143195e2a711f051bdbfcfb08088db2e"},{"path":"LICENSE","mode":"100644","type":"blob","sha":"261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64","size":11357,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64"},{"path":"README.md","mode":"100644","type":"blob","sha":"ea1a34a54f88f528e6864bc7cbf32cf2d794b964","size":1033,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/ea1a34a54f88f528e6864bc7cbf32cf2d794b964"},{"path":"apis","mode":"040000","type":"tree","sha":"702eeb95f9ed5b124ce84a96a9205f49b7c7379c","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/702eeb95f9ed5b124ce84a96a9205f49b7c7379c"},{"path":"apis/githubapi","mode":"040000","type":"tree","sha":"eebffec78929c432e2c7f63d9c53e5ef3de9d54c","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/eebffec78929c432e2c7f63d9c53e5ef3de9d54c"},{"path":"apis/githubapi/apis.go","mode":"100644","type":"blob","sha":"ba744632d38f23a3e8813ea1a9ab55c58f2ca450","size":1360,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/ba744632d38f23a3e8813ea1a9ab55c58f2ca450"},{"path":"apis/githubapi/datastructures.go","mode":"100644","type":"blob","sha":"57286350dd1cc21f9074c566c685046005299d35","size":256,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/57286350dd1cc21f9074c566c685046005299d35"},{"path":"apis/githubapi/methods.go","mode":"100644","type":"blob","sha":"c71f7fe3f62757473aef30a7613fbb2052bbe7ef","size":426,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/c71f7fe3f62757473aef30a7613fbb2052bbe7ef"},{"path":"apis/http.go","mode":"100644","type":"blob","sha":"e65dcd117797fd02fa7b02a7d4813a15cc6e7289","size":1869,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/e65dcd117797fd02fa7b02a7d4813a15cc6e7289"},{"path":"files","mode":"040000","type":"tree","sha":"57c59318121a8918f6ff259cdd3306c249215fdc","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/57c59318121a8918f6ff259cdd3306c249215fdc"},{"path":"files/file0.json","mode":"100644","type":"blob","sha":"ae85c07ec156f563e2d4184ce292219321f415ec","size":17,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/ae85c07ec156f563e2d4184ce292219321f415ec"},{"path":"files/file0.text","mode":"100644","type":"blob","sha":"e25bb064ec2d827c3a4a3a93797642e6d9b9b1f3","size":10,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/e25bb064ec2d827c3a4a3a93797642e6d9b9b1f3"},{"path":"files/file0.yaml","mode":"100644","type":"blob","sha":"31358e0ec5b5136961e0795962cdd1a0b51307ce","size":11,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/31358e0ec5b5136961e0795962cdd1a0b51307ce"},{"path":"files/file1.json","mode":"100644","type":"blob","sha":"41a0cb5ad5cadd8d12df5c8f493a41f370cab03d","size":17,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/41a0cb5ad5cadd8d12df5c8f493a41f370cab03d"},{"path":"files/file1.yaml","mode":"100644","type":"blob","sha":"cda23e65facdc68d2ea60a94715053085bacf97f","size":11,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/cda23e65facdc68d2ea60a94715053085bacf97f"},{"path":"files/file2.yaml","mode":"100644","type":"blob","sha":"177289fb9bc6aa7b19336111e902f92c8b4d28e7","size":11,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/177289fb9bc6aa7b19336111e902f92c8b4d28e7"},{"path":"githubparser","mode":"040000","type":"tree","sha":"7ca01c978d838116cca40c2fbe2a3743219abd43","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/7ca01c978d838116cca40c2fbe2a3743219abd43"},{"path":"githubparser/v1","mode":"040000","type":"tree","sha":"f7498c41927bb478faa94294ced8ad0eb0022448","url":"https://api.github.com/repos/armosec/go-git-url/git/trees/f7498c41927bb478faa94294ced8ad0eb0022448"},{"path":"githubparser/v1/datastructures.go","mode":"100644","type":"blob","sha":"3567877260d4cdad12eeb9facac72e5e4597348b","size":250,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/3567877260d4cdad12eeb9facac72e5e4597348b"},{"path":"githubparser/v1/parser.go","mode":"100644","type":"blob","sha":"bafc7b5cba10cf094ebdd054158b3cfdaa4389cc","size":3997,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/bafc7b5cba10cf094ebdd054158b3cfdaa4389cc"},{"path":"go.mod","mode":"100644","type":"blob","sha":"f5832d5bdf68408a5a8e21f51f6f3d1bf5b8e956","size":103,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/f5832d5bdf68408a5a8e21f51f6f3d1bf5b8e956"},{"path":"go.sum","mode":"100644","type":"blob","sha":"207ee15dca443269452c6c96d4b4df62c48e7e5e","size":974,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/207ee15dca443269452c6c96d4b4df62c48e7e5e"},{"path":"init.go","mode":"100644","type":"blob","sha":"6b0a179812d9df99602f2563bcd759af1920b809","size":638,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/6b0a179812d9df99602f2563bcd759af1920b809"},{"path":"interface.go","mode":"100644","type":"blob","sha":"9ccf2806ef1f897d0ebf5278b3353ca25d7d1288","size":401,"url":"https://api.github.com/repos/armosec/go-git-url/git/blobs/9ccf2806ef1f897d0ebf5278b3353ca25d7d1288"}],"truncated":false}`

type MockGitHubAPI struct {
}

func NewMockGitHubAPI() *MockGitHubAPI { return &MockGitHubAPI{} }

func (gh *MockGitHubAPI) GetRepoTree(owner, repo, branch string, headres *Headres) (*Tree, error) {
	t := &Tree{}
	switch fmt.Sprintf("%s/%s", owner, repo) {
	case "armosec/go-git-url":
		json.Unmarshal([]byte(mockTreeUrlGitGo), t)
	}
	return t, nil
}

func (gh MockGitHubAPI) GetDefaultBranchName(owner, repo string, headres *Headres) (string, error) {
	return "master", nil
}