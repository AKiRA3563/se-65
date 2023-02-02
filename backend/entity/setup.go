package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

)

var db *gorm.DB

func DB() *gorm.DB {
	
	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("se-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&DiagnosisRecord{},
		&TreatmentRecord{},
		&HistorySheet{},
		&Disease{},
		&Medicine{},
		&Employee{},
		&PatientRegister{},
		&TreatmentRecord{},
	)

	db = database

	// Disease Data
	Disease1 := Disease{
		Name: "ไม่มี",
	}
	db.Model(&Disease{}).Create(Disease1)

	// db.Model(&Disease{}).Create([]map[string]interface{}{
	// 	{"Name": "ไข้หวัดใหญ่"},
	// 	{"Name": "ไข้เลือดออก"},
	// 	{"Name": "ไข้ หรือ ไข้ไม่ทราบสาเหตุ"},
	// 	{"Name": "ติดเชื้อระบบทางเดินอาหารจากแบคทีเรียชนิดเฉียบพลัน"},
	// 	{"Name": "ทางเดินอาหารอักเสบเฉียบพลันจากไวรัส"},
	// 	{"Name": "พิษสุนัขบ้า"},
	// 	{"Name": "พิษจากแก๊ส สารไอระเหย"},
	// 	{"Name": "มาลาเรีย"},
	// 	{"Name": "โรคตาแดง โรคตาอักเสบ"},

	// 	{"Name": "โรคเบาหวาน"},
	// 	{"Name": "โรคเพศสัมพันธุ์อื่น ๆ"},
	// 	{"Name": "โรคภูมิแพ้"},
	// 	{"Name": "อาหารเป็นพิษ"},
	// 	{"Name": "เอดส์"},
	// })


	// Medicine Data
	Medicine1 := Medicine{
		Name: "ไม่มี", 
		Description: "none",
	}
	db.Model(&Medicine{}).Create(Medicine1)

	// db.Model(&Medicine{}).Create([]map[string]interface{}{
	// 	{"Name": "Paracetamol", "Description": "ใช้บรรเทาปวด ลดไข้",},
	// 	{"Name": "Antacil Gel", "Description": "บรรเทาอาการปวดท้อง ท้องอืด จุกเสียด แน่น",},
	// 	// {"Name": "", "Description": "",},
	// 	// {"Name": "", "Description": "",},
	// 	// {"Name": "", "Description": "",},
	// })


	password, err := bcrypt.GenerateFromPassword([]byte("qwerty"), 14)

	Employee1 := Employee{
		FirstName:  "Namjai",
		LastName:   "Meedee",
		Email:      "namn@gmail.com",
		Password:   string(password),
		BirthDay:   time.Now(),
		//Gender:     "Female",
		IDCard:   "1234567890098",
		Phonenumber:     "0812345678",
		
	}
	db.Model(&Employee{}).Create(&Employee1)

	Patient1 := PatientRegister{
		FirstName:  "Namjai",
		LastName:   "Meedee",
		Age: 		24,
		Mobile:     "09234569872",
		Email:      "namn@gmail.com",

	}
	db.Model(&PatientRegister{}).Create(&Patient1)
}
