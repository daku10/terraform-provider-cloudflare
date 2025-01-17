package email_routing_address

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *EmailRoutingAddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Email Routing Address](https://developers.cloudflare.com/email-routing/setup/email-routing-addresses/#destination-addresses) resource allows you to manage Cloudflare Email Routing Destination Addresses.
	`),

		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			consts.IDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.IDSchemaDescription,
				Computed:            true,
			},
			"tag": schema.StringAttribute{
				MarkdownDescription: "Destination address identifier.",
				Computed:            true,
			},
			"email": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The contact email address of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"verified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date and time the destination address has been verified. Null means not verified yet.",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date and time the destination address has been created.",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date and time the destination address has been modified.",
			},
		},
	}
}
