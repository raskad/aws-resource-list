package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func getBackup(config aws.Config) (resources resourceMap) {
	client := backup.New(config)

	backupBackupPlanIDs := getBackupBackupPlanIDs(client)
	backupBackupSelectionIDs := getBackupBackupSelectionIDs(client, backupBackupPlanIDs)
	backupBackupVaultNames := getBackupBackupVaultNames(client)

	resources = resourceMap{
		backupBackupPlan:      backupBackupPlanIDs,
		backupBackupSelection: backupBackupSelectionIDs,
		backupBackupVault:     backupBackupVaultNames,
	}
	return
}

func getBackupBackupPlanIDs(client *backup.Client) (resources []string) {
	req := client.ListBackupPlansRequest(&backup.ListBackupPlansInput{})
	p := backup.NewListBackupPlansPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.BackupPlansList {
			resources = append(resources, *resource.BackupPlanId)
		}
	}
	return
}

func getBackupBackupSelectionIDs(client *backup.Client, backupBackupPlanIDs []string) (resources []string) {
	for _, backupBackupPlanID := range backupBackupPlanIDs {
		req := client.ListBackupSelectionsRequest(&backup.ListBackupSelectionsInput{
			BackupPlanId: aws.String(backupBackupPlanID),
		})
		p := backup.NewListBackupSelectionsPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.BackupSelectionsList {
				resources = append(resources, *resource.SelectionId)
			}
		}
	}
	return
}

func getBackupBackupVaultNames(client *backup.Client) (resources []string) {
	req := client.ListBackupVaultsRequest(&backup.ListBackupVaultsInput{})
	p := backup.NewListBackupVaultsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.BackupVaultList {
			resources = append(resources, *resource.BackupVaultName)
		}
	}
	return
}
