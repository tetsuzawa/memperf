package memperf

type Frames struct {
	Items []Frame `json:"items"`
}

type Frame struct {
	Id                      int64   `json:"id"`
	MediumId                int64   `json:"medium_id"`
	BlockCampaignCategories []int64 `json:"block_campaign_categories"`
	BlockCreativeTags       []int64 `json:"block_creative_tags"`
}
