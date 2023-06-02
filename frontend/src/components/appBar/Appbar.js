import * as React from 'react';
import { NavLink } from "react-router-dom";
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import TemporaryDrawer from '../drawer/Drawer';
export default function ButtonAppBar({children}) {
  return (
    <div>
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
            <TemporaryDrawer/>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Citizen folder
          </Typography>
          <NavLink to="/login" style={{textDecoration: 'none', color: "white"}}>
            <Button color="inherit">Login</Button>
          </NavLink>
        </Toolbar>
      </AppBar>
    </Box>
    {children}
    </div>
  );
}