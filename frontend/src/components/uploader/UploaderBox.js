import React from 'react'
import {useState} from 'react'
import './UploaderBox.css'
import {MdCloudUpload, MdDelete} from 'react-icons/md'
import {AiFillFileImage} from 'react-icons/ai'

function UploaderBox(){
    const [image, setImage] = useState(null)
    const [fileName, setFileName] = useState("No selected file")
    return (
        <main>
            <form action=""
            onClick={() => document.querySelector(".input-field").click()}>
                <input type="file" accept='application/pdf' className='input-field' hidden
                    onChange={({target:{files}}) => {files[0] && setFileName(files[0].name)
                    if(files){
                        setImage(URL.createObjectURL(files[0]))
                    }
                }}/>
            
            {image ? <img src={image} width={300} height={300} alt={fileName}/> : <><MdCloudUpload color='#1475cf' size={60}/><p>Browse Files to upload</p></>}
            </form>
            <section className='uploaded-row'>
                <AiFillFileImage color='#1475cf'/>
                <span className='upload-content'>
                    {fileName} -
                    <MdDelete onClick={()=>{setFileName("No selected file"), setImage(null)}}/>
                </span>
            </section>
        </main>
    )
}

function FileUpload(){
    const [file,setFile] = useState()
    function handleFile(event){
        setFile(event.target.files[0])
    }
    return(
        <div>
            <h2>File Upload</h2>
            <form>
                <input type="file" name="file" onChange={handleFile}/>
                <button>Upload</button>
            </form>
        </div>
    )
}

export default UploaderBox;