import { GenderInterface } from "./IEmployee";

export interface PatientRegisterInterface {
    ID:         number;
    FirstName:  string;
    LastName:   string;
    Gender:     GenderInterface;
    Age:        number;
    Mobile:     string;
    Email:      string;

}
