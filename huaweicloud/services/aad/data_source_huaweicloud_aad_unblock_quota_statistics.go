package aad

import (
	"context"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API AAD GET /v1/unblockservice/{domain_id}/unblock-quota-statistics
func DataSourceAADUnblockQuotaStatistics() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAADUnblockQuotaStatisticsRead,

		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specified the account ID of IAM user.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total unblocking times.",
			},
			"total_unblocking_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total unblocking times.",
			},
			"remaining_unblocking_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The manual unblocking times.",
			},
			"unblocking_quota_today": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The automatic unblocking times.",
			},
			"remaining_unblocking_quota_today": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current blocked IP number.",
			},
		},
	}
}

func dataSourceAADUnblockQuotaStatisticsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg    = meta.(*config.Config)
		region = cfg.GetRegion(d)
		mErr   *multierror.Error
	)

	client, err := cfg.NewServiceClient("aad", region)
	if err != nil {
		return diag.Errorf("error creating AAD client: %s", err)
	}

	httpUrl := "v1/unblockservice/{domain_id}/unblock-quota-statistics"
	httpUrl = strings.ReplaceAll(httpUrl, "{domain_id}", d.Get("domain_id").(string))
	listPath := client.Endpoint + httpUrl
	opt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	requestResp, err := client.Request("GET", listPath, &opt)
	if err != nil {
		return diag.Errorf("error retrieving AAD block statistics: %s", err)
	}

	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Errorf("unable to generate ID: %s", err)
	}
	d.SetId(id)

	mErr = multierror.Append(
		mErr,
		d.Set("type", utils.PathSearch("type", respBody, nil)),
		d.Set("total_unblocking_quota", utils.PathSearch("total_unblocking_quota", respBody, nil)),
		d.Set("remaining_unblocking_quota", utils.PathSearch("remaining_unblocking_quota", respBody, nil)),
		d.Set("unblocking_quota_today", utils.PathSearch("unblocking_quota_today", respBody, nil)),
		d.Set("remaining_unblocking_quota_today", utils.PathSearch("remaining_unblocking_quota_today", respBody, nil)),
	)
	if mErr.ErrorOrNil() != nil {
		return diag.Errorf("error retrieving AAD block statistics: %s", mErr)
	}

	return nil
}
