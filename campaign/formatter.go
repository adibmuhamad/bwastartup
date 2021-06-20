package campaign

import "strings"

type CampaignFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	Slug string `json:"slug"`
}

type CampaignUserFormatter struct {
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl string `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}

type CampaignDetailFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	BackerCount int `json:"backer_count"`
	UserID int `json:"user_id"`
	Slug string `json:"slug"`
	Perks []string `json:"perks"`
	User CampaignUserFormatter `json:"user"`
	Images []CampaignImageFormatter `json:"images"`
}


func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormater := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormater := FormatCampaign(campaign)
		campaignsFormater = append(campaignsFormater, campaignFormater)
	}

	return campaignsFormater

}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.BackerCount = campaign.BackerCount
	campaignDetailFormatter.ImageUrl = ""
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	perks := []string{}

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}