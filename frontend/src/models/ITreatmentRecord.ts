import { DiagnosisRecordInterface } from "./IDiagnosisRecord";

export interface TreatmentRecordInterface {
    No?: number;

    PatientID?: number;
	//Patient?: PatientRegisterInterface;

	DoctorID?: number;
	//Doctor?: EmployeeInterface;

    DiagnosisRecordID?: number;
    DiagnosisRecord?:   DiagnosisRecordInterface;

    MedicineID?: number;
    Medicine?:   MedicineInterface;
    
    MedicineQuantity?: number;
    Treatment?: string;
    Note?: string;
    Appointment?: boolean;//int;
    Date?: Date | null;
}

export interface MedicineInterface {
    ID?: number;

    Name?: string;
    Description?: string;
    Quantity?: number;
}