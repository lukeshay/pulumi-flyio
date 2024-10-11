package provider

import (
	"fmt"
	"math/rand"
	"reflect"
	"slices"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/r3labs/diff/v3"
)

type generateDiffResponseOpts struct {
	ReplaceProps             []string
	DeleteBeforeReplaceProps []string
}

func generateDiffResponse(a, b interface{}, opts generateDiffResponseOpts) (p.DiffResponse, error) {
	changelog, err := diff.Diff(a, b)
	if err != nil {
		return p.DiffResponse{}, err
	}

	diffRes := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          false,
	}

	detailedDiff := map[string]p.PropertyDiff{}

	for _, change := range changelog {
		if len(change.Path) == 0 {
			continue
		}

		path := strings.Join(change.Path, ".")
		if slices.Contains(opts.DeleteBeforeReplaceProps, path) {
			diffRes.DeleteBeforeReplace = true
		}

		propertyDiff := p.PropertyDiff{
			InputDiff: true,
		}

		replace := slices.Contains(opts.ReplaceProps, path) || slices.Contains(opts.DeleteBeforeReplaceProps, path)

		if replace {
			if change.Type == diff.CREATE {
				propertyDiff.Kind = p.AddReplace
			} else if change.Type == diff.DELETE {
				propertyDiff.Kind = p.DeleteReplace
			} else {
				propertyDiff.Kind = p.UpdateReplace
			}
		} else {
			if change.Type == diff.CREATE {
				propertyDiff.Kind = p.Add
			} else if change.Type == diff.DELETE {
				propertyDiff.Kind = p.Delete
			} else {
				propertyDiff.Kind = p.Update
			}
		}

		detailedDiff[path] = propertyDiff
	}

	if len(detailedDiff) > 0 {
		diffRes.HasChanges = true
		diffRes.DetailedDiff = detailedDiff
	}

	return diffRes, err
}

func isFirstNilAndSecondNotNil(first, second interface{}) bool {
	return first == nil && second != nil
}

func areBothNotNilAndNotEqual(first, second interface{}) bool {
	return first != nil && second != nil && !reflect.DeepEqual(reflect.ValueOf(first).Elem(), reflect.ValueOf(second).Elem())
}

type FlyRes interface {
	Status() string
}

func resErr(msg string, res FlyRes, body []byte) error {
	return fmt.Errorf("[%s] %s: %s", res.Status(), msg, body)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
