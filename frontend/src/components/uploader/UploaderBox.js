import React from 'react'
import {useState} from 'react'
import './UploaderBox.css'

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
        <div>
            <h1>File Upload</h1>
            <form onSubmit={handleUpload}>
                <input type="file" className='uploaded-row' name="file" accept='application/pdf' onChange={handleFile}/>
                <button>Upload</button>
            </form>
        </div>
    )
}

export default UploaderBox;




