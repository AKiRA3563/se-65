export interface EmployeeInterface {
    ID:         number;
    FirstName: string;
    LastName:  string;
    Email:      string;
    Password:   string;
    Role:   RoleInterface;
    Gender: GenderInterface;
}

export interface RoleInterface {
    ID:     number;
    Name:   string;
}

export interface GenderInterface {
    ID:     number;
    Name:   string;
}

