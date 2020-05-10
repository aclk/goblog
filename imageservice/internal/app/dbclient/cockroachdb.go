package dbclient

import (
	"fmt"
	"github.com/aclk/goblog/imageservice/cmd"
	"strconv"

	"context"
	"github.com/aclk/goblog/common/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
)

type IGormClient interface {
	UpdateAccountImage(ctx context.Context, accountImage model.AccountImage) (model.AccountImage, error)
	StoreAccountImage(ctx context.Context, accountImage model.AccountImage) (model.AccountImage, error)
	QueryAccountImage(ctx context.Context, accountId string) (model.AccountImage, error)
	SetupDB(addr string)
	SeedAccountImages() error
	Check() bool
	Close()
}

func (gc *GormClient) StoreAccountImage(ctx context.Context, accountImage model.AccountImage) (model.AccountImage, error) {
	tx := gc.crDB.Begin()
	accountImage.ID = uuid.NewV4().String()
	tx = tx.Create(&accountImage)
	if tx.Error != nil {
		logrus.Errorf("Error storing account image: %v", tx.Error.Error())
		return model.AccountImage{}, tx.Error
	}
	tx = tx.Commit()
	return accountImage, nil
}

func (gc *GormClient) UpdateAccountImage(ctx context.Context, accountImage model.AccountImage) (model.AccountImage, error) {
	tx := gc.crDB.Begin()
	tx = tx.Save(&accountImage)
	if tx.Error != nil {
		logrus.Errorf("Error updating account image: %v", tx.Error.Error())
		return model.AccountImage{}, tx.Error
	}
	tx = tx.Commit()
	accountImage, _ = gc.QueryAccountImage(ctx, accountImage.ID)
	return accountImage, nil
}

type GormClient struct {
	crDB *gorm.DB
}

func NewGormClient(cfg *cmd.Config) *GormClient {
	gc := &GormClient{}
	gc.SetupDB(cfg.CockroachdbConnUrl)
	return gc
}

func (gc *GormClient) Check() bool {
	return gc.crDB != nil
}

func (gc *GormClient) Close() {
	logrus.Infoln("Closing connection to CockroachDB")
	gc.crDB.Close()
}

func (gc *GormClient) QueryAccountImage(ctx context.Context, accountId string) (model.AccountImage, error) {

	if gc.crDB == nil {
		return model.AccountImage{}, fmt.Errorf("Connection to DB not established!")
	}
	acc := model.AccountImage{}
	gc.crDB.First(&acc, "ID = ?", accountId)
	if acc.ID == "" {
		return acc, fmt.Errorf("")
	}
	return acc, nil
}

func (gc *GormClient) SetupDB(addr string) {
	logrus.Infof("Connecting with connection string: '%v'", addr)
	var err error
	gc.crDB, err = gorm.Open("postgres", addr)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	gc.crDB.AutoMigrate(&model.AccountImage{})

	logrus.Info("Successfully connected to DB and executed auto-migration")
}

func (gc *GormClient) SeedAccountImages() error {
	if gc.crDB == nil {
		return fmt.Errorf("Connection to DB not established!")
	}
	gc.crDB.Delete(&model.AccountImage{})

	total := 100
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.AccountImage{
			ID:  key,
			URL: "http://path.to.some.image/" + key + ".png",
		}

		gc.crDB.Create(&acc)
		logrus.Infoln("Successfully created AccountImage instance.")
	}
	return nil
}
