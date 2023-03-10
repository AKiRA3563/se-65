import { EmployeesInterface } from "./IEmployee";
import { HistorySheetsInterface } from "./IHistorySheet";
// import {EmployeesInterface} from "../IEmployee/IEmployee"
// import { HistorySheetsInterface } from "../HistorySheet/IHistorySheet";
//import { PatientRegistersInterface } from "../IPatientRegister/IPatientRegister";

export interface DiagnosisRecordsInterface {
	ID?: number;

	DoctorID?: number;
	Doctor?: EmployeesInterface;

	HistorySheetID?: number;
	HistorySheet?: HistorySheetsInterface;

	DiseaseID?: number;
	Disease?: DiseasesInterface;

	Examination?: string;
	MedicalCertificate?: boolean | null;
	Date?: Date | null;
}

export interface DiseasesInterface {
	ID: number;
	Name: string;
}