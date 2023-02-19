import { GendersInterface } from "./IEmployee";

export interface PatientRegistersInterface {
    ID: string,
    FirstName: string;
    LastName: string;
    IdentificationNumber:   string;
    Age:    number;
    BirthDay: Date | null;
    Mobile: string;
    Email: string;
    Occupation: string;
    Address:    string;
    EmergencyPersonFirstName:           string;
	EmergencyPersonLastName :           string;
	EmergencyPersonMobile :             string;
	EmergencyPersonOccupation :         string;
	EmergencyPersonRelationWithPatient: string;

    Gender: GendersInterface;
    Prefix: string;
    Nationality: string;
    Religion: string;
    BloodType: string;
    MaritalStatus: string;
    SubDistrict: string;
    District: string;
    Province: string;
}
