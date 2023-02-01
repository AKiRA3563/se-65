import React from "react";
import { DiagnosisRecordInterface } from "../models/IDiagnosisRecord";
import { TreatmentRecordInterface } from "../models/ITreatmentRecord";

const apiUrl = "http://localhost:808";

//GET DiagnosisRecord
async function GetDiagnosisRecords() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Conten-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/diagnosis_records`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

//GET Disease
async function GetDisease() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/diseases`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
        if (res.data) {
            return res.data;
        } else {
            return false;
        }
    });

    return res;
}

//GET Treatment
async function GetTreatmentRecords() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/treatment_records`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
        if (res.data) {
            return res.data;
        } else {
            return false;
        }
    });

    return res;
}


//GET Medicine
async function GetMedicine() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/medicines`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
        if (res.data) {
            return res.data;
        } else {
            return false;
        }
    });

    return res;
}

// Get Employee
async function GetEmployee() {
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
  
    let res = await fetch(`${apiUrl}/employee`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          return res.data;
        } else {
          return false;
        }
      });
  
    return res;
}

//================================================================================================

//Create Diagnosis
async function CreateDiagnosisRecord(data: DiagnosisRecordInterface) {
    const reqquestOption = {
        menthod: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/diagnosis_records`, reqquestOption)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

//Create Treatment
async function CreatetreatmentRecord(data: TreatmentRecordInterface) {
    const reqquestOption = {
        menthod: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/treatment_records`, reqquestOption)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

export {
    GetDiagnosisRecords,
    GetDisease,
    GetTreatmentRecords,
    GetMedicine,
    GetEmployee,

    CreateDiagnosisRecord,
    CreatetreatmentRecord,
};