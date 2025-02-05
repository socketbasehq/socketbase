import { Route, Routes } from 'react-router-dom';
import IndexPage from '@socketbase/pages/index';
import LoginPage from './pages/login';

function App() {
  return (
    <Routes>
      <Route index path="/" element={<IndexPage />}></Route>
      <Route path="/login" element={<LoginPage />} />
    </Routes>
  );
}

export default App;
