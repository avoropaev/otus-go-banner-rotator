package apitest

import (
	"github.com/avoropaev/otus-go-banner-rotator/internal/app"
	"github.com/avoropaev/otus-go-banner-rotator/internal/server/pb"
)

func (s *APISuite) TestAddClickSuccess() {
	_, err := s.client.AddClick(s.ctx, &pb.AddClickRequest{
		BannerGuid:      BannerGUID1,
		SlotGuid:        SlotGUID1,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().NoError(err)

	_, err = s.client.AddClick(s.ctx, &pb.AddClickRequest{
		BannerGuid:      BannerGUID1,
		SlotGuid:        SlotGUID1,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().NoError(err)
}

func (s *APISuite) TestAddClickErrors() {
	_, err := s.client.AddClick(s.ctx, &pb.AddClickRequest{
		BannerGuid:      BannerGUIDNotFound,
		SlotGuid:        SlotGUID1,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().ErrorContains(err, app.ErrBannerNotFound.Error())

	_, err = s.client.AddClick(s.ctx, &pb.AddClickRequest{
		BannerGuid:      BannerGUID1,
		SlotGuid:        SlotGUIDNotFound,
		SocialGroupGuid: SocialGroupGUID1,
	})
	s.Require().ErrorContains(err, app.ErrSlotNotFound.Error())

	_, err = s.client.AddClick(s.ctx, &pb.AddClickRequest{
		BannerGuid:      BannerGUID1,
		SlotGuid:        SlotGUID1,
		SocialGroupGuid: SlotGUIDNotFound,
	})
	s.Require().ErrorContains(err, app.ErrSocialGroupNotFound.Error())
}
