import AppBar from '../../components/appBar/Appbar';
import UploaderBox from '../../components/uploader/UploaderBox';

function Uploads(){
    return (
        <AppBar>
            <h1>Uploads</h1>
            <div>
                <UploaderBox />
            </div>
        </AppBar>
    )
}

export default Uploads;