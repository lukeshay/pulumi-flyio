package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Add annotations
type Volume struct{}

var (
	_ infer.CustomResource[VolumeArgs, VolumeState] = (*Volume)(nil)
	_ infer.CustomDelete[VolumeState]               = (*Volume)(nil)
	_ infer.CustomRead[VolumeArgs, VolumeState]     = (*Volume)(nil)
	_ infer.CustomUpdate[VolumeArgs, VolumeState]   = (*Volume)(nil)
)

type VolumeArgs struct {
	flyio.CreateVolumeRequest
	AutoBackupEnabled *bool  `pulumi:"autoBackupEnabled,optional"`
	App               string `pulumi:"app"`
}

type VolumeState struct {
	flyio.Volume
	Input VolumeArgs `pulumi:"input"`
	App   string     `pulumi:"app"`
}

func (v Volume) Create(ctx context.Context, name string, input VolumeArgs, preview bool) (string, VolumeState, error) {
	state := VolumeState{Input: input, App: input.App}
	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.VolumesCreate(ctx, input.App, input.CreateVolumeRequest)
	if err != nil {
		return "", VolumeState{}, err
	}

	result, err := v.parseVolumesCreateResponse(res)
	if err != nil {
		return "", VolumeState{}, err
	}

	if result.JSON200 == nil {
		return "", VolumeState{}, resErr("error creating volume", result, result.Body)
	}
	state.Volume = *result.JSON200

	if input.AutoBackupEnabled != nil && !*input.AutoBackupEnabled {
		res, err = cfg.flyioClient.VolumesUpdate(ctx, input.App, *result.JSON200.Id, flyio.UpdateVolumeRequest{
			AutoBackupEnabled: input.AutoBackupEnabled,
		})
		if err != nil {
			cfg.flyioClient.VolumeDelete(ctx, input.App, *state.Id)
			return "", VolumeState{}, err
		}

		result2, err := flyio.ParseVolumesUpdateResponse(res)
		if err != nil {
			cfg.flyioClient.VolumeDelete(ctx, input.App, *state.Id)
			return "", VolumeState{}, err
		}

		if result2.JSON200 == nil {
			cfg.flyioClient.VolumeDelete(ctx, input.App, *state.Id)
			return "", VolumeState{}, resErr("error updating volume", result, result.Body)
		}

		state.Volume = *result2.JSON200
	}

	return *result.JSON200.Id, state, nil
}

func (Volume) Delete(ctx context.Context, reqID string, state VolumeState) error {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.VolumeDelete(ctx, state.App, *state.Id)
	if err != nil {
		return err
	}

	result, err := flyio.ParseVolumeDeleteResponse(res)
	if err != nil {
		return err
	}

	if result.JSON200 == nil {
		return resErr("error deleting volume", result, result.Body)
	}

	return nil
}

func (Volume) Read(ctx context.Context, id string, inputs VolumeArgs, state VolumeState) (
	canonicalID string, normalizedInputs VolumeArgs, normalizedState VolumeState, err error,
) {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.VolumesGetById(ctx, *state.Name, *state.Id)
	if err != nil {
		return id, inputs, state, err
	}

	result, err := flyio.ParseVolumesGetByIdResponse(res)
	if err != nil {
		return id, inputs, state, err
	}

	if result.JSON200 == nil {
		return id, inputs, state, fmt.Errorf("error getting volume: %s", result.Body)
	}

	state.Volume = *result.JSON200

	return id, inputs, state, nil
}

var volumeDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{"Compute", "ComputeImage", "Encrypted", "Fstype", "MachinesOnly", "Name", "Region", "RequireUniqueZone", "SnapshotId", "SourceVolumeId", "AppName"},
	DeleteBeforeReplaceProps: []string{},
}

func (Volume) Diff(ctx context.Context, id string, state VolumeState, input VolumeArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, volumeDiffOpts)
}

func (m Volume) Update(ctx context.Context, id string, state VolumeState, input VolumeArgs, preview bool) (VolumeState, error) {
	diff, _ := m.Diff(ctx, id, state, input)

	if diff.DeleteBeforeReplace {
		if input.SourceVolumeId == nil {
			input.SourceVolumeId = state.Id
		}

		_, newState, err := m.Create(ctx, id, input, false)
		if err != nil {
			return state, err
		}

		m.Delete(ctx, id, state)

		return newState, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	if isFirstNilAndSecondNotNil(state.AutoBackupEnabled, input.AutoBackupEnabled) ||
		areBothNotNilAndNotEqual(state.AutoBackupEnabled, input.AutoBackupEnabled) ||
		isFirstNilAndSecondNotNil(state.SnapshotRetention, input.SnapshotRetention) ||
		areBothNotNilAndNotEqual(state.SnapshotRetention, input.SnapshotRetention) {
		res, err := cfg.flyioClient.VolumesUpdate(ctx, state.App, *state.Id, flyio.UpdateVolumeRequest{
			AutoBackupEnabled: input.AutoBackupEnabled,
			SnapshotRetention: input.SnapshotRetention,
		})
		if err != nil {
			return state, err
		}

		result, err := flyio.ParseVolumesUpdateResponse(res)
		if err != nil {
			return state, err
		}

		if result.JSON200 == nil {
			return state, fmt.Errorf("error updating volume: %s", result.Body)
		}

		state.Volume = *result.JSON200
	}

	if isFirstNilAndSecondNotNil(state.SizeGb, input.SizeGb) || areBothNotNilAndNotEqual(state.SizeGb, input.SizeGb) {
		res, err := cfg.flyioClient.VolumesExtend(ctx, state.App, *state.Id, flyio.ExtendVolumeRequest{
			SizeGb: input.SizeGb,
		})
		if err != nil {
			return state, err
		}

		result, err := flyio.ParseVolumesExtendResponse(res)
		if err != nil {
			return state, err
		}

		state.Volume = *result.JSON200.Volume
	}

	state.App = input.App

	return state, nil
}

func (Volume) parseVolumesCreateResponse(rsp *http.Response) (*flyio.VolumesCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &flyio.VolumesCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode < 300:
		var dest flyio.Volume
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
