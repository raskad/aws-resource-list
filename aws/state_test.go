package aws

import (
	"reflect"
	"testing"
)

func TestStateNotinCloudFormation(t *testing.T) {
	state := state{
		Cfn: map[string][]string{
			"AWS::S3::Bucket":  {"arn1"},
			"AWS::EC2::Volume": {},
		},
		Real: map[resourceType][]string{
			s3Bucket:  {"arn1", "exists"},
			ec2Volume: {"existstoo"},
		},
	}
	expected := awsResourceMap{
		s3Bucket:  []string{"exists"},
		ec2Volume: []string{"existstoo"},
	}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudFormationNull(t *testing.T) {
	state := state{
		Cfn: map[string][]string{
			"AWS::S3::Bucket":  {},
			"AWS::EC2::Volume": {},
		},
		Real: map[resourceType][]string{
			s3Bucket:  {},
			ec2Volume: {},
		},
	}
	expected := awsResourceMap{}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudFormationEqual(t *testing.T) {
	state := state{
		Cfn: map[string][]string{
			"AWS::S3::Bucket":  {"arn1", "exists"},
			"AWS::EC2::Volume": {"existstoo"},
		},
		Real: map[resourceType][]string{
			s3Bucket:  {"arn1", "exists"},
			ec2Volume: {"existstoo"},
		},
	}
	expected := awsResourceMap{}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudFormationMissingEntryInCloudFormation(t *testing.T) {
	state := state{
		Cfn: map[string][]string{
			"AWS::S3::Bucket": {"arn1", "exists"},
		},
		Real: map[resourceType][]string{
			s3Bucket:  {"arn1", "exists"},
			ec2Volume: {"existstoo"},
		},
	}
	expected := awsResourceMap{
		ec2Volume: []string{"existstoo"},
	}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudFormationMissingEntryInreal(t *testing.T) {
	state := state{
		Cfn: map[string][]string{
			"AWS::S3::Bucket":  {"arn1", "exists"},
			"AWS::EC2::Volume": {"existstoo"},
		},
		Real: map[resourceType][]string{
			s3Bucket: {"arn1", "exists"},
		},
	}
	expected := awsResourceMap{}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinTerraform(t *testing.T) {
	state := state{
		Tf: map[string][]string{
			"aws_s3_bucket": {"arn1"},
			"aws_volume":    {},
		},
		Real: map[resourceType][]string{
			s3Bucket:  {"arn1", "exists"},
			ec2Volume: {"existstoo"},
		},
	}
	expected := awsResourceMap{
		s3Bucket:  []string{"exists"},
		ec2Volume: []string{"existstoo"},
	}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinTerraformMultiple(t *testing.T) {
	state := state{
		Tf: map[string][]string{
			"aws_fsx_lustre_file_system":  {"lustre1"},
			"aws_fsx_windows_file_system": {"windows1"},
		},
		Real: map[resourceType][]string{
			fsxFileSystem: {"lustre1", "windows1", "windows2"},
		},
	}
	expected := awsResourceMap{
		fsxFileSystem: []string{"windows2"},
	}
	result := state.filter()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestFromCloudFormationTypeOk(t *testing.T) {
	cloudFormationType := "AWS::EC2::Instance"
	expected := ec2Instance
	result, ok := fromCloudFormationType(cloudFormationType)
	if !ok {
		t.Errorf("result  : %v", ok)
		t.Errorf("expected: %v", true)
	}
	if result != expected {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestFromCloudFormationTypeUnknownType(t *testing.T) {
	cloudFormationType := "AWS::NewService::NewType"
	_, ok := fromCloudFormationType(cloudFormationType)
	if ok {
		t.Errorf("result  : %v", ok)
		t.Errorf("expected: %v", false)
	}
}
