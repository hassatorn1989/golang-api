package repositories

import (
	"todo-fiber/internal/dto"
	"todo-fiber/internal/models"

	"gorm.io/gorm"
)

// UserRepository เป็น struct ที่ใช้สำหรับจัดการกับข้อมูลผู้ใช้ในฐานข้อมูล โดยใช้ GORM เป็น ORM (Object-Relational Mapping) เพื่อทำงานกับฐานข้อมูลได้ง่ายขึ้น
type UserRepository struct {
	db *gorm.DB
}

// สร้าง instance ของ UserRepository โดยรับ gorm.DB เป็นพารามิเตอร์
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindAll เป็น method ที่ใช้สำหรับดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล โดยจะคืนค่ากลับเป็น slice ของ UserModel และ error ถ้ามีข้อผิดพลาดเกิดขึ้น
func (ur *UserRepository) FindAll() ([]models.UserModel, error) {
	var users []models.UserModel
	q := ur.db.Find(&users)
	if err := q.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) FindWithPagination(limit int, offset int, search string) ([]models.UserModel, error) {
	var users []models.UserModel
	q := ur.db.Where(
		"name LIKE ? OR email LIKE ?",
		"%"+search+"%", "%"+search+"%",
	)
	q = q.Limit(limit).Offset(offset).Order("created_at DESC").Find(&users)
	if err := q.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.UserModel, error) {
	user := &models.UserModel{}
	q := ur.db.Where("email = ?", email).First(user)
	if err := q.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindByID(id string) (*models.UserModel, error) {
	user := &models.UserModel{}
	q := ur.db.Where("id = ?", id).First(user)
	if err := q.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Create(req *dto.CreateUserRequest) error {
	user := &models.UserModel{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password, // In a real application, ensure this is a securely hashed password
		Role:     req.Role,
	}
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Update(id string, req *dto.UpdateUserRequest) error {
	user := &models.UserModel{}
	q := ur.db.Where("id = ?", id).First(user)

	if err := q.Error; err != nil {
		return err
	}
	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password
	user.Role = req.Role

	if err := ur.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id string) error {
	q := ur.db.Where("id = ?", id).Delete(&models.UserModel{})
	if err := q.Error; err != nil {
		return err
	}
	return nil
}
