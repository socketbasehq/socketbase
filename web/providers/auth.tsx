import { useAuth } from '@socketbase/lib/hooks/use-auth';
import { Loader2 } from 'lucide-react';
import { Outlet, useNavigate } from 'react-router-dom';

function AuthProvider() {
  const navigate = useNavigate();
  const { isLoading, error } = useAuth();

  if (isLoading)
    return (
      <div className="flex h-screen w-screen items-center justify-center">
        <Loader2 className="animate-spin" />
      </div>
    );
  if (error) {
    if ('status' in error && error.status === 401) {
      navigate('/login');
      return;
    }

    return (
      <div className="flex h-screen w-screen items-center justify-center">
        Error: {JSON.stringify(error)}
      </div>
    );
  }

  return <Outlet />;
}

export default AuthProvider;
