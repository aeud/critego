package critego

import (
	bigquery "github.com/aeud/go_google_bigquery"
)

func GetReportBQSchema() *bigquery.BQSchema {
	return bigquery.NewBQSchema([]*bigquery.BQField{
		bigquery.NewBQFieldWithNested("account", "account", bigquery.NewBQSchema([]*bigquery.BQField{
			bigquery.NewBQField("advertiserName", "STRING", "advertiserName"),
			bigquery.NewBQField("email", "STRING", "email"),
			bigquery.NewBQField("currency", "STRING", "currency"),
			bigquery.NewBQField("timezone", "STRING", "timezone"),
			bigquery.NewBQField("country", "STRING", "country"),
		})),
		bigquery.NewBQField("campaignID", "INTEGER", "campaignID"),
		bigquery.NewBQField("dateTimePosix", "STRING", "dateTimePosix"),
		bigquery.NewBQField("dateTime", "STRING", "dateTime"),
		bigquery.NewBQField("categoryID", "INTEGER", "categoryID"),
		bigquery.NewBQField("categoryName", "STRING", "categoryName"),
		bigquery.NewBQField("click", "INTEGER", "click"),
		bigquery.NewBQField("impressions", "INTEGER", "impressions"),
		bigquery.NewBQField("CTR", "FLOAT", "CTR"),
		bigquery.NewBQField("revcpc", "FLOAT", "revcpc"),
		bigquery.NewBQField("ecpm", "FLOAT", "ecpm"),
		bigquery.NewBQField("cost", "FLOAT", "cost"),
		bigquery.NewBQField("sales", "FLOAT", "sales"),
		bigquery.NewBQField("convRate", "FLOAT", "convRate"),
		bigquery.NewBQField("orderValue", "FLOAT", "orderValue"),
		bigquery.NewBQField("salesPostView", "INTEGER", "salesPostView"),
		bigquery.NewBQField("convRatePostView", "FLOAT", "convRatePostView"),
		bigquery.NewBQField("orderValuePostView", "FLOAT", "orderValuePostView"),
		bigquery.NewBQField("overallCompetitionWin", "FLOAT", "overallCompetitionWin"),
		bigquery.NewBQField("costPerOrder", "FLOAT", "costPerOrder"),
	})
}

func GetCampaignBQSchema() *bigquery.BQSchema {
	return bigquery.NewBQSchema([]*bigquery.BQField{
		bigquery.NewBQFieldWithNested("account", "account", bigquery.NewBQSchema([]*bigquery.BQField{
			bigquery.NewBQField("advertiserName", "STRING", "advertiserName"),
			bigquery.NewBQField("email", "STRING", "email"),
			bigquery.NewBQField("currency", "STRING", "currency"),
			bigquery.NewBQField("timezone", "STRING", "timezone"),
			bigquery.NewBQField("country", "STRING", "country"),
		})),
		bigquery.NewBQField("campaignID", "INTEGER", "campaignID"),
		bigquery.NewBQField("campaignName", "STRING", "campaignName"),
		bigquery.NewBQFieldWithNested("campaignBid", "campaignBid", bigquery.NewBQSchema([]*bigquery.BQField{
			bigquery.NewBQField("biddingStrategy", "STRING", "biddingStrategy"),
			bigquery.NewBQField("cpcBid", "FLOAT", "cpcBid"),
		})),
		bigquery.NewBQField("budgetID", "INTEGER", "budgetID"),
		bigquery.NewBQField("remainingDays", "INTEGER", "remainingDays"),
		bigquery.NewBQField("status", "STRING", "status"),
		bigquery.NewBQFieldWithRepeated("categoryBids", "categoryBids", bigquery.NewBQSchema([]*bigquery.BQField{
			bigquery.NewBQField("campaignCategoryUID", "INTEGER", "campaignCategoryUID"),
			bigquery.NewBQField("campaignID", "INTEGER", "campaignID"),
			bigquery.NewBQField("categoryID", "INTEGER", "categoryID"),
			bigquery.NewBQField("selected", "BOOLEAN", "selected"),
			bigquery.NewBQFieldWithNested("bidInformation", "bidInformation", bigquery.NewBQSchema([]*bigquery.BQField{
				bigquery.NewBQField("biddingStrategy", "STRING", "biddingStrategy"),
				bigquery.NewBQField("cpcBid", "FLOAT", "cpcBid"),
			})),
		})),
	})
}

func GetCategoryBQSchema() *bigquery.BQSchema {
	return bigquery.NewBQSchema([]*bigquery.BQField{
		bigquery.NewBQFieldWithNested("account", "account", bigquery.NewBQSchema([]*bigquery.BQField{
			bigquery.NewBQField("advertiserName", "STRING", "advertiserName"),
			bigquery.NewBQField("email", "STRING", "email"),
			bigquery.NewBQField("currency", "STRING", "currency"),
			bigquery.NewBQField("timezone", "STRING", "timezone"),
			bigquery.NewBQField("country", "STRING", "country"),
		})),
		bigquery.NewBQField("categoryID", "INTEGER", "categoryID"),
		bigquery.NewBQField("categoryName", "STRING", "categoryName"),
		bigquery.NewBQField("avgPrice", "FLOAT", "avgPrice"),
		bigquery.NewBQField("numberOfProducts", "INTEGER", "numberOfProducts"),
		bigquery.NewBQField("selected", "BOOLEAN", "selected"),
	})
}
