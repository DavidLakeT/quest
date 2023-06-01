import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import { NavLink } from "react-router-dom";
import { registerCitizen } from '../../services/citizenServices/CitizenServices';

const handleRegisterClick = async () => {
    try {
      const citizenData = {
        citizenId: 123456789,
        name: "Jacobo",
        address: "Envigado",
        email: "jacobo@eafit.edu.co",
        operatorid: 1
      }

      const createdCitizen = await registerCitizen(citizenData);
      console.log('Ciudadano creado:', createdCitizen);

      // Realiza cualquier acción adicional después de crear el ciudadano, como mostrar un mensaje de éxito, redirigir a otra página, etc.
    } catch (error) {
      console.error('Error al crear ciudadano:', error.message);

      // Realiza cualquier acción adicional en caso de error, como mostrar un mensaje de error al usuario.
    }
  };

function Register() {
  return (
    <AppBar>
      <h1>Register</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Id'/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Usuario'/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Password'/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <Button variant='contained' onClick={handleRegisterClick}>Registrar usuario</Button>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
                <NavLink to="/login">
                    LogIn
                </NavLink>
        </Grid>
      </Grid>
    </AppBar>
  );
}

export default Register;
