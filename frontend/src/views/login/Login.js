import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import { NavLink } from "react-router-dom";

function Login() {
  return (
    <AppBar>
      <h1>Login</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Usuario'/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Password'/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
            <Button variant='contained'>Iniciar sesion</Button>
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
