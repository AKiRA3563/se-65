import { EmployeeInterface } from "./IEmployee";

export interface DiagnosisRecordInterface {
	No?: number;

	PatientID?: number;
	//Patient?: PatientRegisterInterface;

	DoctorID?: number;
	Doctor?: EmployeeInterface;

	HistorySheetID?: number;
	//HistorySheet?: HistorySheetInterface;

	DiseaseID?: number;
	Disease?: DiseaseInterface;
	
	Examination?: string;
	Medicalcertification?: boolean;//int;
	Date?: Date | null;
}

export interface DiseaseInterface {
	ID?: number;

	Name?: string;
	Description?: string;
}