// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Roller network.
var MainnetBootnodes = []string{
	// Roller Foundation Go Bootnodes
	"enode://672887b53328555df3cb4dc107e8c434b3e8dacc6672aeb0df85a41808aaed1a0f0343b0ed5d65a25065fcb89d79a2059cb3232af4a4549e1324a0fa8b15b79f@108.61.161.232:30303", //seed1 (Japan)
    "enode://e982046189448702898a6e1d8d4041a91e7fe96abdafe16085773da4f1e3d0c349d2577285e5c09565075454994de5454542a0d9c5bfde8d45f5305d2e93e2e7@45.63.101.57:30303", //seed1 (London)
    "enode://a24907b86b481c09032f6f6fc0db3907c3f68fa3c59bdc4a45c2eaa8df637e42d056368b4140d16db2cca260ffcb1f4b8e785b2aba3bd5ba75b33d9a519cc6f5@95.179.157.139:30303", //seed1 (Amsterdam)
    "enode://403ce417251c990d70967523784a9d41f1e746d64f498a903eb485f9156b3c1d7d397126bb410f45d7742ae23101f3ae5d558bdb8948fba09d6d72005f1c93d1@125.212.245.168:30303", //seed1 (Vn)

	

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	//"enode://fb4a06943051c62802d01d2d7654d5fa4ab0d9ebb871c8ec8f819d9d6b6f3660f8df97908a16421c2ac2cbb85e5a7ab9fc189cbafe974b492a36ef05674780b5@192.168.100.3:30303", //west-us
	//"enode://b20575289b95024f98ccfbd809ce5148082bc819280394d6c0af085c8d7e8681784f515529811c7794eb54448bc9fb08dffe24f0ce2a824c9063bbcc220cfb9a@210.211.111.3:30303", //west-us
	//"enode://ecafd273c612f8743bc1e7d5cf0b512fab0a541be7c93f0fe14b15858e5cf2adbe391494945f101e18273cda23ba480c98a434eafbe6779240728bbd70295df0@210.211.111.12:30303", //west-us
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
