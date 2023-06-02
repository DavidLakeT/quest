import AppBar from '../../components/appBar/Appbar';
import UploaderBox from '../../components/uploader/UploaderBox';

function Uploads(){
    return (
    <div>
        <AppBar/>
        <div className='uploadBox'>
            <UploaderBox />
        </div>
    </div>
    )
}

export default Uploads;