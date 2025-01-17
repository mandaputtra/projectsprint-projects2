package repositories

import (
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/models"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{
		db: db,
	}
}

func (r *ActivityRepository) Create(activity *models.Activity) (*models.Activity, error) {
	if err := r.db.Create(activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *ActivityRepository) GetAll(limit, offset int, userId string) ([]*models.Activity, error) {
	var activities []*models.Activity
	query := r.db.Limit(limit).Offset(offset)

	err := query.Find(&activities, "user_id = ?", userId).Error
	return activities, err
}

// func (r *ActivityRepository) GetAllWithoutPaginationById(email string) ([]*models.Activity, error) {
// 	var activitys []*models.Activity

// 	// Lakukan query untuk mencari data dengan ID yang sesuai
// 	err := r.db.Where("id ILIKE ?", "%"+email+"%").Find(&activitys).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return activitys, nil
// }

func (r *ActivityRepository) GetOne(id, userId string) (*models.Activity, error) {
	var activity models.Activity
	err := r.db.First(&activity, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

// func (r *ActivityRepository) UpdateActivity(data *models.Activity) (*models.Activity, error) {

// 	updateErr := r.db.Model(&models.Activity{}).
// 		Where("id = ?", data.ID).
// 		Updates(data).
// 		First(data).Error

// 	if updateErr != nil {
// 		return nil, updateErr
// 	}
// 	return data, nil
// }

// func (r *ActivityRepository) UpdateActivityId(id string, data *models.Activity) (*models.Activity, error) {
// 	updateErr := r.db.Model(&models.Activity{}).Where("id = ?", id).Updates(data).First(data).Error

// 	if updateErr != nil {
// 		return nil, updateErr
// 	}
// 	return data, nil
// }

// func (r *ActivityRepository) DeleteById(id string) error {

// 	activity, err := r.GetOne(id)
// 	if err != nil {
// 		return err
// 	}

// 	deleteErr := r.db.Delete(activity).Error
// 	return deleteErr
// }
