package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/backup"
)

func getBackup(session *session.Session) (resources resourceMap) {
	client := backup.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		backupBackupPlan:      getBackupBackupPlan(client),
		backupBackupSelection: getBackupBackupSelection(client),
		backupBackupVault:     getBackupBackupVault(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getBackupBackupPlan(client *backup.Backup) (r resourceSliceError) {
	r.err = client.ListBackupPlansPages(&backup.ListBackupPlansInput{}, func(page *backup.ListBackupPlansOutput, lastPage bool) bool {
		logDebug("List BackupBackupPlan resources page. Remaining pages", page.NextToken)
		for _, resource := range page.BackupPlansList {
			logDebug("Got BackupBackupPlan resource with PhysicalResourceId", *resource.BackupPlanId)
			r.resources = append(r.resources, *resource.BackupPlanId)
		}
		return true
	})
	return
}

func getBackupBackupSelection(client *backup.Backup) (r resourceSliceError) {
	r.err = client.ListBackupSelectionsPages(&backup.ListBackupSelectionsInput{}, func(page *backup.ListBackupSelectionsOutput, lastPage bool) bool {
		logDebug("List BackupBackupSelection resources page. Remaining pages", page.NextToken)
		for _, resource := range page.BackupSelectionsList {
			logDebug("Got BackupBackupSelection resource with PhysicalResourceId", *resource.SelectionId)
			r.resources = append(r.resources, *resource.SelectionId)
		}
		return true
	})
	return
}

func getBackupBackupVault(client *backup.Backup) (r resourceSliceError) {
	r.err = client.ListBackupVaultsPages(&backup.ListBackupVaultsInput{}, func(page *backup.ListBackupVaultsOutput, lastPage bool) bool {
		logDebug("List BackupBackupVault resources page. Remaining pages", page.NextToken)
		for _, resource := range page.BackupVaultList {
			logDebug("Got BackupBackupVault resource with PhysicalResourceId", *resource.BackupVaultName)
			r.resources = append(r.resources, *resource.BackupVaultName)
		}
		return true
	})
	return
}
