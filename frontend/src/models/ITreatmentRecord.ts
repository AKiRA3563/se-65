import { DiagnosisRecordsInterface } from "./IDiagnosisRecord";
import { EmployeesInterface } from "./IEmployee";
// import { PatientRegistersInterface } from "./IPatientRegister";

export interface TreatmentRecordInterface {
    ID?: number;

    // PatientRegisterID?: number;
    // PatientRegister?:   PatientRegisterInterface;

	DoctorID?: number;
	Doctor?: EmployeesInterface;

    DiagnosisRecordID?: number;
    DiagnosisRecord?:   DiagnosisRecordsInterface;

    MedicineID?: number;
    Medicine?:   MedicineInterface;
    
    MedicineQuantity?: number;
    Treatment?: string;
    Note?: string;
    
    Appointment?: boolean | null;
    Date?: Date | null;
}

export interface MedicineInterface {
    ID: number;

    Name: string;
    Description: string;
    Price: number;
}