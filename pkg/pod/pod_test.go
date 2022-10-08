package pod

import (
	"testing"

	"github.com/shreyasHpandya/podctl/pkg/client"
)

func TestList(t *testing.T) {
	client := client.GetClient()
	//TODO: create test namespace and install few pods, and check if List returns correct results
	List(client, "", "")
}
