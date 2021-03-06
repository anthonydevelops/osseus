// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rest_template_structure.proto

package restmodel

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//TemplateStructure holds the directory and folder structure of the project
type TemplateStructure struct {
	// List of file objects describing directory structure of generated files
	File                 []*File  `protobuf:"bytes,1,rep,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateStructure) Reset()         { *m = TemplateStructure{} }
func (m *TemplateStructure) String() string { return proto.CompactTextString(m) }
func (*TemplateStructure) ProtoMessage()    {}
func (*TemplateStructure) Descriptor() ([]byte, []int) {
	return fileDescriptor_e578eed11d4f8452, []int{0}
}
func (m *TemplateStructure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateStructure.Unmarshal(m, b)
}
func (m *TemplateStructure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateStructure.Marshal(b, m, deterministic)
}
func (m *TemplateStructure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateStructure.Merge(m, src)
}
func (m *TemplateStructure) XXX_Size() int {
	return xxx_messageInfo_TemplateStructure.Size(m)
}
func (m *TemplateStructure) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateStructure.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateStructure proto.InternalMessageInfo

func (m *TemplateStructure) GetFile() []*File {
	if m != nil {
		return m.File
	}
	return nil
}

func (*TemplateStructure) XXX_MessageName() string {
	return "restmodel.TemplateStructure"
}

// File holds the folder path, children, type, and etcdkey of the given file
type File struct {
	// Name of generated code file or folder
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Path of generated code file
	AbsolutePath string `protobuf:"bytes,2,opt,name=absolute_path,json=absolutePath,proto3" json:"absolute_path,omitempty"`
	// "Folder" if a folder in the directory or "File" if code file
	FileType string `protobuf:"bytes,3,opt,name=fileType,proto3" json:"fileType,omitempty"`
	// Key the file is stored under in etcd
	EtcdKey string `protobuf:"bytes,4,opt,name=etcdKey,proto3" json:"etcdKey,omitempty"`
	// Absolute path(s) of direct children folders of the file, empty list if no children
	Children             []string `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_e578eed11d4f8452, []int{1}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetAbsolutePath() string {
	if m != nil {
		return m.AbsolutePath
	}
	return ""
}

func (m *File) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

func (m *File) GetEtcdKey() string {
	if m != nil {
		return m.EtcdKey
	}
	return ""
}

func (m *File) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

func (*File) XXX_MessageName() string {
	return "restmodel.File"
}

// FileContent holds the content of the go file
type FileContent struct {
	// Content of the generated code file
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileContent) Reset()         { *m = FileContent{} }
func (m *FileContent) String() string { return proto.CompactTextString(m) }
func (*FileContent) ProtoMessage()    {}
func (*FileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e578eed11d4f8452, []int{2}
}
func (m *FileContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileContent.Unmarshal(m, b)
}
func (m *FileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileContent.Marshal(b, m, deterministic)
}
func (m *FileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileContent.Merge(m, src)
}
func (m *FileContent) XXX_Size() int {
	return xxx_messageInfo_FileContent.Size(m)
}
func (m *FileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_FileContent.DiscardUnknown(m)
}

var xxx_messageInfo_FileContent proto.InternalMessageInfo

func (m *FileContent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (*FileContent) XXX_MessageName() string {
	return "restmodel.FileContent"
}
func init() {
	proto.RegisterType((*TemplateStructure)(nil), "restmodel.TemplateStructure")
	proto.RegisterType((*File)(nil), "restmodel.File")
	proto.RegisterType((*FileContent)(nil), "restmodel.FileContent")
}

func init() { proto.RegisterFile("rest_template_structure.proto", fileDescriptor_e578eed11d4f8452) }

var fileDescriptor_e578eed11d4f8452 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x62, 0xfe, 0xd4, 0x05, 0x21, 0x3c, 0x59, 0x95, 0x40, 0x55, 0x3a, 0xd0, 0x85,
	0x54, 0x82, 0x85, 0x19, 0x24, 0x16, 0x16, 0x54, 0xba, 0x47, 0x8e, 0x73, 0x4d, 0x22, 0x39, 0x71,
	0xe4, 0x9c, 0x87, 0x3e, 0x06, 0x6f, 0xc5, 0x7b, 0xf0, 0x22, 0xc8, 0xd7, 0x3a, 0xdb, 0xfd, 0xfc,
	0xdd, 0xf7, 0xd3, 0xc9, 0xfc, 0xde, 0xc1, 0x88, 0x05, 0x42, 0x37, 0x18, 0x85, 0x50, 0x8c, 0xe8,
	0xbc, 0x46, 0xef, 0x20, 0x1f, 0x9c, 0x45, 0x2b, 0x66, 0x21, 0xee, 0x6c, 0x05, 0x66, 0xf1, 0x54,
	0xb7, 0xd8, 0xf8, 0x32, 0xd7, 0xb6, 0xdb, 0xd4, 0xb6, 0xb6, 0x1b, 0xda, 0x28, 0xfd, 0x9e, 0x88,
	0x80, 0xa6, 0x63, 0x33, 0x7b, 0xe5, 0x77, 0xbb, 0x93, 0xf5, 0x3b, 0x4a, 0xc5, 0x8a, 0xb3, 0x7d,
	0x6b, 0x40, 0x26, 0xcb, 0x74, 0x3d, 0x7f, 0xbe, 0xcd, 0x27, 0x7b, 0xfe, 0xd1, 0x1a, 0xd8, 0x52,
	0x98, 0xfd, 0x24, 0x9c, 0x05, 0x14, 0x82, 0xb3, 0x5e, 0x75, 0x61, 0x3b, 0x59, 0xcf, 0xb6, 0x34,
	0x8b, 0x15, 0xbf, 0x51, 0xe5, 0x68, 0x8d, 0x47, 0x28, 0x06, 0x85, 0x8d, 0x3c, 0xa3, 0xf0, 0x3a,
	0x3e, 0x7e, 0x29, 0x6c, 0xc4, 0x82, 0x5f, 0x05, 0xd3, 0xee, 0x30, 0x80, 0x4c, 0x29, 0x9f, 0x58,
	0x48, 0x7e, 0x09, 0xa8, 0xab, 0x4f, 0x38, 0x48, 0x46, 0x51, 0xc4, 0xd0, 0xd2, 0x4d, 0x6b, 0x2a,
	0x07, 0xbd, 0x3c, 0x5f, 0xa6, 0xa1, 0x15, 0x39, 0x7b, 0xe4, 0xf3, 0x70, 0xd2, 0xbb, 0xed, 0x11,
	0x7a, 0x0c, 0x12, 0x7d, 0x1c, 0x4f, 0xc7, 0x45, 0x7c, 0x63, 0xbf, 0x7f, 0x0f, 0x49, 0x79, 0x41,
	0x7f, 0xf0, 0xf2, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x10, 0xd1, 0x81, 0xaf, 0x5e, 0x01, 0x00, 0x00,
}
