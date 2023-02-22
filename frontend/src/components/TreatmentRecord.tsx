import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import SearchIcon from '@mui/icons-material/Search';
import EditIcon from '@mui/icons-material/Edit';
import { DataGrid, GridColDef, GridValueGetterParams } from "@mui/x-data-grid";
import { TreatmentRecordsInterface } from "../models/ITreatmentRecord";
import { GetTreatmentRecord } from "../services/HttpClientService";
import moment from "moment";

function TreatmentRecord() {
    const [treatmentRecord, setTreatmentRecord] = useState<TreatmentRecordsInterface[]>([])
    const [openedit, setOpenEdit] = useState(false);
    

    function getPatientFullName(params: GridValueGetterParams) {
        return `${params.row.DiagnosisRecord.HistorySheet.PatientRegister.FirstName || ''} ${params.row.DiagnosisRecord.HistorySheet.PatientRegister.LastName || ''}`
    }

    function getDoctorFullName(params: GridValueGetterParams) {
        return `${params.row.Doctor.Title.Name || ''}${params.row.Doctor.FirstName || ''} ${params.row.Doctor.LastName || ''}`
    }

    const handleEdit = () => {
        setOpenEdit(true);
    }

    const getTreatmentRecord = async () => {
        let res = await GetTreatmentRecord();
        if (res) {
            setTreatmentRecord(res);
        }
    };

    useEffect(() => {
        getTreatmentRecord();
    }, []);

    const columns: GridColDef[] = [
        { field: "ID", headerName: "No.", width: 50 },
        {   field: "Doctor",
            headerName: "แพทย์ผู้ตรวจ",
            width: 190,
            valueGetter: getDoctorFullName,
        },
        {
            field: "Patient",
            headerName: "ผู้ป่วย",
            width: 170,
            valueGetter: getPatientFullName,
        },
        {
            field: "Treatment",
            headerName: "การรักาษา",
            width: 200,
            valueGetter: (params) => params.row.Treatment,
        },
        {
            field: "ืNote",
            headerName: "หมายเหตุ",
            width: 200,
            valueGetter: (params) => params.row.Note,
        },
        {   field: "Medicine", 
            headerName: "รายการยา", 
            width: 100,
            renderCell: () => (
                <div>
                    &nbsp;
                  <Button 
                    onClick={handleEdit}
                    variant="contained" 
                    size="small" 
                    startIcon={<SearchIcon />}
                    color="success"> 
                    แสดง
                  </Button>
                </div>
              ),
        },
        {   field: "Appointment", headerName: "การนัดหมาย", width: 120,
            renderCell : params => {
                if (params.row.Appointment === true) {
                    return <div>ใช่</div>
                } return <div>ไม่</div>
            },
        },
        
        {   field: "Date", headerName: "Date", width: 140,
            valueGetter: (params) => moment(params.row.Date).format("DD/MM/YYYY") 
        },
    ];


    return (
        <div>
            <Container maxWidth="lg">
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
                        rows={treatmentRecord}
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