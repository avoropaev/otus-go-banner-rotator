package apitest

import (
	"github.com/avoropaev/otus-go-banner-rotator/internal/app"
	"github.com/avoropaev/otus-go-banner-rotator/internal/server/pb"
)

func (s *APISuite) TestGetBannerSuccess() {
	banner, err := s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotGuid5,
		SocialGroupGuid: SocialGroupGuid1,
	})
	s.Require().NoError(err)
	s.Require().NotNil(banner)
}

func (s *APISuite) TestGetBannerErrors() {
	_, err := s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotGuidNotFound,
		SocialGroupGuid: SocialGroupGuid1,
	})
	s.Require().ErrorContains(err, app.ErrSlotNotFound.Error())

	_, err = s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotGuid1,
		SocialGroupGuid: SocialGroupGuidNotFound,
	})
	s.Require().ErrorContains(err, app.ErrSocialGroupNotFound.Error())

	_, err = s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotWithoutLinks,
		SocialGroupGuid: SocialGroupGuid1,
	})
	s.Require().ErrorContains(err, app.ErrNoOneBannerFoundForSlot.Error())
}
