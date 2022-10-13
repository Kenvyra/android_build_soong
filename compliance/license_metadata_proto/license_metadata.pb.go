// Copyright 2021 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: license_metadata.proto

package license_metadata_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LicenseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// package_name identifies the source package. License texts are named relative to the package name.
	PackageName *string `protobuf:"bytes,1,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
	// module_name identifies the target
	ModuleName    *string  `protobuf:"bytes,14,opt,name=module_name,json=moduleName" json:"module_name,omitempty"`
	ModuleTypes   []string `protobuf:"bytes,2,rep,name=module_types,json=moduleTypes" json:"module_types,omitempty"`
	ModuleClasses []string `protobuf:"bytes,3,rep,name=module_classes,json=moduleClasses" json:"module_classes,omitempty"`
	// projects identifies the git project(s) containing the associated source code.
	Projects []string `protobuf:"bytes,4,rep,name=projects" json:"projects,omitempty"`
	// license_kinds lists the kinds of licenses. e.g. SPDX-license-identifier-Apache-2.0 or legacy_notice.
	LicenseKinds []string `protobuf:"bytes,5,rep,name=license_kinds,json=licenseKinds" json:"license_kinds,omitempty"`
	// license_conditions lists the conditions that apply to the license kinds. e.g. notice or restricted.
	LicenseConditions []string `protobuf:"bytes,6,rep,name=license_conditions,json=licenseConditions" json:"license_conditions,omitempty"`
	// license_texts lists the filenames of the associated license text(s).
	LicenseTexts []string `protobuf:"bytes,7,rep,name=license_texts,json=licenseTexts" json:"license_texts,omitempty"`
	// is_container is true for target types that merely aggregate. e.g. .img or .zip files.
	IsContainer *bool `protobuf:"varint,8,opt,name=is_container,json=isContainer" json:"is_container,omitempty"`
	// built lists the built targets.
	Built []string `protobuf:"bytes,9,rep,name=built" json:"built,omitempty"`
	// installed lists the installed targets.
	Installed []string `protobuf:"bytes,10,rep,name=installed" json:"installed,omitempty"`
	// installMap identifies the substitutions to make to path names when moving into installed location.
	InstallMap []*InstallMap `protobuf:"bytes,11,rep,name=install_map,json=installMap" json:"install_map,omitempty"`
	// sources lists the targets depended on.
	Sources []string `protobuf:"bytes,12,rep,name=sources" json:"sources,omitempty"`
	// deps lists the license metadata files depended on.
	Deps []*AnnotatedDependency `protobuf:"bytes,13,rep,name=deps" json:"deps,omitempty"`
}

func (x *LicenseMetadata) Reset() {
	*x = LicenseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_license_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseMetadata) ProtoMessage() {}

func (x *LicenseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_license_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseMetadata.ProtoReflect.Descriptor instead.
func (*LicenseMetadata) Descriptor() ([]byte, []int) {
	return file_license_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *LicenseMetadata) GetPackageName() string {
	if x != nil && x.PackageName != nil {
		return *x.PackageName
	}
	return ""
}

func (x *LicenseMetadata) GetModuleName() string {
	if x != nil && x.ModuleName != nil {
		return *x.ModuleName
	}
	return ""
}

func (x *LicenseMetadata) GetModuleTypes() []string {
	if x != nil {
		return x.ModuleTypes
	}
	return nil
}

func (x *LicenseMetadata) GetModuleClasses() []string {
	if x != nil {
		return x.ModuleClasses
	}
	return nil
}

func (x *LicenseMetadata) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *LicenseMetadata) GetLicenseKinds() []string {
	if x != nil {
		return x.LicenseKinds
	}
	return nil
}

func (x *LicenseMetadata) GetLicenseConditions() []string {
	if x != nil {
		return x.LicenseConditions
	}
	return nil
}

func (x *LicenseMetadata) GetLicenseTexts() []string {
	if x != nil {
		return x.LicenseTexts
	}
	return nil
}

func (x *LicenseMetadata) GetIsContainer() bool {
	if x != nil && x.IsContainer != nil {
		return *x.IsContainer
	}
	return false
}

func (x *LicenseMetadata) GetBuilt() []string {
	if x != nil {
		return x.Built
	}
	return nil
}

func (x *LicenseMetadata) GetInstalled() []string {
	if x != nil {
		return x.Installed
	}
	return nil
}

func (x *LicenseMetadata) GetInstallMap() []*InstallMap {
	if x != nil {
		return x.InstallMap
	}
	return nil
}

func (x *LicenseMetadata) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *LicenseMetadata) GetDeps() []*AnnotatedDependency {
	if x != nil {
		return x.Deps
	}
	return nil
}

// InstallMap messages describe the mapping from an input filesystem file to the path to the file
// in a container.
type InstallMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The input path on the filesystem.
	FromPath *string `protobuf:"bytes,1,opt,name=from_path,json=fromPath" json:"from_path,omitempty"`
	// The path to the file inside the container or installed location.
	ContainerPath *string `protobuf:"bytes,2,opt,name=container_path,json=containerPath" json:"container_path,omitempty"`
}

func (x *InstallMap) Reset() {
	*x = InstallMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_license_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallMap) ProtoMessage() {}

func (x *InstallMap) ProtoReflect() protoreflect.Message {
	mi := &file_license_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallMap.ProtoReflect.Descriptor instead.
func (*InstallMap) Descriptor() ([]byte, []int) {
	return file_license_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *InstallMap) GetFromPath() string {
	if x != nil && x.FromPath != nil {
		return *x.FromPath
	}
	return ""
}

func (x *InstallMap) GetContainerPath() string {
	if x != nil && x.ContainerPath != nil {
		return *x.ContainerPath
	}
	return ""
}

// AnnotateDepencency messages describe edges in the build graph.
type AnnotatedDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to the dependency license metadata file.
	File *string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	// The annotations attached to the dependency.
	Annotations []string `protobuf:"bytes,2,rep,name=annotations" json:"annotations,omitempty"`
}

func (x *AnnotatedDependency) Reset() {
	*x = AnnotatedDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_license_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotatedDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotatedDependency) ProtoMessage() {}

func (x *AnnotatedDependency) ProtoReflect() protoreflect.Message {
	mi := &file_license_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotatedDependency.ProtoReflect.Descriptor instead.
func (*AnnotatedDependency) Descriptor() ([]byte, []int) {
	return file_license_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *AnnotatedDependency) GetFile() string {
	if x != nil && x.File != nil {
		return *x.File
	}
	return ""
}

func (x *AnnotatedDependency) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_license_metadata_proto protoreflect.FileDescriptor

var file_license_metadata_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xab, 0x04, 0x0a, 0x0f, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x69,
	0x6e, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75,
	0x69, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x43,
	0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x4d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a,
	0x04, 0x64, 0x65, 0x70, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x22, 0x50,
	0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x4b, 0x0a, 0x13, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x31, 0x5a,
	0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2f, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_license_metadata_proto_rawDescOnce sync.Once
	file_license_metadata_proto_rawDescData = file_license_metadata_proto_rawDesc
)

func file_license_metadata_proto_rawDescGZIP() []byte {
	file_license_metadata_proto_rawDescOnce.Do(func() {
		file_license_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_license_metadata_proto_rawDescData)
	})
	return file_license_metadata_proto_rawDescData
}

var file_license_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_license_metadata_proto_goTypes = []interface{}{
	(*LicenseMetadata)(nil),     // 0: build_license_metadata.LicenseMetadata
	(*InstallMap)(nil),          // 1: build_license_metadata.InstallMap
	(*AnnotatedDependency)(nil), // 2: build_license_metadata.AnnotatedDependency
}
var file_license_metadata_proto_depIdxs = []int32{
	1, // 0: build_license_metadata.LicenseMetadata.install_map:type_name -> build_license_metadata.InstallMap
	2, // 1: build_license_metadata.LicenseMetadata.deps:type_name -> build_license_metadata.AnnotatedDependency
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_license_metadata_proto_init() }
func file_license_metadata_proto_init() {
	if File_license_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_license_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_license_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_license_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotatedDependency); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_license_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_license_metadata_proto_goTypes,
		DependencyIndexes: file_license_metadata_proto_depIdxs,
		MessageInfos:      file_license_metadata_proto_msgTypes,
	}.Build()
	File_license_metadata_proto = out.File
	file_license_metadata_proto_rawDesc = nil
	file_license_metadata_proto_goTypes = nil
	file_license_metadata_proto_depIdxs = nil
}
