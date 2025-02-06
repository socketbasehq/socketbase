import { useGetAppQuery } from '@socketbase/api/apps';
import {
  LayoutIcon,
  RocketIcon,
  KeyIcon,
  FunctionSquareIcon,
  BarChart3Icon,
  TerminalIcon,
  AlertTriangleIcon,
  WebhookIcon,
  UsersIcon,
  SettingsIcon,
  Loader2,
  ArrowLeft,
} from 'lucide-react';
import {
  Link,
  Route,
  Routes,
  useLocation,
  useNavigate,
  useParams,
} from 'react-router-dom';
import AppOverview from './id/overview';
import AppKeys from './id/keys';
import { Button } from '@socketbase/components/ui/button';

const sidebarItems = [
  {
    title: 'Overview',
    icon: LayoutIcon,
    path: '/apps/:id',
  },
  {
    title: 'Getting Started',
    icon: RocketIcon,
    path: '/apps/:id/getting-started',
  },
  {
    title: 'App Keys',
    icon: KeyIcon,
    path: '/apps/:id/keys',
  },
  {
    title: 'Functions',
    icon: FunctionSquareIcon,
    path: '/apps/:id/functions',
  },
  {
    title: 'Stats',
    icon: BarChart3Icon,
    path: '/apps/:id/stats',
  },
  {
    title: 'Debug Console',
    icon: TerminalIcon,
    path: '/apps/:id/debug',
  },
  {
    title: 'Error Logs',
    icon: AlertTriangleIcon,
    path: '/apps/:id/errors',
  },
  {
    title: 'Webhooks',
    icon: WebhookIcon,
    path: '/apps/:id/webhooks',
  },
  {
    title: 'Collaborators',
    icon: UsersIcon,
    path: '/apps/:id/collaborators',
  },
  {
    title: 'App Settings',
    icon: SettingsIcon,
    path: '/apps/:id/settings',
  },
];

function AppEditorPage() {
  const { id } = useParams();
  const location = useLocation();
  const navigate = useNavigate();

  const { data, isLoading } = useGetAppQuery({ id: id || '' });

  if (isLoading) {
    return (
      <div className="h-40 flex justify-center items-center">
        <Loader2 className="animate-spin" size={20} />
      </div>
    );
  }

  if (!data?.data) {
    return (
      <div className="flex flex-col justify-center items-center h-60">
        <h1 className="text-2xl font-bold">App not found</h1>
        <p className="text-sm text-muted-foreground">
          The app you are looking for does not exist.
        </p>
        <Link to="/apps" className="mt-4">
          <Button variant={'outline'}>Go to apps</Button>
        </Link>
      </div>
    );
  }

  return (
    <div>
      <div className="flex items-center gap-2">
        <ArrowLeft
          onClick={() => navigate('/apps')}
          className="cursor-pointer"
          size={20}
        />
        <h1 className="text-xl font-bold">{data?.data.name}</h1>
      </div>

      <div className="flex items-start gap-8 mt-5">
        <nav className="w-52">
          {sidebarItems.map(item => {
            const path = item.path.replace(':id', id || '');
            const isActive = location.pathname === path;
            return (
              <Link
                to={path}
                draggable={false}
                className={`flex text-sm items-center gap-2 select-none px-4 py-2 rounded-xl ${
                  isActive ? 'bg-primary/5 text-primary' : 'hover:bg-muted'
                }`}
              >
                <item.icon size={20} />
                <span>{item.title}</span>
              </Link>
            );
          })}
        </nav>
        <main className=" flex-1">
          <Routes>
            <Route path="/" element={<AppOverview />} />
            <Route path="/keys" element={<AppKeys app={data?.data} />} />
          </Routes>
        </main>
      </div>
    </div>
  );
}

export default AppEditorPage;
