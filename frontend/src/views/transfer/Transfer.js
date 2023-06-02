import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import React, { useState } from 'react';
import { transferCitizen } from '../../services/citizenServices/CitizenServices';

function Transfer() {
  const [id, setId] = useState('');
  const [currentOperatorId, setCurrentOperatorId] = useState('');
  const [newOperatorId, setNewOperatorId] = useState('');

  const handleTransferClick = async () => {
    try {
      const transferData = {
        citizenId: id,
        currentOperatorId: currentOperatorId,
        newOperatorId: newOperatorId
      }

      const transfer = await transferCitizen(transferData);
      console.log('Transfered citizen:', transfer);

      setId('')
      setCurrentOperatorId('')
      setNewOperatorId('')

    } catch (error) {
      console.error('Error al crear ciudadano:', error.message);

    }
  }

  return (
    <AppBar>
      <h1>Transfer</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Id'  value={id} onChange={(event) => setId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Current operator id'  value={currentOperatorId} onChange={(event) => setCurrentOperatorId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='New operator id'  value={newOperatorId} onChange={(event) => setNewOperatorId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
            <Button variant='contained' onClick={handleTransferClick}>Transfer citizen</Button>
        </Grid>
      </Grid>
    </AppBar>
  );
}

export default Transfer;