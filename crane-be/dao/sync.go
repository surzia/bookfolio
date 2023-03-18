package dao

import (
	"papercrane/models"

	"gorm.io/gorm"
)

type SyncDao struct {
	db *gorm.DB
}

func NewSyncDao(db *gorm.DB) *SyncDao {
	syncDao := &SyncDao{db: db}

	return syncDao
}

func (s *SyncDao) CreateSync(req *models.SaveSyncReq) {
	ret := s.db.Create(req.ToSync())
	if ret.Error != nil {
		panic(ret.Error)
	}
}

func (s *SyncDao) FetchUsernameAndPasswordByType(server string) (string, string) {
	var sync models.Sync
	if err := s.db.Where("type = ?", server).First(&sync).Error; err != nil {
		return "", ""
	}
	return sync.Username, sync.Password
}

func (s *SyncDao) CheckStatus(key string) bool {
	var sync models.Sync
	result := s.db.Find(&sync, models.Sync{Type: key})
	return result.RowsAffected != 0
}
