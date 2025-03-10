package provider

import (
	"context"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/superfly/fly-go"
)

// TODO: Add annotations
type Certificate struct{}

var (
	_ infer.CustomResource[CertificateArgs, CertificateState] = (*Certificate)(nil)
	_ infer.CustomRead[CertificateArgs, CertificateState]     = (*Certificate)(nil)
)

type CertificateArgs struct {
	App      string `json:"app" pulumi:"app"`
	Hostname string `json:"hostname" pulumi:"hostname"`
}

type CertificateStateChecks struct {
	ARecords              []string `json:"aRecords" pulumi:"aRecords"`
	AAAARecords           []string `json:"aaaaRecords" pulumi:"aaaaRecords"`
	CNAMERecords          []string `json:"cnameRecords" pulumi:"cnameRecords"`
	SOA                   string   `json:"soa" pulumi:"soa"`
	DNSProvider           string   `json:"dnsProvider" pulumi:"dnsProvider"`
	DNSVerificationRecord string   `json:"dnsVerificationRecord" pulumi:"dnsVerificationRecord"`
	ResolvedAddresses     []string `json:"resolvedAddresses" pulumi:"resolvedAddresses"`
}

type CertificateState struct {
	Input                     CertificateArgs        `pulumi:"input"`
	App                       string                 `json:"app" pulumi:"app"`
	Hostname                  string                 `json:"hostname" pulumi:"hostname"`
	CreatedAt                 time.Time              `json:"createdAt" pulumi:"createdAt"`
	ID                        string                 `json:"flyId" pulumi:"flyId"`
	AcmeDNSConfigured         bool                   `json:"acmeDnsConfigured" pulumi:"acmeDnsConfigured"`
	AcmeALPNConfigured        bool                   `json:"acmeAlpnConfigured" pulumi:"acmeAlpnConfigured"`
	Configured                bool                   `json:"configured" pulumi:"configured"`
	CertificateAuthority      string                 `json:"certificateAuthority" pulumi:"certificateAuthority"`
	DNSProvider               string                 `json:"dnsProvider" pulumi:"dnsProvider"`
	DNSValidationInstructions string                 `json:"dnsValidationInstructions" pulumi:"dnsValidationInstructions"`
	DNSValidationHostname     string                 `json:"dnsValidationHostname" pulumi:"dnsValidationHostname"`
	DNSValidationTarget       string                 `json:"dnsValidationTarget" pulumi:"dnsValidationTarget"`
	Source                    string                 `json:"source" pulumi:"source"`
	ClientStatus              string                 `json:"clientStatus" pulumi:"clientStatus"`
	IsApex                    bool                   `json:"isApex" pulumi:"isApex"`
	IsWildcard                bool                   `json:"isWildcard" pulumi:"isWildcard"`
	Checks                    CertificateStateChecks `json:"checks" pulumi:"checks"`
}

func (Certificate) Create(ctx context.Context, name string, input CertificateArgs, preview bool) (string, CertificateState, error) {
	if preview {
		state := CertificateState{
			Input:     input,
			Hostname:  input.Hostname,
			App:       input.App,
			CreatedAt: time.Now(),
		}

		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	certificate, checks, err := cfg.flyClient.AddCertificate(ctx, input.App, input.Hostname)
	if err != nil {
		return name, CertificateState{}, err
	}
	state := buildCerticateState(input, certificate, checks)

	return name, state, nil
}

func (Certificate) Read(ctx context.Context, id string, input CertificateArgs, state CertificateState) (
	string, CertificateArgs, CertificateState, error,
) {
	cfg := infer.GetConfig[Config](ctx)

	certificate, checks, err := cfg.flyClient.CheckAppCertificate(ctx, state.Input.App, state.Input.Hostname)
	if err != nil {
		return id, input, state, err
	}

	state = buildCerticateState(input, certificate, checks)

	return id, input, state, nil
}

func (Certificate) Delete(ctx context.Context, reqID string, state CertificateState) error {
	cfg := infer.GetConfig[Config](ctx)

	p.GetLogger(ctx).Infof("Deleting Certificate %s from %s", state.Hostname, state.App)

	_, err := cfg.flyClient.DeleteCertificate(ctx, state.Input.App, state.Input.Hostname)

	return err
}

var certificateDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{},
	DeleteBeforeReplaceProps: []string{"App", "Hostname"},
}

func (Certificate) Diff(ctx context.Context, id string, state CertificateState, input CertificateArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, certificateDiffOpts)
}

func buildCerticateState(input CertificateArgs, certificate *fly.AppCertificate, checks *fly.HostnameCheck) CertificateState {
	return CertificateState{
		Input:                     input,
		App:                       input.App,
		Hostname:                  input.Hostname,
		CreatedAt:                 certificate.CreatedAt,
		ID:                        certificate.ID,
		AcmeDNSConfigured:         certificate.AcmeDNSConfigured,
		AcmeALPNConfigured:        certificate.AcmeALPNConfigured,
		Configured:                certificate.Configured,
		CertificateAuthority:      certificate.CertificateAuthority,
		DNSProvider:               certificate.DNSProvider,
		DNSValidationInstructions: certificate.DNSValidationInstructions,
		DNSValidationHostname:     certificate.DNSValidationHostname,
		DNSValidationTarget:       certificate.DNSValidationTarget,
		Source:                    certificate.Source,
		ClientStatus:              certificate.ClientStatus,
		IsApex:                    certificate.IsApex,
		IsWildcard:                certificate.IsWildcard,
		Checks: CertificateStateChecks{
			ARecords:              checks.ARecords,
			AAAARecords:           checks.AAAARecords,
			CNAMERecords:          checks.CNAMERecords,
			SOA:                   checks.SOA,
			DNSProvider:           checks.DNSProvider,
			DNSVerificationRecord: checks.DNSVerificationRecord,
			ResolvedAddresses:     checks.ResolvedAddresses,
		},
	}
}
