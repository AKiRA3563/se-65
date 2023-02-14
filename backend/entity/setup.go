package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		&Disease{},

		&HistorySheet{},
		&PatientRegister{},

		&TreatmentRecord{},
		&Medicine{},

		// &MedicalCertificate{},

		&Employee{},
		&Gender{},
		&Role{},
	)
	db = database

	//================Disease Data======================
	disease1 := Disease{
		Name: "ไม่มี",
	}
	db.Model(&Disease{}).Create(&disease1)

	disease2 := Disease{
		Name: "ไข้หวัดใหญ่",
	}
	db.Model(&Disease{}).Create(&disease2)

	disease3 := Disease{
		Name: "ไข้เลือดออก",
	}
	db.Model(&Disease{}).Create(&disease3)

	disease4 := Disease{
		Name: "ไข้หวัดใหญ่",
	}
	db.Model(&Disease{}).Create(&disease4)

	disease5 := Disease{
		Name: "ไข้ หรือ ไข้ไม่ทราบสาเหตุ",
	}
	db.Model(&Disease{}).Create(&disease5)

	disease6 := Disease{
		Name: "ติดเชื้อระบบทางเดินอาหารจากแบคทีเรียชนิดเฉียบพลัน",
	}
	db.Model(&Disease{}).Create(&disease6)

	disease7 := Disease{
		Name: "ทางเดินอาหารอักเสบเฉียบพลันจากไวรัส",
	}
	db.Model(&Disease{}).Create(&disease7)

	disease8 := Disease{
		Name: "พิษสุนัขบ้า",
	}
	db.Model(&Disease{}).Create(&disease8)

	disease9 := Disease{
		Name: "พิษจากแก๊ส สารไอระเหย",
	}
	db.Model(&Disease{}).Create(&disease9)

	disease10 := Disease{
		Name: "โรคตาแดง โรคตาอักเสบ",
	}
	db.Model(&Disease{}).Create(&disease10)

	disease11 := Disease{
		Name: "โรคเบาหวาน",
	}
	db.Model(&Disease{}).Create(&disease11)

	disease12 := Disease{
		Name: "โรคเพศสัมพันธุ์อื่น ๆ",
	}
	db.Model(&Disease{}).Create(&disease12)

	disease13 := Disease{
		Name: "โรคภูมิแพ้",
	}
	db.Model(&Disease{}).Create(&disease13)

	disease14 := Disease{
		Name: "อาหารเป็นพิษ",
	}
	db.Model(&Disease{}).Create(&disease14)

	//=================Medicine Data====================
	medicine1 := Medicine{
		Name:        "ไม่มี",
		Description: "none",
		Price:       0,
	}
	db.Model(&Medicine{}).Create(&medicine1)

	medicine2 := Medicine{
		Name:        "Paracetamol",
		Description: "ใช้บรรเทาปวด ลดไข้",
		Price:       15,
	}
	db.Model(&Medicine{}).Create(&medicine2)

	medicine3 := Medicine{
		Name:        "Antacil Gel",
		Description: "บรรเทาอาการปวดท้อง ท้องอืด จุกเสียด แน่น",
		Price:       42,
	}
	db.Model(&Medicine{}).Create(&medicine3)

	//=================Employee========================
	password, err := bcrypt.GenerateFromPassword([]byte("qwerty"), 14)

	gender1 := Gender{
		Name: "Male",
	}
	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{
		Name: "Female",
	}
	db.Model(&Gender{}).Create(&gender2)

	role1 := Role{
		Name: "Doctor",
	}
	db.Model(&Role{}).Create(&role1)

	role2 := Role{
		Name: "Nurse",
	}
	db.Model(&Role{}).Create(&role2)

	employee1 := Employee{
		FirstName:   "Namjai",
		LastName:    "Meedee",
		Email:       "namn@gmail.com",
		Password:    string(password),
		Birthday:    time.Now(),
		Gender:      gender1,
		Role:        role2,
		IDCard:      "1234567890098",
		Phonenumber: "0812345678",
	}
	db.Model(&Employee{}).Create(&employee1)

	employee2 := Employee{
		FirstName:   "Meedee",
		LastName:    "winai",
		Email:       "meem@gmail.com",
		Password:    string(password),
		Birthday:    time.Date(2000, time.August, 1, 0, 0, 0, 0, time.UTC),
		Gender:      gender1,
		Role:        role1,
		IDCard:      "6222023559638",
		Phonenumber: "0912345678",
	}
	db.Model(&Employee{}).Create(&employee2)

	//================Patient======================
	patient1 := PatientRegister{
		FirstName: "Namjai",
		LastName:  "Meedee",
		Gender:    gender2,
		Age:       24,
		Mobile:    "09234569872",
		Email:     "Maidee@gmail.com",
	}
	db.Model(&PatientRegister{}).Create(&patient1)

	patient2 := PatientRegister{
		FirstName: "Noijai",
		LastName:  "Maidee",
		Gender:    gender1,
		Age:       36,
		Mobile:    "06234569872",
		Email:     "noijai@gmail.com",
	}
	db.Model(&PatientRegister{}).Create(&patient2)

	//=============History Sheet=====================
	historysheet1 := HistorySheet{
		Weight:           60.2,
		Height:           173,
		Temperature:      37.6,
		HeartRate:        88,
		OxygenSaturation: 98,
		PatientRegister:  patient2,
		Nurse:            employee1,
	}
	db.Model(&HistorySheet{}).Create(&historysheet1)

	historysheet2 := HistorySheet{
		Weight:           53.6,
		Height:           163,
		Temperature:      38.6,
		HeartRate:        86,
		OxygenSaturation: 98,
		PatientRegister:  patient1,
		Nurse:            employee1,
	}
	db.Model(&HistorySheet{}).Create(&historysheet2)

	//==============Medicalcertificate=================
	// med1 := MedicalCertificate{
	// 	Label: "รับ",
	// }
	// db.Model(&MedicalCertificate{}).Create(&med1)

	// med2 := MedicalCertificate{
	// 	Label: "ไม่รับ",
	// }
	// db.Model(&MedicalCertificate{}).Create(&med2)

	//==============Diagnosis Record====================
	t := true
	// f := false
	diagnosis1 := DiagnosisRecord{
		// PatientRegister:	historysheet1.PatientRegister,
		Doctor:             employee2,
		HistorySheet:       historysheet1,
		Disease:            disease10,
		Examination:        "ผู้ป่วยมีอาการ",
		MedicalCertificate: &t,
		Date:               time.Now(),
	}
	db.Model(&DiagnosisRecord{}).Create(&diagnosis1)

	treatmentrecord1 := TreatmentRecord{
		// PatientRegister: 	diagnosis1.PatientRegister,
		Doctor:           employee2,
		DiagnosisRecord:  diagnosis1,
		Treatment:        "ให้ทานยา และพักผ่อน",
		Note:             "รอดูอาการในการตรวจครั้งหน้า",
		Appointment:      &t,
		Medicine:         medicine2,
		MedicineQuantity: 24,
		 Date:             time.Now(),
	}
	db.Model(&TreatmentRecord{}).Create(&treatmentrecord1)
}
