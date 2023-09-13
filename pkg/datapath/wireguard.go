// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package datapath

import (
	"net"

	"github.com/cilium/cilium/api/v1/models"
)

// WireguardAgent manages the Wireguard peers
type WireguardAgent interface {
	NodeHandler
	UpdatePeer(nodeName, pubKeyHex string, nodeIPv4, nodeIPv6 net.IP) error
	DeletePeer(nodeName string) error
	Status(includePeers bool) (*models.WireguardStatus, error)
}