package response

import "github.com/arvinpaundra/ngekost-api/internal/entity"

type (
	Room struct {
		ID          string  `json:"id"`
		KostId      string  `json:"kost_id"`
		Name        string  `json:"name"`
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
		Description *string `json:"description"`
		Category    *string `json:"category"`
		Image       *string `json:"image"`
	}

	RoomDetail struct {
		Room
		Assets []*RoomAsset `json:"assets"`
	}
)

func ToResponseRooms(rooms []*entity.Room) []*Room {
	var res []*Room

	for _, room := range rooms {
		res = append(res, &Room{
			ID:          room.ID,
			KostId:      room.KostId,
			Name:        room.Name,
			Quantity:    room.Quantity,
			Price:       room.Price,
			Description: room.Description,
			Category:    room.Category,
			Image:       room.Image,
		})
	}

	return res
}

func ToResponseRoom(room *entity.Room, assets []*entity.RoomAsset) *RoomDetail {
	return &RoomDetail{
		Room: Room{
			ID:          room.ID,
			KostId:      room.KostId,
			Name:        room.Name,
			Quantity:    room.Quantity,
			Price:       room.Price,
			Description: room.Description,
			Category:    room.Category,
			Image:       room.Image,
		},
		Assets: ToResponseRoomAssets(assets),
	}
}
