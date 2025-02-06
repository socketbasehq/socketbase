import { Route, Routes } from 'react-router-dom';
import IndexPage from '@socketbase/pages/index';

import { AuthProvider } from './providers/auth';
import MainLayout from './components/layout-main';
import LoginPage from './pages/login';
import AppsIndexPage from './pages/apps';
import AppEditorPage from './pages/apps/editor';

function App() {
  return (
    <Routes>
      <Route path="/" element={<AuthProvider />}>
        <Route element={<MainLayout />}>
          <Route index element={<IndexPage />} />
          <Route path="apps" element={<AppsIndexPage />} />
          <Route path="apps/:id/*" element={<AppEditorPage />} />
        </Route>
      </Route>
      <Route path="/login" element={<LoginPage />} />
    </Routes>
  );
}

export default App;
