package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/backup"
)

func getBackup(session *session.Session) (resources resourceMap) {
	client := backup.New(session)

	backupBackupPlanResourceMap := getBackupBackupPlan(client).unwrap(backupBackupPlan)
	backupBackupPlanIDs := backupBackupPlanResourceMap[backupBackupPlan]

	resources = reduce(
		backupBackupPlanResourceMap,
		getBackupBackupSelection(client, backupBackupPlanIDs).unwrap(backupBackupSelection),
		getBackupBackupVault(client).unwrap(backupBackupVault),
	)
	return
}

func getBackupBackupPlan(client *backup.Backup) (r resourceSliceError) {
	r.err = client.ListBackupPlansPages(&backup.ListBackupPlansInput{}, func(page *backup.ListBackupPlansOutput, lastPage bool) bool {
		for _, resource := range page.BackupPlansList {
			r.resources = append(r.resources, *resource.BackupPlanId)
		}
		return true
	})
	return
}

func getBackupBackupSelection(client *backup.Backup, backupBackupPlanIDs []string) (r resourceSliceError) {
	for _, backupBackupPlanID := range backupBackupPlanIDs {
		r.err = client.ListBackupSelectionsPages(&backup.ListBackupSelectionsInput{
			BackupPlanId: aws.String(backupBackupPlanID),
		}, func(page *backup.ListBackupSelectionsOutput, lastPage bool) bool {
			for _, resource := range page.BackupSelectionsList {
				r.resources = append(r.resources, *resource.SelectionId)
			}
			return true
		})
	}
	return
}

func getBackupBackupVault(client *backup.Backup) (r resourceSliceError) {
	r.err = client.ListBackupVaultsPages(&backup.ListBackupVaultsInput{}, func(page *backup.ListBackupVaultsOutput, lastPage bool) bool {
		for _, resource := range page.BackupVaultList {
			r.resources = append(r.resources, *resource.BackupVaultName)
		}
		return true
	})
	return
}
