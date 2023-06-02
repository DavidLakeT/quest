import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import React, { useState } from 'react';
import { citizenDocuments } from '../../services/citizenServices/CitizenServices';

function Documents() {
  const [id, setId] = useState('');
  const [documents, setDocuments] = useState([]);
  const handleDocumentClick = async () => {
    try {
      const documents = await citizenDocuments(id);
      console.log('Documents ', documents.message);
      setDocuments(documents.message)

      setId('')

    } catch (error) {
      console.error('Error al crear ciudadano:', error.message);

    }
  };

  return (
    <AppBar>
      <h1>Documents</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Id'  value={id} onChange={(event) => setId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <Button variant='contained' onClick={handleDocumentClick}>Search documents</Button>
        </Grid>
      </Grid>
      <div>
        <h2>Documents List:</h2>
        <ul style={{ textAlign: 'center' }}>
        {documents.map((document, index) => (
          <li key={index}>
            <a href={document.URL} target="_blank" rel="noopener noreferrer">
              {document.Title}
            </a>
          </li>
        ))}
      </ul>
      </div>
    </AppBar>
  );
}

export default Documents;