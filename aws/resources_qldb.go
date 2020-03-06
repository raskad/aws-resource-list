package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

func getQLDB(config aws.Config) (resources resourceMap) {
	client := qldb.New(config)
	resources = reduce(
		getQLDBLedger(client).unwrap(qLDBLedger),
	)
	return
}

func getQLDBLedger(client *qldb.Client) (r resourceSliceError) {
	req := client.ListLedgersRequest(&qldb.ListLedgersInput{})
	p := qldb.NewListLedgersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Ledgers {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}
