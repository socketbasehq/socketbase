import { App } from '@socketbase/api/apps';
import { Button } from '@socketbase/components/ui/button';
import { Input } from '@socketbase/components/ui/input';
import { Label } from '@socketbase/components/ui/label';
import { toast } from 'sonner';

type AppKeysProps = {
  app: App;
};

function AppKeys({ app }: AppKeysProps) {
  return (
    <div className="space-y-4">
      <div className="grid w-full max-w-sm items-center gap-1.5">
        <Label htmlFor="app-id">App ID</Label>
        <Input id="app-id" type="text" value={app.id} disabled />
      </div>

      <div className="grid w-full max-w-sm items-center gap-1.5">
        <Label htmlFor="app-key">App Key</Label>
        <Input id="app-key" type="text" value={app.app_key} disabled />
      </div>

      <div className="grid w-full max-w-sm items-center gap-1.5">
        <Label htmlFor="app-secret">App Secret</Label>
        <Input id="app-secret" type="text" value={app.app_secret} disabled />
      </div>

      <div className="space-x-2">
        <Button
          size="sm"
          variant="outline"
          onClick={() => {
            navigator.clipboard.writeText(
              `SOCKETBASE_APP_ID=${app.id}\nSOCKETBASE_APP_KEY=${app.app_key}\nSOCKETBASE_APP_SECRET=${app.app_secret}`,
            );
            toast.success('Copied to clipboard');
          }}
        >
          Copy to clipboard
        </Button>
        <Button size="sm">Revoke</Button>
      </div>
    </div>
  );
}

export default AppKeys;
