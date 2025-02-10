import { Route, Routes } from 'react-router-dom';
import loadable from '@loadable/component';
import AuthProvider from './providers/auth';
import MainLayout from './components/layout-main';

const IndexPage = loadable(() => import('@socketbase/pages/index'));
const LoginPage = loadable(() => import('./pages/login'));
const AppsIndexPage = loadable(() => import('./pages/apps'));
const AppEditorPage = loadable(() => import('./pages/apps/editor'));

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
