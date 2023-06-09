import * as React from 'react';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import Button from '@mui/material/Button';
import List from '@mui/material/List';
import Divider from '@mui/material/Divider';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import MenuIcon from '@mui/icons-material/Menu';
import DocumentScannerIcon from '@mui/icons-material/DocumentScanner';
import CloudUploadIcon from '@mui/icons-material/CloudUpload';
import AssignmentTurnedInIcon from '@mui/icons-material/AssignmentTurnedIn';
import PublishedWithChangesIcon from '@mui/icons-material/PublishedWithChanges';
import { NavLink } from "react-router-dom";

export default function TemporaryDrawer() {
  const [state, setState] = React.useState({
    left: false,
  });

  const toggleDrawer = (anchor, open) => (event) => {
    if (event.type === 'keydown' && (event.key === 'Tab' || event.key === 'Shift')) {
      return;
    }

    setState({ ...state, [anchor]: open });
  };

  const list = (anchor) => (
    <Box
      sx={{ width: anchor === 'top' || anchor === 'bottom' ? 'auto' : 250 }}
      role="presentation"
      onClick={toggleDrawer(anchor, false)}
      onKeyDown={toggleDrawer(anchor, false)}
    >
      <List>
        <ListItem>
        <NavLink to="/documents" style={{textDecoration: 'none', color: "black"}}>
          <ListItemButton>
              <ListItemIcon>
                <DocumentScannerIcon/>
              </ListItemIcon>
              <ListItemText primary='My documents' />
            </ListItemButton>
        </NavLink>
        </ListItem>
        <Divider/>
        <ListItem>
        <NavLink to="/upload" style={{textDecoration: 'none', color: "black"}}>
          <ListItemButton>
              <ListItemIcon>
                <CloudUploadIcon/>
              </ListItemIcon>
              <ListItemText primary='Upload' />
            </ListItemButton>
        </NavLink>
        </ListItem>
        <Divider/>
        <ListItem>
          <NavLink to="/request" style={{textDecoration: 'none', color: "black"}}>
            <ListItemButton>
                <ListItemIcon>
            <AssignmentTurnedInIcon/>
            </ListItemIcon>
            <ListItemText primary='Request document' />
            </ListItemButton>
          </NavLink>
        </ListItem>
        <Divider/>
        <ListItem>
          <NavLink to="/transfer" style={{textDecoration: 'none', color: "black"}}>
            <ListItemButton>
                <ListItemIcon>
            <PublishedWithChangesIcon/>
            </ListItemIcon>
            <ListItemText primary='Transfer operator' />
            </ListItemButton>
          </NavLink>
        </ListItem>
        <Divider/>
      </List>
    </Box>
  );

  return (
    <div>
      {['left'].map((anchor) => (
        <React.Fragment key={anchor}>
          <Button onClick={toggleDrawer(anchor, true)} sx={{color: 'white'}}><MenuIcon /></Button>
          <Drawer
            anchor={anchor}
            open={state[anchor]}
            onClose={toggleDrawer(anchor, false)}
          >
            {list(anchor)}
          </Drawer>
        </React.Fragment>
      ))}
    </div>
  );
}