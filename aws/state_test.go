package aws

import (
	"reflect"
	"testing"
)

func TestStateNotinCloudformation(t *testing.T) {
	state := state{
		cfn: map[resourceType][]string{
			s3Bucket:  []string{"arn1"},
			ec2Volume: []string{},
		},
		real: map[resourceType][]string{
			s3Bucket:  []string{"arn1", "exists"},
			ec2Volume: []string{"existstoo"},
		},
	}
	expected := resourceMap{
		s3Bucket:  []string{"exists"},
		ec2Volume: []string{"existstoo"},
	}
	result := state.filter(real, cfn)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateCloudformationNotinreal(t *testing.T) {
	state := state{
		real: map[resourceType][]string{
			s3Bucket:  []string{"arn1"},
			ec2Volume: []string{},
		},
		cfn: map[resourceType][]string{
			s3Bucket:  []string{"arn1", "exists"},
			ec2Volume: []string{"existstoo"},
		},
	}
	expected := resourceMap{
		s3Bucket:  []string{"exists"},
		ec2Volume: []string{"existstoo"},
	}
	result := state.filter(cfn, real)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudformationNull(t *testing.T) {
	state := state{
		cfn: map[resourceType][]string{
			s3Bucket:  []string{},
			ec2Volume: []string{},
		},
		real: map[resourceType][]string{
			s3Bucket:  []string{},
			ec2Volume: []string{},
		},
	}
	expected := resourceMap{}
	result := state.filter(real, cfn)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudformationEqual(t *testing.T) {
	state := state{
		cfn: map[resourceType][]string{
			s3Bucket:  []string{"arn1", "exists"},
			ec2Volume: []string{"existstoo"},
		},
		real: map[resourceType][]string{
			s3Bucket:  []string{"arn1", "exists"},
			ec2Volume: []string{"existstoo"},
		},
	}
	expected := resourceMap{}
	result := state.filter(real, cfn)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudformationMissingEntryInCloudformation(t *testing.T) {
	state := state{
		cfn: map[resourceType][]string{
			s3Bucket: []string{"arn1", "exists"},
		},
		real: map[resourceType][]string{
			s3Bucket:  []string{"arn1", "exists"},
			ec2Volume: []string{"existstoo"},
		},
	}
	expected := resourceMap{
		ec2Volume: []string{"existstoo"},
	}
	result := state.filter(real, cfn)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result  : %v", result)
		t.Errorf("expected: %v", expected)
	}
}

func TestStateNotinCloudformationMissingEntryInreal(t *testing.T) {
	state := state{
		cfn: map[resourceType][]string{
			s3Bucket:  []string{"arn1", "exists"},
			ec2Volume: []string{"existstoo"},
		},
		real: map[resourceType][]string{
			s3Bucket: []string{"arn1", "exists"},
		},
	}
	expected := resourceMap{}
	result := state.filter(real, cfn)
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
