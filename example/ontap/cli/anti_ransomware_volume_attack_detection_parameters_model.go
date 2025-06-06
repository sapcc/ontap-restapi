// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sapcc/ontap-restapi/example/ontap/models"
	"github.com/spf13/cobra"
)

// Schema cli for AntiRansomwareVolumeAttackDetectionParameters

// register flags to command
func registerModelAntiRansomwareVolumeAttackDetectionParametersFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileCreateOpRate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileDeleteOpRate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileRenameOpRate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnHighEntropyDataRate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnNeverSeenBeforeFileExtension(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropFileCreateOpRateSurgeNotifyPercent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropFileDeleteOpRateSurgeNotifyPercent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropFileRenameOpRateSurgeNotifyPercent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropHighEntropyDataSurgeNotifyPercent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionCountNotifyThreshold(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionDurationInHours(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntiRansomwareVolumeAttackDetectionParametersPropRelaxingPopularFileExtensions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileCreateOpRate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagBasedOnFileCreateOpRateDescription := `Specifies whether attack detection is based on the file create operations rate. This parameter is valid only for NAS volumes.`

	var flagBasedOnFileCreateOpRateName string
	if cmdPrefix == "" {
		flagBasedOnFileCreateOpRateName = "based_on_file_create_op_rate"
	} else {
		flagBasedOnFileCreateOpRateName = fmt.Sprintf("%v.based_on_file_create_op_rate", cmdPrefix)
	}

	var flagBasedOnFileCreateOpRateDefault bool

	_ = cmd.PersistentFlags().Bool(flagBasedOnFileCreateOpRateName, flagBasedOnFileCreateOpRateDefault, flagBasedOnFileCreateOpRateDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileDeleteOpRate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagBasedOnFileDeleteOpRateDescription := `Specifies whether attack detection is based on the file delete operations rate. This parameter is valid only for NAS volumes.`

	var flagBasedOnFileDeleteOpRateName string
	if cmdPrefix == "" {
		flagBasedOnFileDeleteOpRateName = "based_on_file_delete_op_rate"
	} else {
		flagBasedOnFileDeleteOpRateName = fmt.Sprintf("%v.based_on_file_delete_op_rate", cmdPrefix)
	}

	var flagBasedOnFileDeleteOpRateDefault bool

	_ = cmd.PersistentFlags().Bool(flagBasedOnFileDeleteOpRateName, flagBasedOnFileDeleteOpRateDefault, flagBasedOnFileDeleteOpRateDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileRenameOpRate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagBasedOnFileRenameOpRateDescription := `Specifies whether attack detection is based on the file rename operations rate. This parameter is valid only for NAS volumes.`

	var flagBasedOnFileRenameOpRateName string
	if cmdPrefix == "" {
		flagBasedOnFileRenameOpRateName = "based_on_file_rename_op_rate"
	} else {
		flagBasedOnFileRenameOpRateName = fmt.Sprintf("%v.based_on_file_rename_op_rate", cmdPrefix)
	}

	var flagBasedOnFileRenameOpRateDefault bool

	_ = cmd.PersistentFlags().Bool(flagBasedOnFileRenameOpRateName, flagBasedOnFileRenameOpRateDefault, flagBasedOnFileRenameOpRateDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnHighEntropyDataRate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagBasedOnHighEntropyDataRateDescription := `Specifies whether a high entropy data rate should be considered for attack detection.`

	var flagBasedOnHighEntropyDataRateName string
	if cmdPrefix == "" {
		flagBasedOnHighEntropyDataRateName = "based_on_high_entropy_data_rate"
	} else {
		flagBasedOnHighEntropyDataRateName = fmt.Sprintf("%v.based_on_high_entropy_data_rate", cmdPrefix)
	}

	var flagBasedOnHighEntropyDataRateDefault bool = true

	_ = cmd.PersistentFlags().Bool(flagBasedOnHighEntropyDataRateName, flagBasedOnHighEntropyDataRateDefault, flagBasedOnHighEntropyDataRateDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropBasedOnNeverSeenBeforeFileExtension(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagBasedOnNeverSeenBeforeFileExtensionDescription := `Specifies whether file extensions never seen before should be considered for attack detection.`

	var flagBasedOnNeverSeenBeforeFileExtensionName string
	if cmdPrefix == "" {
		flagBasedOnNeverSeenBeforeFileExtensionName = "based_on_never_seen_before_file_extension"
	} else {
		flagBasedOnNeverSeenBeforeFileExtensionName = fmt.Sprintf("%v.based_on_never_seen_before_file_extension", cmdPrefix)
	}

	var flagBasedOnNeverSeenBeforeFileExtensionDefault bool = true

	_ = cmd.PersistentFlags().Bool(flagBasedOnNeverSeenBeforeFileExtensionName, flagBasedOnNeverSeenBeforeFileExtensionDefault, flagBasedOnNeverSeenBeforeFileExtensionDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropFileCreateOpRateSurgeNotifyPercent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFileCreateOpRateSurgeNotifyPercentDescription := `Specifies the percentage of surge in the file create rate up to which it is considered normal behavior.`

	var flagFileCreateOpRateSurgeNotifyPercentName string
	if cmdPrefix == "" {
		flagFileCreateOpRateSurgeNotifyPercentName = "file_create_op_rate_surge_notify_percent"
	} else {
		flagFileCreateOpRateSurgeNotifyPercentName = fmt.Sprintf("%v.file_create_op_rate_surge_notify_percent", cmdPrefix)
	}

	var flagFileCreateOpRateSurgeNotifyPercentDefault int64

	_ = cmd.PersistentFlags().Int64(flagFileCreateOpRateSurgeNotifyPercentName, flagFileCreateOpRateSurgeNotifyPercentDefault, flagFileCreateOpRateSurgeNotifyPercentDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropFileDeleteOpRateSurgeNotifyPercent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFileDeleteOpRateSurgeNotifyPercentDescription := `Specifies the percentage of surge in the file delete rate up to which it is considered normal behavior.`

	var flagFileDeleteOpRateSurgeNotifyPercentName string
	if cmdPrefix == "" {
		flagFileDeleteOpRateSurgeNotifyPercentName = "file_delete_op_rate_surge_notify_percent"
	} else {
		flagFileDeleteOpRateSurgeNotifyPercentName = fmt.Sprintf("%v.file_delete_op_rate_surge_notify_percent", cmdPrefix)
	}

	var flagFileDeleteOpRateSurgeNotifyPercentDefault int64 = 100

	_ = cmd.PersistentFlags().Int64(flagFileDeleteOpRateSurgeNotifyPercentName, flagFileDeleteOpRateSurgeNotifyPercentDefault, flagFileDeleteOpRateSurgeNotifyPercentDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropFileRenameOpRateSurgeNotifyPercent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFileRenameOpRateSurgeNotifyPercentDescription := `Specifies the percent of surge in the file rename rate up to which it is considered normal behavior.`

	var flagFileRenameOpRateSurgeNotifyPercentName string
	if cmdPrefix == "" {
		flagFileRenameOpRateSurgeNotifyPercentName = "file_rename_op_rate_surge_notify_percent"
	} else {
		flagFileRenameOpRateSurgeNotifyPercentName = fmt.Sprintf("%v.file_rename_op_rate_surge_notify_percent", cmdPrefix)
	}

	var flagFileRenameOpRateSurgeNotifyPercentDefault int64 = 100

	_ = cmd.PersistentFlags().Int64(flagFileRenameOpRateSurgeNotifyPercentName, flagFileRenameOpRateSurgeNotifyPercentDefault, flagFileRenameOpRateSurgeNotifyPercentDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropHighEntropyDataSurgeNotifyPercent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagHighEntropyDataSurgeNotifyPercentDescription := `Specifies the percentage of surge in high entropy data up to which it is considered as normal behavior. For example, if the usual high entropy data rate in the volume is 5% and if this parameter is set to 100%, it will be considered as an unusual surge if the high entropy data rate of the volume exceeds 10% at any time. Similarly, if this parameter is set to 400%, it will be considered as an unusual surge if the high entropy data rate of the volume exceeds 25%, and so on.`

	var flagHighEntropyDataSurgeNotifyPercentName string
	if cmdPrefix == "" {
		flagHighEntropyDataSurgeNotifyPercentName = "high_entropy_data_surge_notify_percent"
	} else {
		flagHighEntropyDataSurgeNotifyPercentName = fmt.Sprintf("%v.high_entropy_data_surge_notify_percent", cmdPrefix)
	}

	var flagHighEntropyDataSurgeNotifyPercentDefault int64 = 100

	_ = cmd.PersistentFlags().Int64(flagHighEntropyDataSurgeNotifyPercentName, flagHighEntropyDataSurgeNotifyPercentDefault, flagHighEntropyDataSurgeNotifyPercentDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionCountNotifyThreshold(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNeverSeenBeforeFileExtensionCountNotifyThresholdDescription := `Specifies the number of files found with a never seen before file extension up to which it is considered normal behavior.`

	var flagNeverSeenBeforeFileExtensionCountNotifyThresholdName string
	if cmdPrefix == "" {
		flagNeverSeenBeforeFileExtensionCountNotifyThresholdName = "never_seen_before_file_extension_count_notify_threshold"
	} else {
		flagNeverSeenBeforeFileExtensionCountNotifyThresholdName = fmt.Sprintf("%v.never_seen_before_file_extension_count_notify_threshold", cmdPrefix)
	}

	var flagNeverSeenBeforeFileExtensionCountNotifyThresholdDefault int64 = 20

	_ = cmd.PersistentFlags().Int64(flagNeverSeenBeforeFileExtensionCountNotifyThresholdName, flagNeverSeenBeforeFileExtensionCountNotifyThresholdDefault, flagNeverSeenBeforeFileExtensionCountNotifyThresholdDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionDurationInHours(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNeverSeenBeforeFileExtensionDurationInHoursDescription := `Specifies the duration within which the specified number of files found with never seen before file extensions is considered normal behavior.`

	var flagNeverSeenBeforeFileExtensionDurationInHoursName string
	if cmdPrefix == "" {
		flagNeverSeenBeforeFileExtensionDurationInHoursName = "never_seen_before_file_extension_duration_in_hours"
	} else {
		flagNeverSeenBeforeFileExtensionDurationInHoursName = fmt.Sprintf("%v.never_seen_before_file_extension_duration_in_hours", cmdPrefix)
	}

	var flagNeverSeenBeforeFileExtensionDurationInHoursDefault int64 = 24

	_ = cmd.PersistentFlags().Int64(flagNeverSeenBeforeFileExtensionDurationInHoursName, flagNeverSeenBeforeFileExtensionDurationInHoursDefault, flagNeverSeenBeforeFileExtensionDurationInHoursDescription)

	return nil
}

func registerAntiRansomwareVolumeAttackDetectionParametersPropRelaxingPopularFileExtensions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagRelaxingPopularFileExtensionsDescription := `Specifies whether popular file extensions should be relaxed from being treated as a suspect for the attack. Some popular file extensions are .txt, .pdf, and so on.`

	var flagRelaxingPopularFileExtensionsName string
	if cmdPrefix == "" {
		flagRelaxingPopularFileExtensionsName = "relaxing_popular_file_extensions"
	} else {
		flagRelaxingPopularFileExtensionsName = fmt.Sprintf("%v.relaxing_popular_file_extensions", cmdPrefix)
	}

	var flagRelaxingPopularFileExtensionsDefault bool = true

	_ = cmd.PersistentFlags().Bool(flagRelaxingPopularFileExtensionsName, flagRelaxingPopularFileExtensionsDefault, flagRelaxingPopularFileExtensionsDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAntiRansomwareVolumeAttackDetectionParametersFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, BasedOnFileCreateOpRateAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileCreateOpRateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || BasedOnFileCreateOpRateAdded

	err, BasedOnFileDeleteOpRateAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileDeleteOpRateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || BasedOnFileDeleteOpRateAdded

	err, BasedOnFileRenameOpRateAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileRenameOpRateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || BasedOnFileRenameOpRateAdded

	err, BasedOnHighEntropyDataRateAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnHighEntropyDataRateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || BasedOnHighEntropyDataRateAdded

	err, BasedOnNeverSeenBeforeFileExtensionAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnNeverSeenBeforeFileExtensionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || BasedOnNeverSeenBeforeFileExtensionAdded

	err, FileCreateOpRateSurgeNotifyPercentAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropFileCreateOpRateSurgeNotifyPercentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FileCreateOpRateSurgeNotifyPercentAdded

	err, FileDeleteOpRateSurgeNotifyPercentAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropFileDeleteOpRateSurgeNotifyPercentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FileDeleteOpRateSurgeNotifyPercentAdded

	err, FileRenameOpRateSurgeNotifyPercentAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropFileRenameOpRateSurgeNotifyPercentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FileRenameOpRateSurgeNotifyPercentAdded

	err, HighEntropyDataSurgeNotifyPercentAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropHighEntropyDataSurgeNotifyPercentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || HighEntropyDataSurgeNotifyPercentAdded

	err, NeverSeenBeforeFileExtensionCountNotifyThresholdAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionCountNotifyThresholdFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NeverSeenBeforeFileExtensionCountNotifyThresholdAdded

	err, NeverSeenBeforeFileExtensionDurationInHoursAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionDurationInHoursFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NeverSeenBeforeFileExtensionDurationInHoursAdded

	err, RelaxingPopularFileExtensionsAdded := retrieveAntiRansomwareVolumeAttackDetectionParametersPropRelaxingPopularFileExtensionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RelaxingPopularFileExtensionsAdded

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileCreateOpRateFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagBasedOnFileCreateOpRateName := fmt.Sprintf("%v.based_on_file_create_op_rate", cmdPrefix)
	if cmd.Flags().Changed(flagBasedOnFileCreateOpRateName) {

		var flagBasedOnFileCreateOpRateName string
		if cmdPrefix == "" {
			flagBasedOnFileCreateOpRateName = "based_on_file_create_op_rate"
		} else {
			flagBasedOnFileCreateOpRateName = fmt.Sprintf("%v.based_on_file_create_op_rate", cmdPrefix)
		}

		flagBasedOnFileCreateOpRateValue, err := cmd.Flags().GetBool(flagBasedOnFileCreateOpRateName)
		if err != nil {
			return err, false
		}
		m.BasedOnFileCreateOpRate = &flagBasedOnFileCreateOpRateValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileDeleteOpRateFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagBasedOnFileDeleteOpRateName := fmt.Sprintf("%v.based_on_file_delete_op_rate", cmdPrefix)
	if cmd.Flags().Changed(flagBasedOnFileDeleteOpRateName) {

		var flagBasedOnFileDeleteOpRateName string
		if cmdPrefix == "" {
			flagBasedOnFileDeleteOpRateName = "based_on_file_delete_op_rate"
		} else {
			flagBasedOnFileDeleteOpRateName = fmt.Sprintf("%v.based_on_file_delete_op_rate", cmdPrefix)
		}

		flagBasedOnFileDeleteOpRateValue, err := cmd.Flags().GetBool(flagBasedOnFileDeleteOpRateName)
		if err != nil {
			return err, false
		}
		m.BasedOnFileDeleteOpRate = &flagBasedOnFileDeleteOpRateValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnFileRenameOpRateFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagBasedOnFileRenameOpRateName := fmt.Sprintf("%v.based_on_file_rename_op_rate", cmdPrefix)
	if cmd.Flags().Changed(flagBasedOnFileRenameOpRateName) {

		var flagBasedOnFileRenameOpRateName string
		if cmdPrefix == "" {
			flagBasedOnFileRenameOpRateName = "based_on_file_rename_op_rate"
		} else {
			flagBasedOnFileRenameOpRateName = fmt.Sprintf("%v.based_on_file_rename_op_rate", cmdPrefix)
		}

		flagBasedOnFileRenameOpRateValue, err := cmd.Flags().GetBool(flagBasedOnFileRenameOpRateName)
		if err != nil {
			return err, false
		}
		m.BasedOnFileRenameOpRate = &flagBasedOnFileRenameOpRateValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnHighEntropyDataRateFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagBasedOnHighEntropyDataRateName := fmt.Sprintf("%v.based_on_high_entropy_data_rate", cmdPrefix)
	if cmd.Flags().Changed(flagBasedOnHighEntropyDataRateName) {

		var flagBasedOnHighEntropyDataRateName string
		if cmdPrefix == "" {
			flagBasedOnHighEntropyDataRateName = "based_on_high_entropy_data_rate"
		} else {
			flagBasedOnHighEntropyDataRateName = fmt.Sprintf("%v.based_on_high_entropy_data_rate", cmdPrefix)
		}

		flagBasedOnHighEntropyDataRateValue, err := cmd.Flags().GetBool(flagBasedOnHighEntropyDataRateName)
		if err != nil {
			return err, false
		}
		m.BasedOnHighEntropyDataRate = &flagBasedOnHighEntropyDataRateValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropBasedOnNeverSeenBeforeFileExtensionFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagBasedOnNeverSeenBeforeFileExtensionName := fmt.Sprintf("%v.based_on_never_seen_before_file_extension", cmdPrefix)
	if cmd.Flags().Changed(flagBasedOnNeverSeenBeforeFileExtensionName) {

		var flagBasedOnNeverSeenBeforeFileExtensionName string
		if cmdPrefix == "" {
			flagBasedOnNeverSeenBeforeFileExtensionName = "based_on_never_seen_before_file_extension"
		} else {
			flagBasedOnNeverSeenBeforeFileExtensionName = fmt.Sprintf("%v.based_on_never_seen_before_file_extension", cmdPrefix)
		}

		flagBasedOnNeverSeenBeforeFileExtensionValue, err := cmd.Flags().GetBool(flagBasedOnNeverSeenBeforeFileExtensionName)
		if err != nil {
			return err, false
		}
		m.BasedOnNeverSeenBeforeFileExtension = &flagBasedOnNeverSeenBeforeFileExtensionValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropFileCreateOpRateSurgeNotifyPercentFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFileCreateOpRateSurgeNotifyPercentName := fmt.Sprintf("%v.file_create_op_rate_surge_notify_percent", cmdPrefix)
	if cmd.Flags().Changed(flagFileCreateOpRateSurgeNotifyPercentName) {

		var flagFileCreateOpRateSurgeNotifyPercentName string
		if cmdPrefix == "" {
			flagFileCreateOpRateSurgeNotifyPercentName = "file_create_op_rate_surge_notify_percent"
		} else {
			flagFileCreateOpRateSurgeNotifyPercentName = fmt.Sprintf("%v.file_create_op_rate_surge_notify_percent", cmdPrefix)
		}

		flagFileCreateOpRateSurgeNotifyPercentValue, err := cmd.Flags().GetInt64(flagFileCreateOpRateSurgeNotifyPercentName)
		if err != nil {
			return err, false
		}
		m.FileCreateOpRateSurgeNotifyPercent = &flagFileCreateOpRateSurgeNotifyPercentValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropFileDeleteOpRateSurgeNotifyPercentFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFileDeleteOpRateSurgeNotifyPercentName := fmt.Sprintf("%v.file_delete_op_rate_surge_notify_percent", cmdPrefix)
	if cmd.Flags().Changed(flagFileDeleteOpRateSurgeNotifyPercentName) {

		var flagFileDeleteOpRateSurgeNotifyPercentName string
		if cmdPrefix == "" {
			flagFileDeleteOpRateSurgeNotifyPercentName = "file_delete_op_rate_surge_notify_percent"
		} else {
			flagFileDeleteOpRateSurgeNotifyPercentName = fmt.Sprintf("%v.file_delete_op_rate_surge_notify_percent", cmdPrefix)
		}

		flagFileDeleteOpRateSurgeNotifyPercentValue, err := cmd.Flags().GetInt64(flagFileDeleteOpRateSurgeNotifyPercentName)
		if err != nil {
			return err, false
		}
		m.FileDeleteOpRateSurgeNotifyPercent = &flagFileDeleteOpRateSurgeNotifyPercentValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropFileRenameOpRateSurgeNotifyPercentFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFileRenameOpRateSurgeNotifyPercentName := fmt.Sprintf("%v.file_rename_op_rate_surge_notify_percent", cmdPrefix)
	if cmd.Flags().Changed(flagFileRenameOpRateSurgeNotifyPercentName) {

		var flagFileRenameOpRateSurgeNotifyPercentName string
		if cmdPrefix == "" {
			flagFileRenameOpRateSurgeNotifyPercentName = "file_rename_op_rate_surge_notify_percent"
		} else {
			flagFileRenameOpRateSurgeNotifyPercentName = fmt.Sprintf("%v.file_rename_op_rate_surge_notify_percent", cmdPrefix)
		}

		flagFileRenameOpRateSurgeNotifyPercentValue, err := cmd.Flags().GetInt64(flagFileRenameOpRateSurgeNotifyPercentName)
		if err != nil {
			return err, false
		}
		m.FileRenameOpRateSurgeNotifyPercent = &flagFileRenameOpRateSurgeNotifyPercentValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropHighEntropyDataSurgeNotifyPercentFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagHighEntropyDataSurgeNotifyPercentName := fmt.Sprintf("%v.high_entropy_data_surge_notify_percent", cmdPrefix)
	if cmd.Flags().Changed(flagHighEntropyDataSurgeNotifyPercentName) {

		var flagHighEntropyDataSurgeNotifyPercentName string
		if cmdPrefix == "" {
			flagHighEntropyDataSurgeNotifyPercentName = "high_entropy_data_surge_notify_percent"
		} else {
			flagHighEntropyDataSurgeNotifyPercentName = fmt.Sprintf("%v.high_entropy_data_surge_notify_percent", cmdPrefix)
		}

		flagHighEntropyDataSurgeNotifyPercentValue, err := cmd.Flags().GetInt64(flagHighEntropyDataSurgeNotifyPercentName)
		if err != nil {
			return err, false
		}
		m.HighEntropyDataSurgeNotifyPercent = &flagHighEntropyDataSurgeNotifyPercentValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionCountNotifyThresholdFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNeverSeenBeforeFileExtensionCountNotifyThresholdName := fmt.Sprintf("%v.never_seen_before_file_extension_count_notify_threshold", cmdPrefix)
	if cmd.Flags().Changed(flagNeverSeenBeforeFileExtensionCountNotifyThresholdName) {

		var flagNeverSeenBeforeFileExtensionCountNotifyThresholdName string
		if cmdPrefix == "" {
			flagNeverSeenBeforeFileExtensionCountNotifyThresholdName = "never_seen_before_file_extension_count_notify_threshold"
		} else {
			flagNeverSeenBeforeFileExtensionCountNotifyThresholdName = fmt.Sprintf("%v.never_seen_before_file_extension_count_notify_threshold", cmdPrefix)
		}

		flagNeverSeenBeforeFileExtensionCountNotifyThresholdValue, err := cmd.Flags().GetInt64(flagNeverSeenBeforeFileExtensionCountNotifyThresholdName)
		if err != nil {
			return err, false
		}
		m.NeverSeenBeforeFileExtensionCountNotifyThreshold = &flagNeverSeenBeforeFileExtensionCountNotifyThresholdValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropNeverSeenBeforeFileExtensionDurationInHoursFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNeverSeenBeforeFileExtensionDurationInHoursName := fmt.Sprintf("%v.never_seen_before_file_extension_duration_in_hours", cmdPrefix)
	if cmd.Flags().Changed(flagNeverSeenBeforeFileExtensionDurationInHoursName) {

		var flagNeverSeenBeforeFileExtensionDurationInHoursName string
		if cmdPrefix == "" {
			flagNeverSeenBeforeFileExtensionDurationInHoursName = "never_seen_before_file_extension_duration_in_hours"
		} else {
			flagNeverSeenBeforeFileExtensionDurationInHoursName = fmt.Sprintf("%v.never_seen_before_file_extension_duration_in_hours", cmdPrefix)
		}

		flagNeverSeenBeforeFileExtensionDurationInHoursValue, err := cmd.Flags().GetInt64(flagNeverSeenBeforeFileExtensionDurationInHoursName)
		if err != nil {
			return err, false
		}
		m.NeverSeenBeforeFileExtensionDurationInHours = &flagNeverSeenBeforeFileExtensionDurationInHoursValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntiRansomwareVolumeAttackDetectionParametersPropRelaxingPopularFileExtensionsFlags(depth int, m *models.AntiRansomwareVolumeAttackDetectionParameters, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagRelaxingPopularFileExtensionsName := fmt.Sprintf("%v.relaxing_popular_file_extensions", cmdPrefix)
	if cmd.Flags().Changed(flagRelaxingPopularFileExtensionsName) {

		var flagRelaxingPopularFileExtensionsName string
		if cmdPrefix == "" {
			flagRelaxingPopularFileExtensionsName = "relaxing_popular_file_extensions"
		} else {
			flagRelaxingPopularFileExtensionsName = fmt.Sprintf("%v.relaxing_popular_file_extensions", cmdPrefix)
		}

		flagRelaxingPopularFileExtensionsValue, err := cmd.Flags().GetBool(flagRelaxingPopularFileExtensionsName)
		if err != nil {
			return err, false
		}
		m.RelaxingPopularFileExtensions = &flagRelaxingPopularFileExtensionsValue

		retAdded = true
	}

	return nil, retAdded
}
