package provider

import (
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
	AppName           string `pulumi:"appName"`
}

type VolumeState struct {
	flyio.Volume
	Input   VolumeArgs `pulumi:"input"`
	AppName string     `pulumi:"appName"`
}

func (v Volume) Create(ctx p.Context, name string, input VolumeArgs, preview bool) (string, VolumeState, error) {
	state := VolumeState{Input: input, AppName: input.AppName}
	if preview {
		return name, state, nil
	}

	client, err := getFlyClient()
	if err != nil {
		return "", VolumeState{}, err
	}

	res, err := client.VolumesCreate(ctx, input.AppName, input.CreateVolumeRequest)
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
		res, err = client.VolumesUpdate(ctx, input.AppName, *result.JSON200.Id, flyio.UpdateVolumeRequest{
			AutoBackupEnabled: input.AutoBackupEnabled,
		})
		if err != nil {
			client.VolumeDelete(ctx, input.AppName, *state.Id)
			return "", VolumeState{}, err
		}

		result2, err := flyio.ParseVolumesUpdateResponse(res)
		if err != nil {
			client.VolumeDelete(ctx, input.AppName, *state.Id)
			return "", VolumeState{}, err
		}

		if result2.JSON200 == nil {
			client.VolumeDelete(ctx, input.AppName, *state.Id)
			return "", VolumeState{}, resErr("error updating volume", result, result.Body)
		}

		state.Volume = *result2.JSON200
	}

	return *result.JSON200.Id, state, nil
}

func (Volume) Delete(ctx p.Context, reqID string, state VolumeState) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	res, err := client.VolumeDelete(ctx, state.AppName, *state.Id)
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

func (Volume) Read(ctx p.Context, id string, inputs VolumeArgs, state VolumeState) (
	canonicalID string, normalizedInputs VolumeArgs, normalizedState VolumeState, err error,
) {
	client, err := getFlyClient()
	if err != nil {
		return id, inputs, state, err
	}

	res, err := client.VolumesGetById(ctx, *state.Name, *state.Id)
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
	ReplaceProps:             []string{},
	DeleteBeforeReplaceProps: []string{"Compute", "ComputeImage", "Encrypted", "Fstype", "MachinesOnly", "Name", "Region", "RequireUniqueZone", "SnapshotId", "SourceVolumeId", "AppName"},
}

func (Volume) Diff(ctx p.Context, id string, state VolumeState, input VolumeArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, volumeDiffOpts)
}

func (m Volume) Update(ctx p.Context, id string, state VolumeState, input VolumeArgs, preview bool) (VolumeState, error) {
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

	client, err := getFlyClient()
	if err != nil {
		return state, err
	}

	if isFirstNilAndSecondNotNil(state.AutoBackupEnabled, input.AutoBackupEnabled) ||
		areBothNotNilAndNotEqual(state.AutoBackupEnabled, input.AutoBackupEnabled) ||
		isFirstNilAndSecondNotNil(state.SnapshotRetention, input.SnapshotRetention) ||
		areBothNotNilAndNotEqual(state.SnapshotRetention, input.SnapshotRetention) {
		res, err := client.VolumesUpdate(ctx, state.AppName, *state.Id, flyio.UpdateVolumeRequest{
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
		res, err := client.VolumesExtend(ctx, state.AppName, *state.Id, flyio.ExtendVolumeRequest{
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

	state.AppName = input.AppName

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
