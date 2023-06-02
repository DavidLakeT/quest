import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import { NavLink } from "react-router-dom";
import { registerCitizen } from '../../services/citizenServices/CitizenServices';
import React, { useState } from 'react';
import { toast } from 'react-toastify';

function Register() {
  const [id, setId] = useState('');
  const [name, setName] = useState('');
  const [address, setAddress] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [operatorId, setOperatorId] = useState('');

  const handleRegisterClick = async () => {
    try {
      const citizenData = {
        citizenId: id,
        name: name,
        address: address,
        email: email,
        password: password,
        operatorid: operatorId
      }
      const createdCitizen = await registerCitizen(citizenData);
      console.log('Ciudadano creado:', createdCitizen);

      setId('')
      setName('')
      setAddress('')
      setEmail('')
      setPassword('')
      setOperatorId('')

      toast.success('Succesfully registered', {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 2000,
        hideProgressBar: true,
      });

    } catch (error) {
      console.error('Error al crear ciudadano:', error.message);

    }
  };

  return (
    <AppBar>
      <h1>Register</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Id'  value={id} onChange={(event) => setId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Name'  value={name} onChange={(event) => setName(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Address'  value={address} onChange={(event) => setAddress(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='email'  value={email} onChange={(event) => setEmail(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="password" placeholder='password'  value={password} onChange={(event) => setPassword(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Operator Id'  value={operatorId} onChange={(event) => setOperatorId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <Button variant='contained' onClick={handleRegisterClick}>Register citizen</Button>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
                <NavLink to="/login">
                    Login
                </NavLink>
        </Grid>
      </Grid>
    </AppBar>
  );
}

export default Register;
