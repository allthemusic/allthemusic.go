package index

import (
	"fmt"
	ipfsCmds "github.com/jbenet/go-ipfs/core/commands"
)

func AddMetadata(metadataBlob string) (key string, err error) {
	fmt.Printf("putCmd: %#v\n", ipfsCmds.ObjectCmd.Subcommand("put"))
	return
}
