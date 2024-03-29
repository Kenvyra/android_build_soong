package android

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"android/soong/bazel/cquery"
)

func TestRequestResultsAfterInvokeBazel(t *testing.T) {
	label := "//foo:bar"
	cfg := configKey{"arm64_armv8-a", Android}
	bazelContext, _ := testBazelContext(t, map[bazelCommand]string{
		bazelCommand{command: "cquery", expression: "deps(@soong_injection//mixed_builds:buildroot, 2)"}: `//foo:bar|arm64_armv8-a|android>>out/foo/bar.txt`,
	})
	bazelContext.QueueBazelRequest(label, cquery.GetOutputFiles, cfg)
	err := bazelContext.InvokeBazel()
	if err != nil {
		t.Fatalf("Did not expect error invoking Bazel, but got %s", err)
	}
	g, err := bazelContext.GetOutputFiles(label, cfg)
	if err != nil {
		t.Errorf("Expected cquery results after running InvokeBazel(), but got err %v", err)
	} else if w := []string{"out/foo/bar.txt"}; !reflect.DeepEqual(w, g) {
		t.Errorf("Expected output %s, got %s", w, g)
	}
}

func TestInvokeBazelWritesBazelFiles(t *testing.T) {
	bazelContext, baseDir := testBazelContext(t, map[bazelCommand]string{})
	err := bazelContext.InvokeBazel()
	if err != nil {
		t.Fatalf("Did not expect error invoking Bazel, but got %s", err)
	}
	if _, err := os.Stat(filepath.Join(baseDir, "soong_injection", "mixed_builds", "main.bzl")); os.IsNotExist(err) {
		t.Errorf("Expected main.bzl to exist, but it does not")
	} else if err != nil {
		t.Errorf("Unexpected error stating main.bzl %s", err)
	}

	if _, err := os.Stat(filepath.Join(baseDir, "soong_injection", "mixed_builds", "BUILD.bazel")); os.IsNotExist(err) {
		t.Errorf("Expected BUILD.bazel to exist, but it does not")
	} else if err != nil {
		t.Errorf("Unexpected error stating BUILD.bazel %s", err)
	}

	if _, err := os.Stat(filepath.Join(baseDir, "soong_injection", "WORKSPACE.bazel")); os.IsNotExist(err) {
		t.Errorf("Expected WORKSPACE.bazel to exist, but it does not")
	} else if err != nil {
		t.Errorf("Unexpected error stating WORKSPACE.bazel %s", err)
	}
}

func TestInvokeBazelPopulatesBuildStatements(t *testing.T) {
	bazelContext, _ := testBazelContext(t, map[bazelCommand]string{
		bazelCommand{command: "aquery", expression: "deps(@soong_injection//mixed_builds:buildroot)"}: `
{
  "artifacts": [{
    "id": 1,
    "pathFragmentId": 1
  }, {
    "id": 2,
    "pathFragmentId": 2
  }],
  "actions": [{
    "targetId": 1,
    "actionKey": "x",
    "mnemonic": "x",
    "arguments": ["touch", "foo"],
    "inputDepSetIds": [1],
    "outputIds": [1],
    "primaryOutputId": 1
  }],
  "depSetOfFiles": [{
    "id": 1,
    "directArtifactIds": [1, 2]
  }],
  "pathFragments": [{
    "id": 1,
    "label": "one"
  }, {
    "id": 2,
    "label": "two"
  }]
}`,
	})
	err := bazelContext.InvokeBazel()
	if err != nil {
		t.Fatalf("Did not expect error invoking Bazel, but got %s", err)
	}

	got := bazelContext.BuildStatementsToRegister()
	if want := 1; len(got) != want {
		t.Errorf("Expected %d registered build statements, got %#v", want, got)
	}
}

func testBazelContext(t *testing.T, bazelCommandResults map[bazelCommand]string) (*bazelContext, string) {
	t.Helper()
	p := bazelPaths{
		soongOutDir:  t.TempDir(),
		outputBase:   "outputbase",
		workspaceDir: "workspace_dir",
	}
	aqueryCommand := bazelCommand{command: "aquery", expression: "deps(@soong_injection//mixed_builds:buildroot)"}
	if _, exists := bazelCommandResults[aqueryCommand]; !exists {
		bazelCommandResults[aqueryCommand] = "{}\n"
	}
	runner := &mockBazelRunner{bazelCommandResults: bazelCommandResults}
	return &bazelContext{
		bazelRunner: runner,
		paths:       &p,
		requests:    map[cqueryKey]bool{},
	}, p.soongOutDir
}
