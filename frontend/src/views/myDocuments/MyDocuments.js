import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import React, { useState } from 'react';
import { citizenDocuments } from '../../services/citizenServices/CitizenServices';
import 'react-toastify/dist/ReactToastify.css';

function Documents() {
  const [id, setId] = useState('');
  const [documents, setDocuments] = useState(null);
  const handleDocumentClick = async () => {
    try {
      setDocuments(null);
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
          <h2 style={{marginTop: 90}}>Documents List:</h2>
          {documents === null ? (
            <p>No documents were found with that Citizen ID</p>
          ) : (
            <ul style={{ textAlign: 'center', listStyleType: 'none', padding: 0 }}>
              {documents.map((document, index) => (
                <li key={index} style={{ border: '1px solid black', borderRadius: '15px', padding: '15px', margin: '11px 680px 0 680px' }}>
                  <a href={document.URL} target="_blank" rel="noopener noreferrer" style={{ color: 'black' }}>
                    {document.Title}
                  </a>
                </li>
              ))}
            </ul>
          )}
        </div>
    </AppBar>
  );
}

export default Documents;