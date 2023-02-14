import { EmployeeInterface } from "./IEmployee";
import { PatientRegisterInterface } from "./IPatientRegister";

export interface HistorySheetInterface {
    ID: number;
    Weight: Float32Array;
    Height: Float32Array;
    BMI:    Float32Array;
    Temperature:    Float32Array;
    HeartRate:      number;
    OxygenSaturation:   number;
    
    PatientRegisterID:  number
    PatientRegister:    PatientRegisterInterface;
    
    Nurse:      EmployeeInterface;
}