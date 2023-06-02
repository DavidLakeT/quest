import { Button, Grid, TextField } from '@mui/material';
import AppBar from '../../components/appBar/Appbar';
import React from 'react'
import {useState} from 'react'
import './UploaderBox.css'
import { uploadFile } from '../../services/documentServices/DocumentServices';


function UploaderBox() {
  const [id, setId] = useState('');
  const [url, setUrl] = useState('');
  const [title, setTitle] = useState('');

  const handleUploadClick = async () => {
    try {
      const uploadData = {
        citizenId: id,
        documentUrl: url,
        documentTitle: title,
      }
      const upload = await uploadFile(uploadData);
      console.log('Uploaded document: ', upload);

      setId('')
      setUrl('')
      setTitle('')

    } catch (error) {
      console.error('Error uploading document', error.message);

    }
  };

  return (
    <div>
      <h1>Upload documents</h1>
      <Grid container display="row" justifyContent="center" gap={2}>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Id'  value={id} onChange={(event) => setId(parseInt(event.target.value))}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Url'  value={url} onChange={(event) => setUrl(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <TextField type="text" placeholder='Title'  value={title} onChange={(event) => setTitle(event.target.value)}/>
        </Grid>
        <Grid item xs={12} sm={12} md={12}>
        <Button variant='contained' onClick={handleUploadClick}>Upload documents</Button>
        </Grid>
      </Grid>
    </div>
  );
}
/*
function UploaderBox(){
    const [file,setFile] = useState()

    function handleFile(event){
        setFile(event.target.files[0])
    }
    function handleUpload(){
        const formData = new FormData()
        formData.append('file',file)
        fetch(
            'http://127.0.0.1:3001/apis/document/uploadDocument',
            {
                method:"POST",
                body: formData
            }
        ).then((response) => response.json())
        .then(
            (result) => {
                console.log('succes',result)
            }
        )
        .catch(error => {
            console.error("error:", error)
        })
    }
    return(
    )
}
        <div>
            <h1>File Upload</h1>
            <form onSubmit={handleUpload}>
                <input type="file" className='uploaded-row' name="file" accept='application/pdf' onChange={handleFile}/>
                <button>Upload</button>
            </form>
        </div>
*/
export default UploaderBox;