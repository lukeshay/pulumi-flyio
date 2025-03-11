package provider

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"net"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"

	"github.com/pulumi/pulumi-go-provider/infer"
	"golang.org/x/crypto/curve25519"
)

type WireGuardPeer struct{}

var _ infer.Annotated = (*WireGuardPeer)(nil)

func (w *WireGuardPeer) Annotate(anno infer.Annotator) {
	anno.Describe(&w, "A Fly.io WireGuard peer for private network connectivity.")
}

var (
	_ infer.CustomResource[WireGuardPeerArgs, WireGuardPeerState] = (*WireGuardPeer)(nil)
	_ infer.CustomRead[WireGuardPeerArgs, WireGuardPeerState]     = (*WireGuardPeer)(nil)
	_ infer.CustomDelete[WireGuardPeerState]                      = (*WireGuardPeer)(nil)
)

type WireGuardPeerArgs struct {
	Org     string `pulumi:"org"`
	Region  string `pulumi:"region,optional"`
	Name    string `pulumi:"name,optional"`
	Network string `pulumi:"network,optional"`
}

var _ infer.Annotated = (*WireGuardPeerArgs)(nil)

func (a *WireGuardPeerArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.Org, "The organization to create the WireGuard peer in.")
	anno.Describe(&a.Region, "The region to create the WireGuard peer in.")
	anno.Describe(&a.Name, "The name of the WireGuard peer.")
	anno.Describe(&a.Network, "The network to create the WireGuard peer in.")
}

type WireGuardPeerState struct {
	PublicKey       string `pulumi:"publicKey" provider:"secret"`
	PrivateKey      string `pulumi:"privateKey" provider:"secret"`
	WireGuardConfig string `pulumi:"wireguardConfig" provider:"secret"`
	PeerIP          string `pulumi:"peerIp"`
	EndpointIP      string `pulumi:"endpointIp"`
	Name            string `pulumi:"name"`
	Org             string `pulumi:"org"`
	Region          string `pulumi:"region"`
	Network         string `pulumi:"network,optional"`
}

var _ infer.Annotated = (*WireGuardPeerState)(nil)

func (s *WireGuardPeerState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.PublicKey, "The public key of the WireGuard peer.")
	anno.Describe(&s.PrivateKey, "The private key of the WireGuard peer.")
	anno.Describe(&s.WireGuardConfig, "The WireGuard configuration for the peer.")
	anno.Describe(&s.PeerIP, "The IP address assigned to the WireGuard peer.")
	anno.Describe(&s.EndpointIP, "The endpoint IP address for the WireGuard peer.")
	anno.Describe(&s.Name, "The name of the WireGuard peer.")
	anno.Describe(&s.Org, "The organization the WireGuard peer belongs to.")
	anno.Describe(&s.Region, "The region the WireGuard peer is in.")
	anno.Describe(&s.Network, "The network the WireGuard peer belongs to.")
}

func (WireGuardPeer) Create(ctx context.Context, name string, input WireGuardPeerArgs, preview bool) (string, WireGuardPeerState, error) {
	state := WireGuardPeerState{
		Name:    input.Name,
		Org:     input.Org,
		Region:  input.Region,
		Network: input.Network,
	}

	if state.Name == "" {
		nameSlice := strings.Split(name, ":")
		state.Name = strings.ToLower(fmt.Sprintf("%s-%s", nameSlice[len(nameSlice)-1], randSeq(8)))

		p.GetLogger(ctx).Infof("Name is %s", state.Name)
	}

	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	org, err := cfg.flyClient.GetOrganizationBySlug(ctx, input.Org)
	if err != nil {
		return "", WireGuardPeerState{}, err
	}

	pubkey, privkey := c25519pair()

	peer, err := cfg.flyClient.CreateWireGuardPeer(ctx, org, state.Region, state.Name, pubkey, state.Network)
	if err != nil {
		return "", WireGuardPeerState{}, err
	}

	state.WireGuardConfig = generateWgConf(state.PrivateKey, state.PublicKey, peer.Peerip, peer.Endpointip)
	state.PublicKey = pubkey
	state.PrivateKey = privkey
	state.PeerIP = peer.Peerip
	state.EndpointIP = peer.Endpointip

	return name, state, nil
}

func (WireGuardPeer) Read(ctx context.Context, id string, inputs WireGuardPeerArgs, state WireGuardPeerState) (string, WireGuardPeerArgs, WireGuardPeerState, error) {
	cfg := infer.GetConfig[Config](ctx)

	peer, err := cfg.flyClient.GetWireGuardPeer(ctx, state.Org, state.Name)
	if err != nil {
		return id, inputs, state, err
	}

	state.WireGuardConfig = generateWgConf(state.PrivateKey, state.PublicKey, peer.Peerip, state.EndpointIP)
	state.PeerIP = peer.Peerip
	state.EndpointIP = peer.GatewayStatus.Endpoint

	return id, inputs, state, nil
}

func (WireGuardPeer) Delete(ctx context.Context, reqID string, state WireGuardPeerState) error {
	cfg := infer.GetConfig[Config](ctx)

	org, err := cfg.flyClient.GetOrganizationBySlug(ctx, state.Org)
	if err != nil {
		return err
	}

	return cfg.flyClient.RemoveWireGuardPeer(ctx, org, state.Name)
}

func c25519pair() (string, string) {
	var private [32]byte
	_, err := rand.Read(private[:])
	if err != nil {
		panic(fmt.Sprintf("reading from random: %s", err))
	}

	public, err := curve25519.X25519(private[:], curve25519.Basepoint)
	if err != nil {
		panic(fmt.Sprintf("can't mult: %s", err))
	}

	return base64.StdEncoding.EncodeToString(public),
		base64.StdEncoding.EncodeToString(private[:])
}

func generateWgConf(privkey string, pubKey string, peerip string, endpointip string) string {
	templateStr := `
[Interface]
PrivateKey = {{.Privkey}}
Address = {{.Peerip}}/120
DNS = {{.DNS}}

[Peer]
PublicKey = {{.Pubkey}}
AllowedIPs = {{.AllowedIPs}}
Endpoint = {{.Endpointip}}:51820
PersistentKeepalive = 15

`
	data := struct {
		Privkey    string
		Pubkey     string
		AllowedIPs string
		DNS        string
		Endpointip string
		Peerip     string
	}{
		Privkey:    privkey,
		Pubkey:     pubKey,
		Endpointip: endpointip,
		Peerip:     peerip,
	}

	addr := net.ParseIP(peerip).To16()
	for i := 6; i < 16; i++ {
		addr[i] = 0
	}

	// BUG(tqbf): can't stay this way
	data.AllowedIPs = fmt.Sprintf("%s/48", addr)

	addr[15] = 3

	data.DNS = addr.String()

	tmpl := template.Must(template.New("name").Parse(templateStr))

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	return buf.String()
}
