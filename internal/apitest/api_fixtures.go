package apitest

const (
	SlotGuid1        = "00000000-0000-0000-0000-000000000001"
	SlotGuid5        = "00000000-0000-0000-0000-000000000005"
	SlotWithoutLinks = "00000000-0000-0000-0000-000000000006"
	SlotGuidNotFound = "00000000-0000-0000-0000-000000000009"

	BannerGuid1        = "00000000-0000-0000-1111-000000000001"
	BannerGuid5        = "00000000-0000-0000-1111-000000000005"
	BannerGuidNotFound = "00000000-0000-0000-1111-000000000009"

	SocialGroupGuid1        = "00000000-0000-0000-2222-000000000001"
	SocialGroupGuidNotFound = "00000000-0000-0000-2222-000000000009"
)

type FixtureLink struct {
	bannerGUID string
	slotGUID   string
}

var (
	Link1 = FixtureLink{bannerGUID: BannerGuid1, slotGUID: SlotGuid1}

	LinkNotFound1 = FixtureLink{bannerGUID: BannerGuid5, slotGUID: SlotGuid1}
)
