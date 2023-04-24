package main

import (
	"bwastartup/campaign"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestKoneksi(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func TestCampaignFindAll(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	campaignRepository := campaign.NewRepository(db)
	campaigns, err := campaignRepository.FindAll()

	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println(len(campaigns))

	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
	}
}

func TestCampaignFindByuserId(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	campaignRepository := campaign.NewRepository(db)
	campaigns, err := campaignRepository.FindByUserID(8)

	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println(len(campaigns))

	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
	}
}
