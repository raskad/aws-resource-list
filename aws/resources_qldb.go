package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

func getQLDB(config aws.Config) (resources resourceMap) {
	client := qldb.New(config)

	qldbLedgerNames := getQLDBLedgerNames(client)

	resources = resourceMap{
		qLDBLedger: qldbLedgerNames,
	}
	return
}

func getQLDBLedgerNames(client *qldb.Client) (resources []string) {
	req := client.ListLedgersRequest(&qldb.ListLedgersInput{})
	p := qldb.NewListLedgersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Ledgers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}
