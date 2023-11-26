package response

import "github.com/arvinpaundra/ngekost-api/internal/entity"

type (
	RoomAsset struct {
		ID     string `json:"id"`
		RoomId string `json:"room_id"`
		Url    string `json:"url"`
		Type   string `json:"type"`
	}
)

func ToResponseRoomAssets(assets []*entity.RoomAsset) []*RoomAsset {
	var res []*RoomAsset

	for _, asset := range assets {
		res = append(res, &RoomAsset{
			ID:     asset.ID,
			RoomId: asset.RoomId,
			Url:    asset.Url,
			Type:   asset.Type,
		})
	}

	return res
}
