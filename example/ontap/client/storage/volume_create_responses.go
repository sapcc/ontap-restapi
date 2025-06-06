// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/ontap-restapi/example/ontap/models"
)

// VolumeCreateReader is a Reader for the VolumeCreate structure.
type VolumeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVolumeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewVolumeCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeCreateCreated creates a VolumeCreateCreated with default headers values
func NewVolumeCreateCreated() *VolumeCreateCreated {
	return &VolumeCreateCreated{}
}

/*
VolumeCreateCreated describes a response with status code 201, with default header values.

Created
*/
type VolumeCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this volume create created response has a 2xx status code
func (o *VolumeCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume create created response has a 3xx status code
func (o *VolumeCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume create created response has a 4xx status code
func (o *VolumeCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume create created response has a 5xx status code
func (o *VolumeCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this volume create created response a status code equal to that given
func (o *VolumeCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the volume create created response
func (o *VolumeCreateCreated) Code() int {
	return 201
}

func (o *VolumeCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes][%d] volumeCreateCreated %s", 201, payload)
}

func (o *VolumeCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes][%d] volumeCreateCreated %s", 201, payload)
}

func (o *VolumeCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *VolumeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateAccepted creates a VolumeCreateAccepted with default headers values
func NewVolumeCreateAccepted() *VolumeCreateAccepted {
	return &VolumeCreateAccepted{}
}

/*
VolumeCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type VolumeCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this volume create accepted response has a 2xx status code
func (o *VolumeCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume create accepted response has a 3xx status code
func (o *VolumeCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume create accepted response has a 4xx status code
func (o *VolumeCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume create accepted response has a 5xx status code
func (o *VolumeCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this volume create accepted response a status code equal to that given
func (o *VolumeCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the volume create accepted response
func (o *VolumeCreateAccepted) Code() int {
	return 202
}

func (o *VolumeCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes][%d] volumeCreateAccepted %s", 202, payload)
}

func (o *VolumeCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes][%d] volumeCreateAccepted %s", 202, payload)
}

func (o *VolumeCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *VolumeCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateDefault creates a VolumeCreateDefault with default headers values
func NewVolumeCreateDefault(code int) *VolumeCreateDefault {
	return &VolumeCreateDefault{
		_statusCode: code,
	}
}

/*
	VolumeCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 460770 | \[Job \"jobid\"\] Job failed. |
| 524561 | Unable to get the information for the specified SVM's root volume because of the specified reason. |
| 524597 | Failed to create the clone volume on the node. |
| 524597 | Failed to create clone volume on the node. |
| 524601 | Parent volume \"name\" not found. |
| 524819 | Failed to create the clone volume. |
| 787140 | One of \"aggregates.uuid\", \"aggregates.name\", or \"style\" must be provided. |
| 787141 | The specified \"aggregates.name\" and \"aggregates.uuid\" refer to different aggregates. |
| 917508 | The target aggregate could not be found. |
| 917526 | The volume name specified is a duplicate. |
| 917534 | The value specified for field \"-size\" is too small. Update the field \"-size\" with the minimum size allowed and retry. |
| 917536 | Job completed with specified error. |
| 917551 | Specified snapshot policy not found. |
| 917597 | Provisioning a volume on a root aggregate is not supported. |
| 917775 | Volume \\\"{0}:{1}\\\" is not online. |
| 917829 | Volume autosize grow threshold must be larger than autosize shrink threshold. |
| 917831 | Volume minimum autosize must be smaller than the maximum autosize. |
| 917835 | Maximum allowed snapshot.reserve_percent value during a volume creation is 90. Use PATCH to set it to a higher value after the volume has been created. |
| 917886 | Volume name is too long. It must be 203 characters or less. |
| 917887 | The first character of the volume name must be a letter or underscore. |
| 918191 | Flexvol tiering min cooling days requires an effective cluster version of ONTAP 9.4 or later. |
| 918194 | Tiering min cooling days not supported for SVMDR. |
| 918195 | Tiering min cooling days not supported for non data volumes. |
| 918196 | Tiering min cooling days not allowed for the provided tiering policy. |
| 918215 | FlexGroup tiering min cooling days requires an effective cluster version of ONTAP 9.5 or later. |
| 918232 | The specified volume identifier fields must be provided. |
| 918233 | The target field cannot be specified for this operation. |
| 918236 | The specified \"parent_volume.uuid\" and \"parent_volume.name\" do not refer to the same volume. |
| 918240 | The target style is an invalid volume style. |
| 918241 | The target style is an unsupported volume style for volume creation. |
| 918242 | When creating a flexible volume, exactly one aggregate must be specified via either \"aggregates.name\" or \"aggregates.uuid\". |
| 918243 | The specified snapshot UUID is not correct for the specified snapshot name. |
| 918244 | Invalid \"volume.type\" for clone volume. |
| 918247 | Specifying a value is not valid for a volume FlexClone creation. |
| 918252 | \"nas.path\" is invalid. |
| 918271 | Failed to get valid export policy. |
| 918290 | cloud retrieval policy requires an effective cluster version of 9.8 or later. |
| 918291 | Invalid volume cloud retrieval policy for the provided tiering policy. |
| 918292 | cloud retrieval policy not supported for non data volume. |
| 918334 | Cache cloud retrieval policy requires an effective cluster version of 9.14 or later. |
| 918521 | The volume maximum autosize must be smaller than or equal to the maximum volume size. |
| 918524 | Volume minimum autosize must be less than or equal to the current volume size. |
| 918652 | \"error\" is an invalid value for field \"-state\". Valid values are \"online\", \"offline\" and \"restricted\". |
| 918701 | The specified operation on the volume endpoint is not supported on this platform. |
| 1638400 | Failed to retrieve snapshot information. |
| 1638587 | The specified snapshot_policy.name and snapshot_policy.uuid refer to different snapshot policies. |
| 1638593 | Operation failed because "snapdb" is disabled. |
| 1638624 | Internal error. Failed to lookup snapshot tags for \\\"{0}\\\". |
| 2621462 | The target SVM does not exist. |
| 2621706 | The specified \"svm.uuid\" and \"svm.name\" do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either \"svm.name\" or \"svm.uuid\" must be supplied. |
| 6881654 | Disabling storage efficiency features is not allowed on this platform. |
| 13107245 | The specified FlexGroup volume layout exceeds the maximum number of volumes on node. |
| 13107307 | specifying this parameter when creating a flexible volume is not supported. |
| 13107326 | Request to provision FlexGroup volume \"name\" failed because the requested size is greater than the maximum size. |
| 13107341 | At least one valid aggregate assigned to SVM \\\"{0}\\\" with sufficient space and a homogeneous storage type is required on every node on the cluster to provision a FlexGroup volume. |
| 13107349 | Operation is only supported on flexible volumes and FlexGroup volumes. |
| 13107405 | Another volume is currently being created with the name \"name\" in SVM \"svm.name\". |
| 13107413 | Creating a FlexGroup volume is not supported on All SAN Arrays and systems that support large LUNs. |
| 13109258 | Cannot enable granular data on volume \"name\" in Vserver \"svm.name\". This setting can only be enabled on FlexGroups. |
| 13109260 | Failed to enable granular data on the volume. |
| 13565983 | A value of zero is not supported for \\\"-uid\\\" or \\\"-gid\\\". |
| 65537463 | Volume encryption keys (VEK) cannot be created or deleted for data Vserver \\\"{0}\\\". External key management has been configured for data Vserver \\\"{0}\\\" but ONTAP is not able to encrypt or decrypt with the key manager. Resolve the external key manager key issues at the key manager's portal before creating any new encrypted volumes. |
| 65537529 | Encrypted volumes cannot be created or deleted for Vserver \\\"{0}\\\" as a rekey operation for the vserver is in progress. Try creating the encrypted volume again after some time. If the problem persists, run the rekey operation again after some time. |
| 65537600 | Encrypted volumes cannot be created or deleted for Vserver \\\"{0}\\\" while the enabled keystore configuration is being switched. If a previous attempt to switch the keystore configuration failed, or was interrupted, the system will continue to prevent encrypted volume creation for Vserver \\\"{0}\\\". Use the \\\"security key-manager keystore enable\\\" command to re-run and complete the operation. |
| 65539430 | Cannot create or delete volumes on Vserver \\\"{0}\\\" while the keystore is being initialized. Wait until the keystore is in the active state, and rerun the volume operation. |
| 65539431 | Cannot create or delete volumes on Vserver \\\"{0}\\\" while the keystore is being disabled. |
| 111411205 | File system analytics requires an effective cluster version of 9.8 or later. |
| 111411206 | The specified \"analytics.state\" is invalid. |
| 111411207 | File system analytics cannot be enabled on volumes that contain LUNs. |
| 111411207 | Volume file system analytics is not supported on volumes that contain LUNs. |
| 111411209 | Volume file system analytics is not supported on FlexCache volumes. |
| 111411210 | Volume file system analytics is not supported on audit staging volumes. |
| 111411211 | Volume file system analytics is not supported on object store server volumes. |
| 111411212 | Volume file system analytics is not supported on SnapMirror destination volumes. |
| 111411216 | Enabling or disabling volume file system analytics is not supported on individual FlexGroup constituents. |
| 111411217 | Volume file system analytics is not supported on SnapLock volumes. |
| 111411230 | Volume file system analytics is not supported on volumes that contain NVMe namespaces. |
| 111411241 | Volume file system analytics is not supported on All SAN Array clusters. |
| 111411252 | Failed to enable file system analytics on volume \"name\" in SVM \"svm.name\" because there is insufficient available space. Ensure there is at least \"space.available_percent`\" available space in the volume, then try the operation again. |
| 111411253 | Failed to enable file system analytics on volume \"name\" in SVM \"svm.name\" because there is insufficient available space. Ensure there is at least \"space.available_percent\" available space in all constituents of the FlexGroup, then try the operation again." |
| 111411257 | Failed to enable file system analytics on volume \"name\" in SVM \"svm.name\" because there is insufficient available space. Increase the volume size to \"size\", then try the operation again." |
| 111415208 | Cannot create a clone of volume "name" in SVM "svm.name" because the file system analytics state is not set to "off" or "on". Use REST API GET method "/api/storage/volumes/uuid?fields=analytics.state" to check the analytics state. |
| 124518405 | Volume activity tracking is not supported on volumes that contain LUNs. |
| 124518407 | Volume activity tracking is not supported on FlexCache volumes. |
| 124518408 | Volume activity tracking is not supported on audit staging volumes. |
| 124518409 | Volume activity tracking is not supported on object store server volumes. |
| 124518410 | Volume activity tracking is not supported on SnapMirror destination volumes. |
| 124518411 | Enabling or disabling volume activity tracking is not supported on individual FlexGroup constituents. |
| 124518412 | Volume activity tracking is not supported on SnapLock volumes. |
| 124518414 | Volume activity tracking is not supported on volumes that contain NVMe namespaces. |
| 124518422 | Volume activity tracking is not supported on All SAN Array clusters. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type VolumeCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volume create default response has a 2xx status code
func (o *VolumeCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume create default response has a 3xx status code
func (o *VolumeCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume create default response has a 4xx status code
func (o *VolumeCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume create default response has a 5xx status code
func (o *VolumeCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume create default response a status code equal to that given
func (o *VolumeCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume create default response
func (o *VolumeCreateDefault) Code() int {
	return o._statusCode
}

func (o *VolumeCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes][%d] volume_create default %s", o._statusCode, payload)
}

func (o *VolumeCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes][%d] volume_create default %s", o._statusCode, payload)
}

func (o *VolumeCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
