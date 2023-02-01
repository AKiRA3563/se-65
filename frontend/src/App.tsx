import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { FormControl, MenuItem, Snackbar } from "@mui/material";
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import Paper from "@mui/material/Paper";
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import MuiAlert, { AlertProps } from "@mui/material/Alert";

//import '../App.css';
import { DiagnosisRecordInterface } from "./models/IDiagnosisRecord";
import { PartientRegisterInterface } from "./models/IPatientRegister";
import { EmployeeInterface } from "./models/IEmployee"; 

import {
  GetEmployee,
  GetEquipment,
  GetUser,
  BorrowLists,
  GetEmployeeLogIn,
} from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function BorrowListCreate() {
  const [employee, setEmployee] = useState<EmployeeInterface[]>([]);
  const [employeeLogIn, setEmployeeLogIn] = useState<EmployeeInterface>();
  const [user, setUser] = useState<UsersInterface[]>([]);
  const [equipment, setEquipment] = useState<EquipmentInterface[]>([]);
  const [borrowList, setBorrowList] = useState<Partial<BorrowListInterface>>({
    Amount: 0,
    BorrowTime: new Date(),
    EmployeeID: employeeLogIn?.ID,
  });

  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };
 
  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof borrowList;
    setBorrowList({
      ...borrowList,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof BorrowListCreate;
    const { value } = event.target;
    setBorrowList({ ...borrowList, [id]: value });
  };

  const getEmployee = async () => {
    let res = await GetEmployee();
    if (res) {
      setEmployee(res);
    }
  };

  const getEmployeeLogIn = async () => {
    let res = await GetEmployeeLogIn();
    if (res) {
      setEmployeeLogIn(res);
    }
  };

  const getEquipment = async () => {
    let res = await GetEquipment();
    if (res) {
      setEquipment(res);
    }
  };

  const getUser = async () => {
    let res = await GetUser();
    if (res) {
      setUser(res);
    }
  };

  useEffect(() => {
    getEmployeeLogIn();
    getUser();
    getEmployee();
    getEquipment();
    setBorrowList({
      ...borrowList,
      ["EmployeeID"]: employeeLogIn?.ID,
    });
  }, [employeeLogIn]);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      EmployeeID: convertType(borrowList?.EmployeeID),
      MemberID: convertType(borrowList?.MemberID),
      EquipmentID: convertType(borrowList?.EquipmentID),
      Amount: convertType(borrowList?.Amount),
      BorrowTime: borrowList?.BorrowTime,
    };

    let res = await BorrowLists(data);
    if (res) {
      setSuccess(true);
      window.location.href="/borrow_list";
    } else {
      setError(true);
    }
  };

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>

      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          เกิดความผิดพลาด
        </Alert>
      </Snackbar>

      <Container maxWidth="md">
        <Paper>
          <Box
            display={"flex"}
            sx={{
              marginTop: 2,
              paddingX: 2,
              paddingY: 2,
            }}
          >
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกรายการยืมอุปกรณ์ 
            </Typography>
          </Box>
          <hr/>
          <Grid container spacing={2} columns={8} className="centered-div">
        
            <Grid item xs={3}>
              <FormControl fullWidth>
                <p>Employee</p>
                <Select
                  fullWidth
                  disabled
                  value={localStorage.getItem('uid')}
                  name="EmployeeID"
                  variant="outlined"
                  sx={{bgcolor: "rgba(0, 0, 0, 0.1)"}}
                >
                {employee.map((item: EmployeeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.FirstName}
                  </option>
                ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <FormControl fullWidth variant="outlined">
                <p>User</p>
                <Select
                  native
                  fullWidth 
                  value={borrowList.MemberID + ""}
                  onChange={handleChange}
                  inputProps={{name: "MemberID",}}
                >
                  <option aria-label="None" value="" >
                    Select User
                  </option>
                  {user.map((item: UsersInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.FirstName}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
        
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>Equipment</p>
                <Select
                  native
                  value={borrowList.EquipmentID + ""}
                  onChange={handleChange}
                  inputProps={{name: "EquipmentID",}}
                >
                  <option aria-label="None" value="" >
                    Select Equipment
                  </option>
                  {equipment.map((item: EquipmentInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Name}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
        
            <Grid item xs={3}>
              <p>Amount</p>
              <TextField 
                fullWidth 
                id="Amount"
                InputProps={{inputProps: {min: 1, max: 100}}}
                type="number"
                variant="outlined"
                defaultValue={0}
                onChange={handleInputChange}
              />
            </Grid>
        
            <Grid item xs={3}>
              <FormControl fullWidth>
                <p>Date & Time</p>
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                  <DateTimePicker
                    value={borrowList?.BorrowTime}
                      onChange={(newValue) => {
                        setBorrowList({
                          ...borrowList,
                          BorrowTime: newValue,
                        });
                      }}
                    renderInput={(props) => <TextField {...props} />}
                    ampm = {false}
                  />
                </LocalizationProvider>
              </FormControl>
            </Grid>
        
            <Grid item xs={6}>
              <Button
                component={RouterLink}
                to="/borrow_list"
                variant="contained"
                sx ={{p: 1, m: 2, mx: 'auto'}}
                color="inherit">
                  Back
              </Button>
              <Button 
                variant="contained" 
                color='success' 
                sx ={{p: 1, m: 2, mx: 'auto', float: "right"}}
                onClick={submit}>
                  Submit
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    </Container>
  );
}


export default App;
