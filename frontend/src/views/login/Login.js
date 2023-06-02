import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import { NavLink } from "react-router-dom";
import React, { useState } from 'react';

function Login() {
  const [id, setId] = useState('');
  const [password, setPassword] = useState('');

  const handleLoginClick = async () => {
    try {
      const citizenData = {
        citizenId: id,
        password: password
      }
      //const createdCitizen = await registerCitizen(citizenData);
      //console.log('Ciudadano creado:', createdCitizen);

      setId('')
      setPassword('')

    } catch (error) {
      console.error('Error al crear ciudadano:', error.message);

    }
  }

  return (
    <AppBar>
      <h1>Login</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Id'  value={id} onChange={(event) => setId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Password'  value={password} onChange={(event) => setId(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
            <Button variant='contained'>Login</Button>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
                <NavLink to="/register">
                    SignUp
                </NavLink>
        </Grid>
      </Grid>
    </AppBar>
  );
}

export default Login;
