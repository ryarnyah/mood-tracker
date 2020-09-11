// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mood.proto

package mood

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Entry) Validate() error {
	if !(this.Record > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Record", fmt.Errorf(`value '%v' must be greater than '0'`, this.Record))
	}
	if !(this.Record < 4) {
		return github_com_mwitkow_go_proto_validators.FieldError("Record", fmt.Errorf(`value '%v' must be less than '4'`, this.Record))
	}
	if !(len(this.Comment) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Comment", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Comment))
	}
	return nil
}
func (this *EntryWithDate) Validate() error {
	if !(this.Record > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Record", fmt.Errorf(`value '%v' must be greater than '0'`, this.Record))
	}
	if !(this.Record < 4) {
		return github_com_mwitkow_go_proto_validators.FieldError("Record", fmt.Errorf(`value '%v' must be less than '4'`, this.Record))
	}
	if !(len(this.Comment) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Comment", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Comment))
	}
	if this.RecordEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RecordEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RecordEntry", err)
		}
	}
	return nil
}

var _regex_AddEntryRequest_EntryAccessCode = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *AddEntryRequest) Validate() error {
	if nil == this.Entry {
		return github_com_mwitkow_go_proto_validators.FieldError("Entry", fmt.Errorf("message must exist"))
	}
	if this.Entry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Entry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Entry", err)
		}
	}
	if !(this.MoodId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodId", fmt.Errorf(`value '%v' must be greater than '0'`, this.MoodId))
	}
	if !_regex_AddEntryRequest_EntryAccessCode.MatchString(this.EntryAccessCode) {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryAccessCode", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.EntryAccessCode))
	}
	if this.EntryAccessCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryAccessCode", fmt.Errorf(`value '%v' must not be an empty string`, this.EntryAccessCode))
	}
	return nil
}
func (this *AddEntryResponse) Validate() error {
	return nil
}

var _regex_GetMoodFromEntryRequest_EntryAccessCode = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetMoodFromEntryRequest) Validate() error {
	if !(this.MoodId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodId", fmt.Errorf(`value '%v' must be greater than '0'`, this.MoodId))
	}
	if !_regex_GetMoodFromEntryRequest_EntryAccessCode.MatchString(this.EntryAccessCode) {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryAccessCode", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.EntryAccessCode))
	}
	if this.EntryAccessCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryAccessCode", fmt.Errorf(`value '%v' must not be an empty string`, this.EntryAccessCode))
	}
	return nil
}
func (this *GetMoodFromEntryResponse) Validate() error {
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	if !(len(this.Title) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Title))
	}
	if !(len(this.Content) < 513) {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must have a length smaller than '513'`, this.Content))
	}
	return nil
}

var _regex_GetMoodRequest_MoodAccessCode = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetMoodRequest) Validate() error {
	if !(this.MoodId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodId", fmt.Errorf(`value '%v' must be greater than '0'`, this.MoodId))
	}
	if !_regex_GetMoodRequest_MoodAccessCode.MatchString(this.MoodAccessCode) {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodAccessCode", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.MoodAccessCode))
	}
	if this.MoodAccessCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodAccessCode", fmt.Errorf(`value '%v' must not be an empty string`, this.MoodAccessCode))
	}
	return nil
}
func (this *RecordStat) Validate() error {
	if this.RecordEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RecordEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RecordEntry", err)
		}
	}
	return nil
}
func (this *MoodStat) Validate() error {
	for _, item := range this.RecordStats {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RecordStats", err)
			}
		}
	}
	return nil
}
func (this *GetMoodResponse) Validate() error {
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	if !(len(this.Title) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Title))
	}
	if !(len(this.Content) < 513) {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must have a length smaller than '513'`, this.Content))
	}
	for _, item := range this.Entries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Entries", err)
			}
		}
	}
	for _, item := range this.Stats {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Stats", err)
			}
		}
	}
	return nil
}

var _regex_CreateMoodRequest_Emails = regexp.MustCompile(`^[a-zA-Z0-9.!#$%%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

func (this *CreateMoodRequest) Validate() error {
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	if !(len(this.Title) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Title))
	}
	if !(len(this.Content) < 513) {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must have a length smaller than '513'`, this.Content))
	}
	if !(this.NumberOfRecordsNeeded < 21) {
		return github_com_mwitkow_go_proto_validators.FieldError("NumberOfRecordsNeeded", fmt.Errorf(`value '%v' must be less than '21'`, this.NumberOfRecordsNeeded))
	}
	if len(this.Emails) > 20 {
		return github_com_mwitkow_go_proto_validators.FieldError("Emails", fmt.Errorf(`value '%v' must contain at most 20 elements`, this.Emails))
	}
	for _, item := range this.Emails {
		if !_regex_CreateMoodRequest_Emails.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Emails", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9.!#$%%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"`, item))
		}
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("Emails", fmt.Errorf(`value '%v' must not be an empty string`, item))
		}
	}
	return nil
}

var _regex_CreateMoodResponse_MoodAccessCode = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *CreateMoodResponse) Validate() error {
	if !(this.MoodId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodId", fmt.Errorf(`value '%v' must be greater than '0'`, this.MoodId))
	}
	if !_regex_CreateMoodResponse_MoodAccessCode.MatchString(this.MoodAccessCode) {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodAccessCode", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.MoodAccessCode))
	}
	if this.MoodAccessCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MoodAccessCode", fmt.Errorf(`value '%v' must not be an empty string`, this.MoodAccessCode))
	}
	return nil
}