package ontap

//go:generate go tool swagger generate cli -f ../../v9.16/ontap_fixed.yml --skip-validation -O volume_create -O volume_collection_get -M volume -M anti_ransomware_volume_attack_detection_parameters -M job_link_response -M job_link -M href -M snapshot_reference -M anti_ransomware_volume_workload -M anti_ransomware_attack_report -M error -M error_arguments -M volume_response -M error_response -M returned_error
