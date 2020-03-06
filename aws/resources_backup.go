package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func getBackup(config aws.Config) (resources resourceMap) {
	client := backup.New(config)

	backupBackupPlanResourceMap := getBackupBackupPlan(client).unwrap(backupBackupPlan)
	backupBackupPlanIDs := backupBackupPlanResourceMap[backupBackupPlan]

	resources = reduce(
		backupBackupPlanResourceMap,
		getBackupBackupSelection(client, backupBackupPlanIDs).unwrap(backupBackupSelection),
		getBackupBackupVault(client).unwrap(backupBackupVault),
	)
	return
}

func getBackupBackupPlan(client *backup.Client) (r resourceSliceError) {
	req := client.ListBackupPlansRequest(&backup.ListBackupPlansInput{})
	p := backup.NewListBackupPlansPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.BackupPlansList {
			r.resources = append(r.resources, *resource.BackupPlanId)
		}
	}
	r.err = p.Err()
	return
}

func getBackupBackupSelection(client *backup.Client, backupBackupPlanIDs []string) (r resourceSliceError) {
	for _, backupBackupPlanID := range backupBackupPlanIDs {
		req := client.ListBackupSelectionsRequest(&backup.ListBackupSelectionsInput{
			BackupPlanId: aws.String(backupBackupPlanID),
		})
		p := backup.NewListBackupSelectionsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.BackupSelectionsList {
				r.resources = append(r.resources, *resource.SelectionId)
			}
		}
		r.err = p.Err()
	}
	return
}

func getBackupBackupVault(client *backup.Client) (r resourceSliceError) {
	req := client.ListBackupVaultsRequest(&backup.ListBackupVaultsInput{})
	p := backup.NewListBackupVaultsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.BackupVaultList {
			r.resources = append(r.resources, *resource.BackupVaultName)
		}
	}
	r.err = p.Err()
	return
}
