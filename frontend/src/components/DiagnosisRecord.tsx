import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { DiagnosisRecordsInterface } from "../models/IDiagnosisRecord";
import { GetDiagnosisRecord } from "../services/HttpClientService";
import moment from "moment";

function DiagnosisRecord() {
    const [diagnosis, setDiagnosis] = useState<DiagnosisRecordsInterface[]>([])

    useEffect(() => {
        getDiagnosisRecord();
    }, []);

    function getPatientFullName(params: { row: { HistorySheet: { PatientRegister: { FirstName: any; LastName: any; }; }; }; }) {
        return `${params.row.HistorySheet.PatientRegister.FirstName || ''} ${params.row.HistorySheet.PatientRegister.LastName || ''}`
    }

    function getDoctorFullName(params: { row: { Doctor: { FirstName: any; LastName: any; }; }; }) {
        return `${params.row.Doctor.FirstName || ''} ${params.row.Doctor.LastName || ''}`
    }

    const getDiagnosisRecord = async () => {
        let res = await GetDiagnosisRecord();
        if (res) {
            setDiagnosis(res);
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
        {
            field: "Examination",
            headerName: "การวินิจฉัย",
            width: 150,
            valueGetter: (params) => params.row.Examination,
        },
        {
            field: "Disease",
            headerName: "วินิจฉัยโรค",
            width: 150,
            valueGetter: (params) => params.row.Disease.Name,
        },
        {   field: "MedicalCertificate", headerName: "ใบรับรองแพทย์", width: 100,
            renderCell : params => {
                if (params.row.MedicalCertificate == true) {
                    return <div>รับ</div>
                } return <div>ไม่รับ</div>
            },
        },
        {   field: "Doctor",
            headerName: "ผู้ตรวจ",
            width: 150,
            valueGetter: getDoctorFullName,
        },
        {   field: "Date", headerName: "Date", width: 200,
            valueGetter: (params) => moment(params.row.Date).format("DD/MM/YYYY") },
    ];

    return (
        <div>
            <Container maxWidth="md">
                <Box display="flex" sx={{ marginTop: 2 }}>
                    <Typography
                     component="h2"
                     variant="h6"
                     color="primary"
                     gutterBottom
                    >
                        รายการผลวินิจฉัย
                    </Typography>
                </Box>
                <Box>
                    <Button
                     component={RouterLink}
                     to="/diagnosis_records/create"
                     variant="contained"
                     color="primary">
                        บันทึกผลการวินิจฉัย
                    </Button>
                </Box>
                <div style={{ height: 400 , width: "100%", marginTop: "20px"}}>
                    <DataGrid
                        rows={diagnosis}
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

export default DiagnosisRecord;