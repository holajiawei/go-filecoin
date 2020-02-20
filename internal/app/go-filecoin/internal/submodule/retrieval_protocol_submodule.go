package submodule

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/piecestore"
	iface "github.com/filecoin-project/go-fil-markets/retrievalmarket"
	impl "github.com/filecoin-project/go-fil-markets/retrievalmarket/impl"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket/network"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/libp2p/go-libp2p-core/host"

	retmkt "github.com/filecoin-project/go-filecoin/internal/app/go-filecoin/retrieval_market_connector"
)

// RetrievalProtocolSubmodule enhances the node with retrieval protocol
// capabilities.
type RetrievalProtocolSubmodule struct {
	RetrievalClient   iface.RetrievalClient
	RetrievalProvider iface.RetrievalProvider
}

// NewRetrievalProtocolSubmodule creates a new retrieval protocol submodule.
func NewRetrievalProtocolSubmodule(
	bs blockstore.Blockstore,
	c *ChainSubmodule,
	mw retmkt.MsgWaiter,
	host host.Host,
	providerAddr address.Address,
	ob retmkt.MsgSender,
	ps piecestore.PieceStore,
	signer retmkt.RetrievalSigner,
	wal retmkt.WalletAPI,
) (*RetrievalProtocolSubmodule, error) {
	panic("TODO: go-fil-markets integration")

	netwk := network.NewFromLibp2pHost(host)
	pnode := retmkt.NewRetrievalProviderNodeConnector(netwk, ps, bs)
	cnode := retmkt.NewRetrievalClientConnector(bs,
		c.ChainReader,
		mw,
		ob,
		ps,
		signer,
		wal,
	)
	rsvlr := retmkt.NewRetrievalPeerResolverConnector()

	return &RetrievalProtocolSubmodule{
		RetrievalClient:   impl.NewClient(netwk, bs, cnode, rsvlr),
		RetrievalProvider: impl.NewProvider(providerAddr, pnode, netwk, ps, bs),
	}, nil
}
