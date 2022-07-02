package apitest

import (
	"github.com/avoropaev/otus-go-banner-rotator/internal/app"
	"github.com/avoropaev/otus-go-banner-rotator/internal/server/pb"
)

func (s *APISuite) TestGetBannerSuccess() {
	countMessage := s.GetCountMessageInAMQP()

	banner, err := s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotGUID5,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().NoError(err)
	s.Require().NotNil(banner)
	s.Require().Equal(countMessage+1, s.GetCountMessageInAMQP())
}

func (s *APISuite) TestGetBannerErrors() {
	countMessage := s.GetCountMessageInAMQP()

	_, err := s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotGUIDNotFound,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().ErrorContains(err, app.ErrSlotNotFound.Error())
	s.Require().Equal(countMessage, s.GetCountMessageInAMQP())

	_, err = s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotGUID1,
		SocialGroupGuid: SocialGroupGUIDNotFound,
	})
	s.Require().ErrorContains(err, app.ErrSocialGroupNotFound.Error())
	s.Require().Equal(countMessage, s.GetCountMessageInAMQP())

	_, err = s.client.GetBanner(s.ctx, &pb.SlotAndSocialGroupRequest{
		SlotGuid:        SlotWithoutLinks,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().ErrorContains(err, app.ErrNoOneBannerFoundForSlot.Error())
	s.Require().Equal(countMessage, s.GetCountMessageInAMQP())
}
