import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { TreatmentRecordInterface } from "../models/ITreatmentRecord";
import { GetTreatmentRecord } from "../services/HttpClientService";
import moment from "moment";

function TreatmentRecord() {
    const [treatment, setTreatment] = useState<TreatmentRecordInterface[]>([])

    useEffect(() => {
        getTreatmentRecord();
    }, []);

    function getPatientFullName(params: { row: { DiagnosisRecord: { HistorySheet: { PatientRegister: { FirstName: any; LastName: any; }; }; }; }; }) {
        return `${params.row.DiagnosisRecord.HistorySheet.PatientRegister.FirstName || ''} ${params.row.DiagnosisRecord.HistorySheet.PatientRegister.LastName || ''}`
    }

    function getDoctorFullName(params: { row: { Doctor: { FirstName: any; LastName: any; }; }; }) {
        return `${params.row.Doctor.FirstName || ''} ${params.row.Doctor.LastName || ''}`
    }


    const getTreatmentRecord = async () => {
        let res = await GetTreatmentRecord();
        if (res) {
            setTreatment(res);
        }
    };

    const columns: GridColDef[] = [
        { field: "ID", headerName: "No.", width: 50 },
        {
            field: "Patient",
            headerName: "ผู้ป่วย",
            width: 200,
            valueGetter: getPatientFullName,
        },
        // {
        //   field: "User",
        //   headerName: "User",
        //   width: 150,
        //   valueGetter: (params) => params.row.Member.FirstName,
        // },
        // { field: "Amount", headerName: "Amount", width: 100 },
        // { field: "Employee",
        //   headerName: "Author",
        //   width: 150,
        //   valueGetter: (params) => params.row.Employee.FirstName,
        // },
        // { field: "BorrowTime", headerName: "Borrow Time", width: 200,
        //   valueGetter: (params) => moment(params.row.BorrowTime).format("DD/MM/YYYY HH:mm") },
    ];

    return (
        <div>
            <Container maxWidth="md">
                <Box display="flex" sx={{ marginTop: 2 }}>
                    <Typography
                     component="h2"
                     variant="h6"
                     color="primary"
                     gutterBottom>
                        รายการการรักษา
                    </Typography>
                </Box>
                <Box>
                    <Button
                     component={RouterLink}
                     to="/treatment_records/create"
                     variant="contained"
                     color="primary">
                        บันทึกการรักษา
                    </Button>
                </Box>
                <div style={{ height: 400 , width: "100%", marginTop: "20px"}}>
                    <DataGrid
                        rows={treatment}
                        getRowId={(row) => row.ID}
                        columns={columns}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                    />
                </div>
            </Container>
        </div>
    )
}

export default TreatmentRecord;