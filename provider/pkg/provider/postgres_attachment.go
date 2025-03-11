package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/superfly/fly-go"
)

type PostgresAttachment struct{}

var _ infer.Annotated = (*PostgresAttachment)(nil)

func (p *PostgresAttachment) Annotate(anno infer.Annotator) {
	anno.Describe(&p, "A Fly.io Postgres attachment connects a Postgres database to an application.")
}

var (
	_ infer.CustomResource[PostgresAttachmentArgs, PostgresAttachmentState] = (*PostgresAttachment)(nil)
	_ infer.CustomDelete[PostgresAttachmentState]                           = (*PostgresAttachment)(nil)
)

type PostgresAttachmentArgs struct {
	App          string  `pulumi:"app"`
	Postgres     string  `pulumi:"postgres"`
	DatabaseName *string `pulumi:"databaseName,optional"`
	DatabaseUser *string `pulumi:"databaseUser,optional"`
	VariableName *string `pulumi:"variableName,optional"`
	ManualEntry  bool    `pulumi:"manualEntry,optional"`
}

var _ infer.Annotated = (*PostgresAttachmentArgs)(nil)

func (a *PostgresAttachmentArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.App, "The application to attach the Postgres database to.")
	anno.Describe(&a.Postgres, "The Postgres cluster to attach.")
	anno.Describe(&a.DatabaseName, "The name of the database to use.")
	anno.Describe(&a.DatabaseUser, "The database user to connect as.")
	anno.Describe(&a.VariableName, "The environment variable name to store the connection string.")
	anno.Describe(&a.ManualEntry, "Whether to manually enter the connection details.")
}

type PostgresAttachmentState struct {
	PostgresAttachmentArgs
	ConnectionString string `pulumi:"connectionString" provider:"secret"`
	VariableName     string `pulumi:"variableName"`
	AttachmentID     string `pulumi:"attachmentId"`
}

var _ infer.Annotated = (*PostgresAttachmentState)(nil)

func (s *PostgresAttachmentState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.PostgresAttachmentArgs, "The input arguments used to create the Postgres attachment.")
	anno.Describe(&s.ConnectionString, "The PostgreSQL connection string.")
	anno.Describe(&s.VariableName, "The environment variable name that contains the connection string.")
	anno.Describe(&s.AttachmentID, "The unique ID of the Postgres attachment.")
}

func (PostgresAttachment) Create(ctx context.Context, name string, input PostgresAttachmentArgs, preview bool) (string, PostgresAttachmentState, error) {
	state := PostgresAttachmentState{
		PostgresAttachmentArgs: input,
	}

	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	attachment, err := cfg.flyClient.AttachPostgresCluster(ctx, fly.AttachPostgresClusterInput{
		AppID:                input.App,
		PostgresClusterAppID: input.Postgres,
		DatabaseName:         input.DatabaseName,
		DatabaseUser:         input.DatabaseUser,
		VariableName:         input.VariableName,
		ManualEntry:          input.ManualEntry,
	})
	if err != nil {
		return name, state, err
	}

	attachments, err := cfg.flyClient.ListPostgresClusterAttachments(ctx, input.App, input.Postgres)
	if err != nil {
		return name, state, err
	}

	state.ConnectionString = attachment.ConnectionString
	state.VariableName = attachment.EnvironmentVariableName

	var attachmentID string
	for _, a := range attachments {
		if a.DatabaseName == *state.DatabaseName && a.DatabaseUser == *state.DatabaseUser && a.EnvironmentVariableName == state.VariableName {
			attachmentID = a.ID
		}
	}

	state.AttachmentID = attachmentID

	return name, state, nil
}

func (PostgresAttachment) Delete(ctx context.Context, reqID string, state PostgresAttachmentState) error {
	return infer.GetConfig[Config](ctx).flyClient.DetachPostgresCluster(ctx, fly.DetachPostgresClusterInput{
		AppID:                       state.App,
		PostgresClusterId:           state.Postgres,
		PostgresClusterAttachmentId: state.AttachmentID,
	})
}
