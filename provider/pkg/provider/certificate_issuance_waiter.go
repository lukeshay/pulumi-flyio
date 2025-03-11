package provider

import (
	"context"
	"fmt"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/superfly/fly-go"
)

type CertificateIssuanceWaiter struct{}

var _ infer.Annotated = (*CertificateIssuanceWaiter)(nil)

func (c *CertificateIssuanceWaiter) Annotate(a infer.Annotator) {
	a.Describe(&c, "A component that waits for a Fly.io SSL/TLS certificate to be fully issued.")
}

var (
	_ infer.CustomResource[CertificateIssuanceWaiterArgs, CertificateIssuanceWaiterState] = (*CertificateIssuanceWaiter)(nil)
	_ infer.CustomDelete[CertificateIssuanceWaiterState]                                  = (*CertificateIssuanceWaiter)(nil)
	_ infer.CustomCreate[CertificateIssuanceWaiterArgs, CertificateIssuanceWaiterState]   = (*CertificateIssuanceWaiter)(nil)
)

type CertificateIssuanceWaiterArgs struct {
	App      string        `json:"app" pulumi:"app"`
	Hostname string        `json:"hostname" pulumi:"hostname"`
	Timeout  time.Duration `json:"timeout" pulumi:"timeout"`
}

var _ infer.Annotated = (*CertificateIssuanceWaiterArgs)(nil)

func (a *CertificateIssuanceWaiterArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.App, "The name of the Fly app that the certificate is associated with.")
	anno.Describe(&a.Hostname, "The hostname for the certificate (e.g., example.com).")
	anno.Describe(&a.Timeout, "The maximum time to wait for the certificate to be fully issued. Formatted like 5s, 5m, etc.")
}

type CertificateIssuanceWaiterState struct {
	Input                CertificateIssuanceWaiterArgs `pulumi:"input"`
	App                  string                        `json:"app" pulumi:"app"`
	Hostname             string                        `json:"hostname" pulumi:"hostname"`
	Timeout              time.Duration                 `json:"timeout" pulumi:"timeout"`
	CertificateID        string                        `json:"certificateId" pulumi:"certificateId"`
	IssuedAt             time.Time                     `json:"issuedAt" pulumi:"issuedAt"`
	ClientStatus         string                        `json:"clientStatus" pulumi:"clientStatus"`
	CertificateAuthority string                        `json:"certificateAuthority" pulumi:"certificateAuthority"`
	ECDSAExpiresAt       *time.Time                    `json:"ecdsaExpiresAt,omitempty" pulumi:"ecdsaExpiresAt,optional"`
	RSAExpiresAt         *time.Time                    `json:"rsaExpiresAt,omitempty" pulumi:"rsaExpiresAt,optional"`
	IsFullyIssued        bool                          `json:"isFullyIssued" pulumi:"isFullyIssued"`
}

var _ infer.Annotated = (*CertificateIssuanceWaiterState)(nil)

func (s *CertificateIssuanceWaiterState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.Input, "The input arguments used for the certificate issuance waiter.")
	anno.Describe(&s.App, "The name of the Fly app.")
	anno.Describe(&s.Hostname, "The hostname for the certificate.")
	anno.Describe(&s.Timeout, "The timeout duration used for waiting.")
	anno.Describe(&s.CertificateID, "The Fly.io certificate ID.")
	anno.Describe(&s.IssuedAt, "When the certificate was fully issued.")
	anno.Describe(&s.ClientStatus, "The status of the certificate.")
	anno.Describe(&s.CertificateAuthority, "The certificate authority used.")
	anno.Describe(&s.ECDSAExpiresAt, "Expiration time for the ECDSA certificate.")
	anno.Describe(&s.RSAExpiresAt, "Expiration time for the RSA certificate.")
	anno.Describe(&s.IsFullyIssued, "Whether the certificate is fully issued (has both ECDSA and RSA certificates).")
}

func (CertificateIssuanceWaiter) Create(ctx context.Context, name string, input CertificateIssuanceWaiterArgs, preview bool) (string, CertificateIssuanceWaiterState, error) {
	state := CertificateIssuanceWaiterState{
		Input:    input,
		App:      input.App,
		Hostname: input.Hostname,
		Timeout:  input.Timeout,
	}

	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	// First check if the certificate exists
	certificate, _, err := cfg.flyClient.CheckAppCertificate(ctx, input.App, input.Hostname)
	if err != nil {
		// If the certificate doesn't exist, we need to create it
		p.GetLogger(ctx).Infof("Certificate for %s does not exist, creating it", input.Hostname)
		certificate, _, err = cfg.flyClient.AddCertificate(ctx, input.App, input.Hostname)
		if err != nil {
			return name, state, fmt.Errorf("failed to create certificate: %w", err)
		}
	}

	state.CertificateID = certificate.ID
	state.ClientStatus = certificate.ClientStatus
	state.CertificateAuthority = certificate.CertificateAuthority

	// Check if the certificate is already fully issued
	if len(certificate.Issued.Nodes) >= 2 {
		p.GetLogger(ctx).Infof("Certificate for %s is already fully issued", input.Hostname)
		state.IsFullyIssued = true
		state.IssuedAt = time.Now()
		updateCertificateExpirations(&state, certificate)
		return name, state, nil
	}

	// Wait for the certificate to be fully issued
	timeout := time.After(input.Timeout * time.Second)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return name, state, fmt.Errorf("timed out waiting for certificate to be issued after %d seconds", int(input.Timeout))
		case <-ticker.C:
			certificate, _, err = cfg.flyClient.CheckAppCertificate(ctx, input.App, input.Hostname)
			if err != nil {
				p.GetLogger(ctx).Warningf("Error checking certificate status: %v", err)
				continue
			}

			state.ClientStatus = certificate.ClientStatus

			// Check if we have both ECDSA and RSA certificates (length of 2)
			if len(certificate.Issued.Nodes) >= 2 {
				p.GetLogger(ctx).Infof("Certificate for %s is fully issued", input.Hostname)
				state.IsFullyIssued = true
				state.IssuedAt = time.Now()
				updateCertificateExpirations(&state, certificate)
				return name, state, nil
			}

			p.GetLogger(ctx).Infof("Waiting for certificate to be fully issued, current status: %s, issued nodes: %d",
				certificate.ClientStatus, len(certificate.Issued.Nodes))
		}
	}
}

func (CertificateIssuanceWaiter) Delete(ctx context.Context, id string, state CertificateIssuanceWaiterState) error {
	// This resource only waits for certificate issuance but doesn't own the certificate itself,
	// so we don't delete the actual certificate when this resource is deleted.
	return nil
}

var certificateIssuanceWaiterDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{"Timeout"},
	DeleteBeforeReplaceProps: []string{"App", "Hostname"},
}

func (CertificateIssuanceWaiter) Diff(ctx context.Context, id string, state CertificateIssuanceWaiterState, input CertificateIssuanceWaiterArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, certificateIssuanceWaiterDiffOpts)
}

func updateCertificateExpirations(state *CertificateIssuanceWaiterState, certificate *fly.AppCertificate) {
	for _, issued := range certificate.Issued.Nodes {
		switch issued.Type {
		case "ecdsa":
			state.ECDSAExpiresAt = &issued.ExpiresAt
		case "rsa":
			state.RSAExpiresAt = &issued.ExpiresAt
		}
	}
}
