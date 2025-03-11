package provider

import (
	"context"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/superfly/fly-go"
)

type Certificate struct{}

var _ infer.Annotated = (*Certificate)(nil)

func (c *Certificate) Annotate(a infer.Annotator) {
	a.Describe(&c, "A Fly.io SSL/TLS certificate for an app's domain.")
}

var (
	_ infer.CustomResource[CertificateArgs, CertificateState] = (*Certificate)(nil)
	_ infer.CustomRead[CertificateArgs, CertificateState]     = (*Certificate)(nil)
)

type CertificateArgs struct {
	App      string `json:"app" pulumi:"app"`
	Hostname string `json:"hostname" pulumi:"hostname"`
}

var _ infer.Annotated = (*CertificateArgs)(nil)

func (a *CertificateArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.App, "The name of the Fly app to add the certificate to.")
	anno.Describe(&a.Hostname, "The hostname for the certificate (e.g., example.com).")
}

type CertificateState struct {
	Input                     CertificateArgs `pulumi:"input"`
	App                       string          `json:"app" pulumi:"app"`
	Hostname                  string          `json:"hostname" pulumi:"hostname"`
	CreatedAt                 time.Time       `json:"createdAt" pulumi:"createdAt"`
	ID                        string          `json:"flyId" pulumi:"flyId"`
	AcmeDNSConfigured         bool            `json:"acmeDnsConfigured" pulumi:"acmeDnsConfigured"`
	AcmeALPNConfigured        bool            `json:"acmeAlpnConfigured" pulumi:"acmeAlpnConfigured"`
	Configured                bool            `json:"configured" pulumi:"configured"`
	CertificateAuthority      string          `json:"certificateAuthority" pulumi:"certificateAuthority"`
	DNSProvider               string          `json:"dnsProvider" pulumi:"dnsProvider"`
	DNSValidationInstructions string          `json:"dnsValidationInstructions" pulumi:"dnsValidationInstructions"`
	DNSValidationHostname     string          `json:"dnsValidationHostname" pulumi:"dnsValidationHostname"`
	DNSValidationTarget       string          `json:"dnsValidationTarget" pulumi:"dnsValidationTarget"`
	Source                    string          `json:"source" pulumi:"source"`
	ClientStatus              string          `json:"clientStatus" pulumi:"clientStatus"`
	IsApex                    bool            `json:"isApex" pulumi:"isApex"`
	IsWildcard                bool            `json:"isWildcard" pulumi:"isWildcard"`
}

var _ infer.Annotated = (*CertificateState)(nil)

func (s *CertificateState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.Input, "The input arguments used to create the certificate.")
	anno.Describe(&s.App, "The name of the Fly app.")
	anno.Describe(&s.Hostname, "The hostname for the certificate.")
	anno.Describe(&s.CreatedAt, "When the certificate was created.")
	anno.Describe(&s.ID, "The Fly.io certificate ID.")
	anno.Describe(&s.AcmeDNSConfigured, "Whether ACME DNS verification is configured.")
	anno.Describe(&s.AcmeALPNConfigured, "Whether ACME ALPN verification is configured.")
	anno.Describe(&s.Configured, "Whether the certificate is fully configured.")
	anno.Describe(&s.CertificateAuthority, "The certificate authority used.")
	anno.Describe(&s.DNSProvider, "The DNS provider for the hostname.")
	anno.Describe(&s.DNSValidationInstructions, "Instructions for DNS validation.")
	anno.Describe(&s.DNSValidationHostname, "Hostname for DNS validation.")
	anno.Describe(&s.DNSValidationTarget, "Target for DNS validation.")
	anno.Describe(&s.Source, "The source of the certificate.")
	anno.Describe(&s.ClientStatus, "The status of the certificate.")
	anno.Describe(&s.IsApex, "Whether the hostname is an apex domain.")
	anno.Describe(&s.IsWildcard, "Whether the certificate is a wildcard certificate.")
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

	certificate, _, err := cfg.flyClient.AddCertificate(ctx, input.App, input.Hostname)
	if err != nil {
		return name, CertificateState{}, err
	}

	p.GetLogger(ctx).Infof("certificate created for %s", input.Hostname)
	state := buildCerticateState(input, certificate)
	return name, state, nil
}

func (Certificate) Read(ctx context.Context, id string, input CertificateArgs, state CertificateState) (
	string, CertificateArgs, CertificateState, error,
) {
	cfg := infer.GetConfig[Config](ctx)

	certificate, _, err := cfg.flyClient.CheckAppCertificate(ctx, state.App, state.Hostname)
	if err != nil {
		return id, input, state, err
	}

	state = buildCerticateState(input, certificate)

	return id, input, state, nil
}

func (Certificate) Delete(ctx context.Context, reqID string, state CertificateState) error {
	cfg := infer.GetConfig[Config](ctx)

	p.GetLogger(ctx).Infof("Deleting Certificate %s from %s", state.Hostname, state.App)

	_, err := cfg.flyClient.DeleteCertificate(ctx, state.App, state.Hostname)

	return err
}

var certificateDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{},
	DeleteBeforeReplaceProps: []string{"App", "Hostname"},
}

func (Certificate) Diff(ctx context.Context, id string, state CertificateState, input CertificateArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, certificateDiffOpts)
}

func buildCerticateState(input CertificateArgs, certificate *fly.AppCertificate) CertificateState {
	state := CertificateState{
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
	}

	return state
}
