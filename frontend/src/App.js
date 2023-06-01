import './App.css';
import { Routes, Route } from 'react-router-dom';
import HomeView from './views/homeView/HomeView';
import Login from './views/login/Login';
import Register from './views/register/Register';
import MyDocuments from './views/myDocuments/MyDocuments';
import Uploads from './views/uploads/Uploads';
import RequestDocument from './views/requestDocument/RequestDocument';



function App() {
  return (
    <div className="App">
      <Routes>
          <Route path="/" element={<HomeView />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/documents" element={<MyDocuments />} />
          <Route path="/upload" element={<Uploads />} />
          <Route path="/request" element={<RequestDocument />} />
       </Routes>
    </div>
  );
}



export default App;
